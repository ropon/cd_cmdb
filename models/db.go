package models

import (
	"github.com/ropon/cd_cmdb/utils"
	"time"
)

type Db struct {
	Id            uint      `json:"id" form:"id" gorm:"primary_key,AUTO_INCREMENT"`
	UUID          string    `json:"uuid" form:"uuid" gorm:"column:uuid" sql:"unique;not null"`
	SvcType       string    `json:"svc_type" form:"svc_type" gorm:"column:svc_type" sql:"not null"`
	AuthorEmail   string    `json:"author_email" form:"author_email" gorm:"column:author_email" sql:"not null"`
	Ports         string    `json:"ports" form:"ports" gorm:"column:ports" sql:"unique;not null"`
	CreateTimeStr string    `json:"create_time" gorm:"-"`
	UpdateTimeStr string    `json:"update_time" gorm:"-"`
	CreateTime    time.Time `json:"-" gorm:"column:create_time" sql:"type:datetime"`
	UpdateTime    time.Time `json:"-" gorm:"column:update_time" sql:"type:datetime"`
}

type DbList []*Db

func (s *Db) TableName() string {
	return "db"
}

// FormatTime 特殊处理时间
func (s *Db) FormatTime() {
	s.CreateTimeStr = utils.FormatTime(s.CreateTime)
	s.UpdateTimeStr = utils.FormatTime(s.UpdateTime)
}
