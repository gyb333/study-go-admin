package models

import (
    "gorm.io/gorm"
	"time"

	"go-admin/common/models"
)

type Article struct {
    gorm.Model
    models.ControlBy
    
    Title string `json:"title" gorm:"type:varchar(128);comment:标题"` // 
    Author string `json:"author" gorm:"type:varchar(128);comment:作者"` // 
    Content string `json:"content" gorm:"type:varchar(255);comment:内容"` // 
    Status string `json:"status" gorm:"type:int(1);comment:状态"` // 
    PublishAt time.Time `json:"publishAt" gorm:"type:timestamp;comment:发布时间"` // 
}

func (Article) TableName() string {
    return "article"
}

func (e *Article) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *Article) GetId() interface{} {
	return e.ID
}
