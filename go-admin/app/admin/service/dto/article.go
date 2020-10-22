package dto

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"time"

	"go-admin/app/admin/models"
	"go-admin/common/dto"
	"go-admin/common/log"
	common "go-admin/common/models"
	"go-admin/tools"
)

type ArticleSearch struct {
	dto.Pagination     `search:"-"`
    ID uint `form:"ID" search:"type:exact;column:id;table:article" comment:"编码"`

    Title string `form:"title" search:"type:exact;column:title;table:article" comment:"标题"`

    Author string `form:"author" search:"type:exact;column:author;table:article" comment:"作者"`

    
}

func (m *ArticleSearch) GetNeedSearch() interface{} {
	return *m
}

func (m *ArticleSearch) Bind(ctx *gin.Context) error {
    msgID := tools.GenerateMsgIDFromContext(ctx)
    err := ctx.ShouldBind(m)
    if err != nil {
    	log.Debugf("MsgID[%s] ShouldBind error: %s", msgID, err.Error())
    }
    return err
}

func (m *ArticleSearch) Generate() dto.Index {
	o := *m
	return &o
}

type ArticleControl struct {
    
    ID uint `uri:"ID" comment:"编码"` // 编码

    Title string `json:"title" comment:"标题"`
    

    Author string `json:"author" comment:"作者"`
    

    Content string `json:"content" comment:"内容"`
    

    Status string `json:"status" comment:"状态"`
    

    PublishAt time.Time `json:"publishAt" comment:"发布时间"`
    
}

func (s *ArticleControl) Bind(ctx *gin.Context) error {
    msgID := tools.GenerateMsgIDFromContext(ctx)
    err := ctx.ShouldBindUri(s)
    if err != nil {
        log.Debugf("MsgID[%s] ShouldBindUri error: %s", msgID, err.Error())
        return err
    }
    err = ctx.ShouldBind(s)
    if err != nil {
        log.Debugf("MsgID[%s] ShouldBind error: %#v", msgID, err.Error())
    }
    return err
}

func (s *ArticleControl) Generate() dto.Control {
	cp := *s
	return &cp
}

func (s *ArticleControl) GenerateM() (common.ActiveRecord, error) {
	return &models.Article{
	
        Model:     gorm.Model{ID: s.ID},
        Title:  s.Title,
        Author:  s.Author,
        Content:  s.Content,
        Status:  s.Status,
        PublishAt:  s.PublishAt,
	}, nil
}

func (s *ArticleControl) GetId() interface{} {
	return s.ID
}

type ArticleById struct {
	dto.ObjectById
}

func (s *ArticleById) Generate() dto.Control {
	cp := *s
	return &cp
}

func (s *ArticleById) GenerateM() (common.ActiveRecord, error) {
	return &models.Article{}, nil
}
