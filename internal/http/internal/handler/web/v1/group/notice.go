package group

import (
	"go-chat/api/pb/web/v1"
	"go-chat/internal/entity"
	"go-chat/internal/pkg/ichat"
	"go-chat/internal/pkg/jsonutil"
	"go-chat/internal/pkg/timeutil"
	"go-chat/internal/repository/model"

	"go-chat/internal/service"
)

type Notice struct {
	groupNoticeService *service.GroupNoticeService
	groupMemberService *service.GroupMemberService
	messageService     *service.MessageService
}

func NewNotice(groupNoticeService *service.GroupNoticeService, groupMemberService *service.GroupMemberService, messageService *service.MessageService) *Notice {
	return &Notice{groupNoticeService: groupNoticeService, groupMemberService: groupMemberService, messageService: messageService}
}

// CreateAndUpdate 添加或编辑群公告
func (c *Notice) CreateAndUpdate(ctx *ichat.Context) error {

	params := &web.GroupNoticeEditRequest{}
	if err := ctx.Context.ShouldBindJSON(params); err != nil {
		return ctx.InvalidParams(err)
	}

	uid := ctx.UserId()

	if !c.groupMemberService.Dao().IsLeader(ctx.Ctx(), int(params.GroupId), uid) {
		return ctx.ErrorBusiness("无权限操作")
	}

	var (
		msg string
		err error
	)

	if params.NoticeId == 0 {
		err = c.groupNoticeService.Create(ctx.Ctx(), &service.GroupNoticeEditOpt{
			UserId:    uid,
			GroupId:   int(params.GroupId),
			NoticeId:  int(params.NoticeId),
			Title:     params.Title,
			Content:   params.Content,
			IsTop:     int(params.IsTop),
			IsConfirm: int(params.IsConfirm),
		})
		msg = "添加群公告成功！"
	} else {
		err = c.groupNoticeService.Update(ctx.Ctx(), &service.GroupNoticeEditOpt{
			GroupId:   int(params.GroupId),
			NoticeId:  int(params.NoticeId),
			Title:     params.Title,
			Content:   params.Content,
			IsTop:     int(params.IsTop),
			IsConfirm: int(params.IsConfirm),
		})
		msg = "更新群公告成功！"
	}

	if err != nil {
		return ctx.ErrorBusiness(err.Error())
	}

	_ = c.messageService.SendSysOther(ctx.Ctx(), &model.TalkRecords{
		TalkType:   model.TalkRecordTalkTypeGroup,
		MsgType:    entity.ChatMsgSysGroupNotice,
		UserId:     uid,
		ReceiverId: int(params.GroupId),
		Extra: jsonutil.Encode(model.TalkRecordExtraGroupNotice{
			OwnerId:   uid,
			OwnerName: "gzydong",
			Title:     params.Title,
			Content:   params.Content,
		}),
	})

	return ctx.Success(nil, msg)
}

// Delete 删除群公告
func (c *Notice) Delete(ctx *ichat.Context) error {

	params := &web.GroupNoticeDeleteRequest{}
	if err := ctx.Context.ShouldBindJSON(params); err != nil {
		return ctx.InvalidParams(err)
	}

	if err := c.groupNoticeService.Delete(ctx.Ctx(), int(params.GroupId), int(params.NoticeId)); err != nil {
		return ctx.ErrorBusiness(err.Error())
	}

	return ctx.Success(nil, "群公告删除成功！")
}

// List 获取群公告列表(所有)
func (c *Notice) List(ctx *ichat.Context) error {

	params := &web.GroupNoticeListRequest{}
	if err := ctx.Context.ShouldBindQuery(params); err != nil {
		return ctx.InvalidParams(err)
	}

	// 判断是否是群成员
	if !c.groupMemberService.Dao().IsMember(ctx.Ctx(), int(params.GroupId), ctx.UserId(), true) {
		return ctx.ErrorBusiness("无获取数据权限！")
	}

	all, err := c.groupNoticeService.Dao().GetListAll(ctx.Ctx(), int(params.GroupId))
	if err != nil {
		return ctx.ErrorBusiness(err.Error())
	}

	items := make([]*web.GroupNoticeListResponse_Item, 0)
	for i := 0; i < len(all); i++ {
		items = append(items, &web.GroupNoticeListResponse_Item{
			Id:           int32(all[i].Id),
			Title:        all[i].Title,
			Content:      all[i].Content,
			IsTop:        int32(all[i].IsTop),
			IsConfirm:    int32(all[i].IsConfirm),
			ConfirmUsers: all[i].ConfirmUsers,
			Avatar:       all[i].Avatar,
			CreatorId:    int32(all[i].CreatorId),
			CreatedAt:    timeutil.FormatDatetime(all[i].CreatedAt),
			UpdatedAt:    timeutil.FormatDatetime(all[i].UpdatedAt),
		})
	}

	return ctx.Success(&web.GroupNoticeListResponse{
		Items: items,
	})
}
