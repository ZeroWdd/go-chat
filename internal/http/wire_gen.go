// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"context"

	"github.com/google/wire"
	_ "github.com/urfave/cli/v2"
	"go-chat/config"
	"go-chat/internal/http/internal/handler"
	"go-chat/internal/http/internal/handler/admin"
	v1_2 "go-chat/internal/http/internal/handler/admin/v1"
	"go-chat/internal/http/internal/handler/open"
	v1_3 "go-chat/internal/http/internal/handler/open/v1"
	"go-chat/internal/http/internal/handler/web"
	"go-chat/internal/http/internal/handler/web/v1"
	"go-chat/internal/http/internal/handler/web/v1/article"
	"go-chat/internal/http/internal/handler/web/v1/contact"
	"go-chat/internal/http/internal/handler/web/v1/group"
	"go-chat/internal/http/internal/handler/web/v1/talk"
	"go-chat/internal/http/internal/router"
	"go-chat/internal/logic"
	_ "go-chat/internal/pkg/validation"
	"go-chat/internal/provider"
	"go-chat/internal/repository/cache"
	"go-chat/internal/repository/dao"
	"go-chat/internal/repository/dao/note"
	"go-chat/internal/repository/dao/organize"
	"go-chat/internal/service"
	note2 "go-chat/internal/service/note"
	organize2 "go-chat/internal/service/organize"
)

// Injectors from wire.go:

