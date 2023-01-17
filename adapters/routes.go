package myPack1

import "github.com/gin-gonic/gin"

func (s *Server) routes() {
	s.router = gin.New()
	core := s.router.Group("/")
	core.Use()
	{
		core.GET("/hola", s.handlerHello)
	}
}
