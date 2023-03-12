package repo

import (
	"github.com/captainherolee/clean-arch-go/internal/pkg/notice"
)

type entityMapper struct{}

func (e entityMapper) toNoticeEntity(dbNotice *Notice) *notice.Notice {
	return &notice.Notice{
		ID:        dbNotice.ID,
		Title:     dbNotice.Title,
		Content:   dbNotice.Content,
		UserTypes: dbNotice.UserTypes,
		Lang:      dbNotice.Lang,
		CreatedAt: dbNotice.CreatedAt,
		UpdatedAt: dbNotice.UpdatedAt,
	}
}
