package router

import (
    "github.com/gin-gonic/gin"

    "go-admin/app/admin/middleware"
    "go-admin/app/admin/models"
    "go-admin/app/admin/service/dto"
    "go-admin/common/actions"
    jwt "go-admin/pkg/jwtauth"
)

func init()  {
	routerCheckRole = append(routerCheckRole, registerArticleRouter)
}

// 需认证的路由代码
func registerArticleRouter(v1 *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) {
    r := v1.Group("/article").Use(authMiddleware.MiddlewareFunc()).Use(middleware.AuthCheckRole())
    {
        model := &models.Article{}
        r.GET("", actions.PermissionAction(), actions.IndexAction(model, new(dto.ArticleSearch), func() interface{} {
            list := make([]models.Article, 0)
            return &list
        }))
        r.GET("/:id", actions.PermissionAction(), actions.ViewAction(new(dto.ArticleById), nil))
        r.POST("", actions.CreateAction(new(dto.ArticleControl)))
        r.PUT("/:id", actions.PermissionAction(), actions.UpdateAction(new(dto.ArticleControl)))
        r.DELETE("", actions.PermissionAction(), actions.DeleteAction(new(dto.ArticleById)))
    }
}
