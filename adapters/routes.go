package myPack1

import "github.com/gin-gonic/gin"

func (s *Server) routes() {
	s.router = gin.New()
	core := s.router.Group("/")
	core.Use()
	{
		core.GET("/hola", s.handlerHello)
		core.POST("/add-bills/:name", s.handlerAddBills)
		core.GET("/buscar/:name", s.handlerBuscar)
		core.POST("/generate", s.handlerGenerate)
	}

}