func Initialize(ctx context.Context, conf *config.Config) *AppProvider {
	client := provider.NewRedisClient(ctx, conf)
	smsCodeCache := cache.NewSmsCodeCache(client)
	smsService := service.NewSmsService(smsCodeCache)
	db := provider.NewMySQLClient(conf)
	baseDao := dao.NewBaseDao(db, client)
	usersDao := dao.NewUserDao(baseDao)
	userService := service.NewUserService(usersDao)
	common := v1.NewCommon(conf, smsService, userService)
	tokenSessionStorage := cache.NewTokenSessionStorage(client)
	redisLock := cache.NewRedisLock(client)
	baseService := service.NewBaseService(db, client)
	unreadStorage := cache.NewUnreadStorage(client)
	messageStorage := cache.NewMessageStorage(client)
	talkVote := cache.NewTalkVote(client)
	talkRecordsVoteDao := dao.NewTalkRecordsVoteDao(baseDao, talkVote)
	relation := cache.NewRelation(client)
	groupMemberDao := dao.NewGroupMemberDao(baseDao, relation)
	serverStorage := cache.NewSidStorage(client)
	clientStorage := cache.NewClientStorage(client, conf, serverStorage)
	filesystem := provider.NewFilesystem(conf)
	splitUploadDao := dao.NewFileSplitUploadDao(baseDao)
	talkMessageService := service.NewTalkMessageService(baseService, conf, unreadStorage, messageStorage, talkRecordsVoteDao, groupMemberDao, serverStorage, clientStorage, filesystem, splitUploadDao)
	httpClient := provider.NewHttpClient()
	requestClient := provider.NewRequestClient(httpClient)
	ipAddressService := service.NewIpAddressService(baseService, conf, requestClient)
	talkSessionDao := dao.NewTalkSessionDao(baseDao)
	talkSessionService := service.NewTalkSessionService(baseService, talkSessionDao)
	articleClassDao := note.NewArticleClassDao(baseDao)
	articleClassService := note2.NewArticleClassService(baseService, articleClassDao)
	robotDao := dao.NewRobotDao(baseDao)
	messageForwardLogic := logic.NewMessageForwardLogic(db)
	sequence := cache.NewSequence(client)
	messageService := service.NewMessageService(baseService, messageForwardLogic, groupMemberDao, splitUploadDao, filesystem, unreadStorage, messageStorage, serverStorage, clientStorage, sequence)
	auth := v1.NewAuth(conf, userService, smsService, tokenSessionStorage, redisLock, talkMessageService, ipAddressService, talkSessionService, articleClassService, robotDao, messageService)
	organizeDao := organize.NewOrganizeDao(baseDao)
	organizeService := organize2.NewOrganizeService(baseService, organizeDao)
	user := v1.NewUser(userService, smsService, organizeService)
	departmentDao := organize.NewDepartmentDao(baseDao)
	deptService := organize2.NewOrganizeDeptService(baseService, departmentDao)
	positionDao := organize.NewPositionDao(baseDao)
	positionService := organize2.NewPositionService(baseService, positionDao)
	v1Organize := v1.NewOrganize(deptService, organizeService, positionService)
	talkService := service.NewTalkService(baseService, groupMemberDao)
	contactRemark := cache.NewContactRemark(client)
	contactDao := dao.NewContactDao(baseDao, contactRemark, relation)
	contactService := service.NewContactService(baseService, contactDao)
	groupDao := dao.NewGroupDao(baseDao)
	groupService := service.NewGroupService(baseService, groupDao, groupMemberDao, relation)
	authPermissionService := service.NewAuthPermissionService(contactDao, groupMemberDao, organizeDao)
	session := talk.NewSession(talkService, talkSessionService, redisLock, userService, clientStorage, messageStorage, contactService, unreadStorage, contactRemark, groupService, authPermissionService)
	splitUploadService := service.NewSplitUploadService(baseService, splitUploadDao, conf, filesystem)
	groupMemberService := service.NewGroupMemberService(baseService, groupMemberDao)
	talkAuthService := service.NewTalkAuthService(organizeDao, contactDao)
	message := talk.NewMessage(talkMessageService, talkService, talkRecordsVoteDao, splitUploadService, contactService, groupMemberService, organizeService, talkAuthService, messageService)
	talkRecordsDao := dao.NewTalkRecordsDao(baseDao)
	talkRecordsService := service.NewTalkRecordsService(baseService, talkVote, talkRecordsVoteDao, groupMemberDao, talkRecordsDao)
	records := talk.NewRecords(talkRecordsService, groupMemberService, filesystem, authPermissionService)
	emoticonDao := dao.NewEmoticonDao(baseDao)
	emoticonService := service.NewEmoticonService(baseService, emoticonDao, filesystem)
	emoticon := v1.NewEmoticon(filesystem, emoticonService, redisLock)
	upload := v1.NewUpload(conf, filesystem, splitUploadService)
	groupNoticeDao := dao.NewGroupNoticeDao(baseDao)
	groupNoticeService := service.NewGroupNoticeService(groupNoticeDao)
	groupGroup := group.NewGroup(groupService, groupMemberService, talkSessionService, userService, redisLock, contactService, groupNoticeService, talkMessageService)
	notice := group.NewNotice(groupNoticeService, groupMemberService)
	groupApplyDao := dao.NewGroupApply(baseDao)
	groupApplyService := service.NewGroupApplyService(baseService, groupApplyDao)
	apply := group.NewApply(groupApplyService, groupMemberService, groupService)
	contactContact := contact.NewContact(contactService, clientStorage, userService, talkSessionService, talkMessageService, organizeService)
	contactApplyService := service.NewContactsApplyService(baseService)
	contactApply := contact.NewApply(contactApplyService, userService, talkMessageService, contactService)
	articleService := note2.NewArticleService(baseService)
	articleAnnexDao := note.NewArticleAnnexDao(baseDao)
	articleAnnexService := note2.NewArticleAnnexService(baseService, articleAnnexDao, filesystem)
	articleArticle := article.NewArticle(articleService, filesystem, articleAnnexService)
	annex := article.NewAnnex(articleAnnexService, filesystem)
	class := article.NewClass(articleClassService)
	articleTagService := note2.NewArticleTagService(baseService)
	tag := article.NewTag(articleTagService)
	sendMessage := talk.NewSendMessage(talkAuthService, messageService)
	webV1 := &web.V1{
		Common:        common,
		Auth:          auth,
		User:          user,
		Organize:      v1Organize,
		Talk:          session,
		TalkMessage:   message,
		TalkRecords:   records,
		Emoticon:      emoticon,
		Upload:        upload,
		Group:         groupGroup,
		GroupNotice:   notice,
		GroupApply:    apply,
		Contact:       contactContact,
		ContactsApply: contactApply,
		Article:       articleArticle,
		ArticleAnnex:  annex,
		ArticleClass:  class,
		ArticleTag:    tag,
		Message:       sendMessage,
	}
	webHandler := &web.Handler{
		V1: webV1,
	}
	index := v1_2.NewIndex()
	v1Auth := v1_2.NewAuth()
	adminV1 := &admin.V1{
		Index: index,
		Auth:  v1Auth,
	}
	v2 := &admin.V2{}
	adminHandler := &admin.Handler{
		V1: adminV1,
		V2: v2,
	}
	v1Index := v1_3.NewIndex()
	openV1 := &open.V1{
		Index: v1Index,
	}
	openHandler := &open.Handler{
		V1: openV1,
	}
	handlerHandler := &handler.Handler{
		Api:   webHandler,
		Admin: adminHandler,
		Open:  openHandler,
	}
	engine := router.NewRouter(conf, handlerHandler, tokenSessionStorage)
	httpServer := provider.NewHttpServer(conf, engine)
	appProvider := &AppProvider{
		Config: conf,
		Server: httpServer,
	}
	return appProvider
}

