package v1

import (
	"strconv"
	"time"

	"github.com/mojocn/base64Captcha"
	"go-chat/api/pb/admin/v1"
	"go-chat/config"
	"go-chat/internal/pkg/encrypt"
	"go-chat/internal/pkg/ichat"
	"go-chat/internal/pkg/jwt"
	"go-chat/internal/repository/cache"
	"go-chat/internal/repository/model"
	"go-chat/internal/repository/repo"
	"gorm.io/gorm"
)

type Auth struct {
	config          *config.Config
	captcha         *base64Captcha.Captcha
	admin           *repo.Admin
	jwtTokenStorage *cache.JwtTokenStorage
}

func NewAuth(config *config.Config, captcha *cache.CaptchaStorage, admin *repo.Admin, jwtTokenStorage *cache.JwtTokenStorage) *Auth {
	return &Auth{config: config, captcha: base64Captcha.NewCaptcha(base64Captcha.DefaultDriverDigit, captcha), admin: admin, jwtTokenStorage: jwtTokenStorage}
}

// Login 登录接口
func (c *Auth) Login(ctx *ichat.Context) error {

	var in admin.AuthLoginRequest
	if err := ctx.Context.ShouldBindJSON(&in); err != nil {
		return ctx.InvalidParams(err)
	}

	if !c.captcha.Verify(in.CaptchaVoucher, in.Captcha, true) {
		return ctx.InvalidParams("验证码填写不正确")
	}

	adminInfo, err := c.admin.FindByWhere(ctx.Ctx(), "username = ? or email = ?", in.Username, in.Username)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return ctx.InvalidParams("账号不存在或密码填写错误!")
		}

		return ctx.Error(err.Error())
	}

	password, err := encrypt.RsaDecrypt(in.Password, c.config.App.PrivateKey)
	if err != nil {
		return ctx.Error(err.Error())
	}

	if !encrypt.VerifyPassword(adminInfo.Password, string(password)) {
		return ctx.InvalidParams("账号不存在或密码填写错误!")
	}

	if adminInfo.Status != model.AdminStatusNormal {
		return ctx.ErrorBusiness("账号已被管理员禁用，如有问题请联系管理员！")
	}

	expiresAt := time.Now().Add(12 * time.Hour)

	// 生成登录凭证
	token := jwt.GenerateToken("admin", c.config.Jwt.Secret, &jwt.Options{
		ExpiresAt: jwt.NewNumericDate(expiresAt),
		ID:        strconv.Itoa(adminInfo.Id),
		Issuer:    "im.admin",
		IssuedAt:  jwt.NewNumericDate(time.Now()),
	})

	return ctx.Success(&admin.AuthLoginResponse{
		Auth: &admin.AccessToken{
			Type:        "Bearer",
			AccessToken: token,
			ExpiresIn:   int32(expiresAt.Unix() - time.Now().Unix()),
		},
	})
}

// Captcha 图形验证码
func (c *Auth) Captcha(ctx *ichat.Context) error {

	generate, base64, err := c.captcha.Generate()
	if err != nil {
		return ctx.ErrorBusiness(err)
	}

	return ctx.Success(&admin.AuthCaptchaResponse{
		Voucher: generate,
		Captcha: base64,
	})
}

// Logout 退出登录接口
func (c *Auth) Logout(ctx *ichat.Context) error {

	session := ctx.JwtSession()
	if session != nil {
		if ex := session.ExpiresAt - time.Now().Unix(); ex > 0 {
			_ = c.jwtTokenStorage.SetBlackList(ctx.Ctx(), session.Token, time.Duration(ex)*time.Second)
		}
	}

	return ctx.Success(nil)
}

// Refresh Token 刷新接口
func (c *Auth) Refresh(ctx *ichat.Context) error {

	// TODO 业务逻辑 ...

	return ctx.Success(nil)
}
