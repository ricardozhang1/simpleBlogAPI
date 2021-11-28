package routers

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
	"net/http"
	models "simple_blog/models/sqlc"
)

type createTagRequest struct {
	Name      string `json:"name"`
	CreatedID int32  `json:"created_id"`
}

func (server *Server) createTag(ctx *gin.Context) {
	var req createTagRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := models.CreateTagParams{
		Name: req.Name,
		CreatedID: req.CreatedID,
	}

	tag, err := server.query.CreateTag(ctx, arg)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			switch pqErr.Code.Name() {
			case "unique_violation":
				ctx.JSON(http.StatusForbidden, errorResponse(err))
			}
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, tag)
}

type getTagRequest struct {
	TagID	int32	`json:"tag_id" binding:"required"`
}

func (server *Server) getTag(ctx *gin.Context) {
	var req getTagRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	tag, err := server.query.GetTag(ctx, req.TagID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusBadRequest, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, tag)
}

type listTagRequest struct {

}

func (server *Server) listTag(ctx *gin.Context) {

}