// wire.go:

var providerSet = wire.NewSet(provider.NewMySQLClient, provider.NewRedisClient, provider.NewHttpClient, provider.NewEmailClient, provider.NewHttpServer, provider.NewFilesystem, provider.NewRequestClient, router.NewRouter, wire.Struct(new(web.Handler), "*"), wire.Struct(new(admin.Handler), "*"), wire.Struct(new(open.Handler), "*"), wire.Struct(new(handler.Handler), "*"), wire.Struct(new(AppProvider), "*"))

var cacheProviderSet = wire.NewSet(cache.NewTokenSessionStorage, cache.NewSidStorage, cache.NewUnreadStorage, cache.NewRedisLock, cache.NewClientStorage, cache.NewMessageStorage, cache.NewTalkVote, cache.NewRoomStorage, cache.NewRelation, cache.NewSmsCodeCache, cache.NewContactRemark, cache.NewSequence)

var daoProviderSet = wire.NewSet(dao.NewBaseDao, dao.NewContactDao, dao.NewGroupMemberDao, dao.NewUserDao, dao.NewGroupDao, dao.NewGroupApply, dao.NewTalkRecordsDao, dao.NewGroupNoticeDao, dao.NewTalkSessionDao, dao.NewEmoticonDao, dao.NewTalkRecordsVoteDao, dao.NewFileSplitUploadDao, note.NewArticleClassDao, note.NewArticleAnnexDao, organize.NewDepartmentDao, organize.NewOrganizeDao, organize.NewPositionDao, dao.NewRobotDao)

var serviceProviderSet = wire.NewSet(service.NewBaseService, service.NewUserService, service.NewSmsService, service.NewTalkService, service.NewTalkMessageService, service.NewGroupService, service.NewGroupMemberService, service.NewGroupNoticeService, service.NewGroupApplyService, service.NewTalkSessionService, service.NewEmoticonService, service.NewTalkRecordsService, service.NewContactService, service.NewContactsApplyService, service.NewSplitUploadService, service.NewIpAddressService, service.NewAuthPermissionService, service.NewMessageService, note2.NewArticleService, note2.NewArticleTagService, note2.NewArticleClassService, note2.NewArticleAnnexService, organize2.NewOrganizeDeptService, organize2.NewOrganizeService, organize2.NewPositionService, service.NewTemplateService, service.NewTalkAuthService, logic.NewMessageForwardLogic)
