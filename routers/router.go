package routers

import (
	"github.com/gin-gonic/gin"
	models "simple_blog/models/sqlc"
	"simple_blog/pkg/setting"
)

type Server struct {
	config setting.Setting
	query *models.Queries
	router *gin.Engine
}

func NewServer(config setting.Setting, q *models.Queries) (*Server, error) {
	// 权限验证的待处理

	server := &Server{
		config: config,
		query: q,
	}

	// 可以在这里注册验证器，对request的验证

	server.initRouter()
	return server, nil
}

func (server *Server) initRouter() {
	// 配置路由
	router := gin.Default()

	router.POST("/user", server.createUser)
	router.POST("/user/login", server.loginUser)

	router.POST("/article", server.createArticle)
	//router.GET("/article/:id", server.getArticle)
	//router.GET("/article", server.listArticle)

	router.POST("/tag", server.createTag)
	router.GET("/tag", server.getTag)
	router.GET("/tags", server.listTag)

	server.router = router
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

// errorResponse 统一处理错误信息返回的格式，方便阅读，简化代码
func errorResponse(err error) gin.H {
	return gin.H{"message": err.Error()}
}

