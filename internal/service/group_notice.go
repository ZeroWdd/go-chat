package service

import (
	"context"
	"time"

	"go-chat/internal/pkg/timeutil"
	"go-chat/internal/repository/model"
	"go-chat/internal/repository/repo"
)

type GroupNoticeService struct {
	*repo.Source
	notice *repo.GroupNotice
}

func NewGroupNoticeService(source *repo.Source, notice *repo.GroupNotice) *GroupNoticeService {
	return &GroupNoticeService{
		Source: source,
		notice: notice,
	}
}

func (s *GroupNoticeService) Dao() *repo.GroupNotice {
	return s.notice
}

type GroupNoticeEditOpt struct {
	UserId    int
	GroupId   int
	NoticeId  int
	Title     string
	Content   string
	IsTop     int
	IsConfirm int
}

// Create 创建群公告
func (s *GroupNoticeService) Create(ctx context.Context, opt *GroupNoticeEditOpt) error {
	return s.notice.Create(ctx, &model.GroupNotice{
		GroupId:      opt.GroupId,
		CreatorId:    opt.UserId,
		Title:        opt.Title,
		Content:      opt.Content,
		IsTop:        opt.IsTop,
		IsConfirm:    opt.IsConfirm,
		ConfirmUsers: "{}",
	})
}

// Update 更新群公告
func (s *GroupNoticeService) Update(ctx context.Context, opt *GroupNoticeEditOpt) error {
	_, err := s.notice.UpdateWhere(ctx, map[string]any{
		"title":      opt.Title,
		"content":    opt.Content,
		"is_top":     opt.IsTop,
		"is_confirm": opt.IsConfirm,
		"updated_at": time.Now(),
	}, "id = ? and group_id = ?", opt.NoticeId, opt.GroupId)
	return err
}

func (s *GroupNoticeService) Delete(ctx context.Context, groupId, noticeId int) error {
	_, err := s.notice.UpdateWhere(ctx, map[string]any{
		"is_delete":  1,
		"deleted_at": timeutil.DateTime(),
		"updated_at": time.Now(),
	}, "id = ? and group_id = ?", noticeId, groupId)
	return err
}
