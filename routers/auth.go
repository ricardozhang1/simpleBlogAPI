package routers

import (
	"database/sql"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
	"net/http"
	models "simple_blog/models/sqlc"
	"time"
)

type createUserRequest struct {
	UserName	string	`json:"username" binding:"required,alphanum"`
	Name		string	`json:"name" binding:"required,min=6"`
	Password	string	`json:"password" binding:"required"`
	Email		string	`json:"email" binding:"required,email"`
}

type userResponse struct {
	ID			int32		`json:"id"`
	UserName	string		`json:"username"`
	Name		string		`json:"name"`
	Email		string		`json:"email"`
	CreatedAt	time.Time	`json:"created_at"`
}

func newUserResponse(user models.BlogUser) userResponse {
	return userResponse{
		ID: user.ID,
		UserName: user.Username,
		Name: user.Name,
		Email: user.Email,
		CreatedAt: user.CreatedAt,
	}
}

func (server *Server) createUser(ctx *gin.Context) {
	var req createUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := models.CreateAuthorParams{
		Email: req.Email,
		Username: req.UserName,
		Name: req.Name,
		Password: req.Password,
	}

	user, err := server.query.CreateAuthor(ctx, arg)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			switch pqErr.Code.Name() {
			case "unique_violation":
				ctx.JSON(http.StatusForbidden, errorResponse(err))
				return
			}
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	rsp := newUserResponse(user)
	ctx.JSON(http.StatusOK, rsp)
}

type loginRequest struct {
	Email		string	`json:"email" binding:"required,email"`
	Password	string	`json:"password" binding:"required"`
}

type loginResponse struct {
	ID			int32	`json:"id"`
	Email 		string	`json:"email"`
	Name		string	`json:"name"`
	CreatedAt	time.Time `json:"created_at"`
}

func (server *Server) loginUser(ctx *gin.Context) {
	var req loginRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	user, err := server.query.GetAuthor(ctx, req.Email)
	if err != nil {
		if err != sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	if user.Password != req.Password {
		err = errors.New("password is not equal")
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}

	resp := loginResponse{
		ID: user.ID,
		Email: user.Email,
		Name: user.Name,
		CreatedAt: user.CreatedAt,
	}
	ctx.JSON(http.StatusOK, resp)
}





