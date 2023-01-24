package myPack1

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type bill struct {
	Name     string    `json:"name"`
	Price    float32   `json:"price"`
	Paid     bool      `json:"paid"`
	DatePaid time.Time `json:"date-paid"`
}

var bill1 = bill{Name: "Pepito"}
var bills = []bill{
	{Name: "juan"},
	{Name: "alex"},
	{Name: "luis"},
}

func (s *Server) handlerHello(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, bill1)
}
func (s *Server) handlerAddBills(c *gin.Context) {
	name := c.Param("name")
	log.Println(name)
	//c.IndentedJSON(http.StatusOK, bills)
}

func (s *Server) handlerBuscar(c *gin.Context) {
	//var errorParam error
	name := c.Param("name")
	if name == "" {
		log.Println("no valido")
		c.IndentedJSON(http.StatusBadRequest, "no valido")
	}
	log.Println(bills)
	bills = append(bills, bill{Name: name})
	log.Println(bills)
	for _, bill := range bills {
		if bill.Name == name {
			log.Println("encontrado")
			bill.DatePaid = time.Now()
			c.IndentedJSON(http.StatusOK, bill)

		}

	}
}
