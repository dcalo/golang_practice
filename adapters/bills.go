package myPack1

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type bill struct {
	Name string `json:"name"`
}

var bill1 = bill{Name: "Pepito"}

func (s *Server) handlerHello(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, bill1)
}
