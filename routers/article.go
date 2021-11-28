package routers

import (
	"github.com/gin-gonic/gin"
)

type createArticleRequest struct {
	TagID     int64  `json:"tag_id" binding:"required"`
	Title     string `json:"title" binding:"required"`
	Intro     string `json:"intro" binding:"required"`
	Content   string `json:"content" binding:"required"`
}

func (server *Server) createArticle(ctx *gin.Context) {

}
