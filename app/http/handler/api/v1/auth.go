package v1

import (
	"fmt"
	"go-chat/app/http/dto"
	"go-chat/app/pkg/utils"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"go-chat/app/cache"
	"go-chat/app/entity"
	"go-chat/app/http/request"
	"go-chat/app/http/response"
	"go-chat/app/pkg/auth"
	"go-chat/app/service"
	"go-chat/config"
)

type Auth struct {
	config             *config.Config
	userService        *service.UserService
	smsService         *service.SmsService
	session            *cache.Session
	redisLock          *cache.RedisLock
	talkMessageService *service.TalkMessageService
	ipAddressService   *service.IpAddressService
}

func NewAuthHandler(
	config *config.Config,
	userService *service.UserService,
	smsService *service.SmsService,
	session *cache.Session,
	redisLock *cache.RedisLock,
	talkMessageService *service.TalkMessageService,
	ipAddressService *service.IpAddressService,
) *Auth {
	return &Auth{
		config:             config,
		userService:        userService,
		smsService:         smsService,
		session:            session,
		redisLock:          redisLock,
		talkMessageService: talkMessageService,
		ipAddressService:   ipAddressService,
	}
}

// Login 登录接口
func (c *Auth) Login(ctx *gin.Context) {
	params := &request.LoginRequest{}
	if err := ctx.ShouldBind(params); err != nil {
		response.InvalidParams(ctx, err)
		return
	}

	user, err := c.userService.Login(params.Mobile, params.Password)
	if err != nil {
		response.InvalidParams(ctx, err)
		return
	}

	ip, _ := utils.GetIP(ctx.Request)

	address, err := c.ipAddressService.FindAddress(ip)
	if err == nil {
		fmt.Println(address)
	}

	_ = c.talkMessageService.SendLoginMessage(ctx.Request.Context(), &service.LoginInfo{
		UserId:   user.Id,
		Ip:       ip,
		Address:  address,
		Platform: params.Platform,
		Agent:    ctx.GetHeader("user-agent"),
	})

	response.Success(ctx, c.createToken(user.Id))
}

// Register 注册接口
func (c *Auth) Register(ctx *gin.Context) {
	params := &request.RegisterRequest{}
	if err := ctx.ShouldBind(params); err != nil {
		response.InvalidParams(ctx, err)
		return
	}

	// 验证短信验证码是否正确
	if !c.smsService.CheckSmsCode(ctx.Request.Context(), entity.SmsRegisterChannel, params.Mobile, params.SmsCode) {
		response.InvalidParams(ctx, "短信验证码填写错误！")
		return
	}

	if _, err := c.userService.Register(params); err != nil {
		response.BusinessError(ctx, err)
		return
	}

	c.smsService.DeleteSmsCode(ctx.Request.Context(), entity.SmsRegisterChannel, params.Mobile)

	response.Success(ctx, gin.H{}, "注册成功！")
}

// Logout 退出登录接口
func (c *Auth) Logout(ctx *gin.Context) {
	c.toBlackList(ctx)

	response.Success(ctx, gin.H{}, "已退出登录！")
}

// Refresh Token 刷新接口
func (c *Auth) Refresh(ctx *gin.Context) {
	tokenInfo := c.createToken(auth.GetAuthUserID(ctx))

	c.toBlackList(ctx)

	response.Success(ctx, tokenInfo)
}

// Forget 账号找回接口
func (c *Auth) Forget(ctx *gin.Context) {
	params := &request.ForgetRequest{}

	if err := ctx.ShouldBind(params); err != nil {
		response.InvalidParams(ctx, err)
		return
	}

	// 验证短信验证码是否正确
	if !c.smsService.CheckSmsCode(ctx.Request.Context(), entity.SmsForgetAccountChannel, params.Mobile, params.SmsCode) {
		response.InvalidParams(ctx, "短信验证码填写错误！")
		return
	}

	if _, err := c.userService.Forget(params); err != nil {
		response.BusinessError(ctx, err)
		return
	}

	c.smsService.DeleteSmsCode(ctx.Request.Context(), entity.SmsForgetAccountChannel, params.Mobile)

	response.Success(ctx, gin.H{}, "账号成功找回")
}

func (c *Auth) createToken(uid int) *dto.Token {
	expiresAt := time.Now().Add(time.Second * time.Duration(c.config.Jwt.ExpiresTime)).Unix()

	// 生成登录凭证
	token := auth.SignJwtToken("api", c.config.Jwt.Secret, &auth.JwtOptions{
		ExpiresAt: expiresAt,
		Id:        strconv.Itoa(uid),
	})

	return &dto.Token{
		Type:      "Bearer",
		Token:     token,
		ExpiresIn: expiresAt,
	}
}

// 设置黑名单
func (c *Auth) toBlackList(ctx *gin.Context) {
	info := ctx.GetStringMapString("jwt")

	expiresAt, _ := strconv.Atoi(info["expires_at"])

	ex := expiresAt - int(time.Now().Unix())

	// 将 session 加入黑名单
	_ = c.session.SetBlackList(ctx.Request.Context(), info["session"], ex)
}
