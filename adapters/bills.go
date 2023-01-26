package myPack1

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"reflect"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"

	my_conn "example.com/myProject/internal/config"
)

type Bill struct {
	Name     string    `json:"name"`
	Price    float32   `json:"price"`
	Paid     bool      `json:"paid"`
	DatePaid time.Time `json:"date-paid"`
}

var bill1 = Bill{Name: "Pepito"}
var bills = []Bill{
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
	bills = append(bills, Bill{Name: name})
	log.Println(bills)
	for _, bill := range bills {
		if bill.Name == name {
			log.Println("encontrado")
			bill.DatePaid = time.Now()
			c.IndentedJSON(http.StatusOK, bill)

		}
	}
}

func (s *Server) handlerGenerate(c *gin.Context) {
	var requestBody Bill

	if err := c.BindJSON(&requestBody); err != nil {
		// DO SOMETHING WITH THE ERROR
		log.Fatal("Body not valid")
	}

	log.Println("Param value: ", requestBody.Name)
	bills = append(bills, requestBody)
	log.Println("new array: ", bills)

	// marshal
	marshal_struct, err := json.Marshal(Bill{Name: "Dario", Price: 10.50, DatePaid: time.Now()})
	if err != nil {
		log.Println("I can not marshal the passed data")
	}
	log.Println("Marshalled data: ", marshal_struct)
	// unmarshal
	json_example := `{"name": "Ana"}`
	json_in_bytes := []byte(json_example)

	numero := `1`
	convert, err := strconv.Atoi(numero)
	if err != nil {
		log.Println("I can not transform passed data")
	}

	log.Println("numero convertido:", convert)

	var unmarshalled_bill Bill
	json.Unmarshal(json_in_bytes, &unmarshalled_bill)

	log.Println("after unmarshall: ", unmarshalled_bill)

	valor1, valor2, valor3 := procesar1(5)
	log.Println("respuestas: ", valor1, valor2, valor3)

	valor4, valor5, valor6 := procesar1(10)
	log.Println("respuestas: ", valor4, valor5, valor6)

	//validacion
	x := T{5}
	log.Println(reflect.TypeOf(x))
	x.a()
	x.b()
	log.Println("Final value:", &x.val)

	// probar conexion de bd
	my_conn.Start_connection()

	// response
	c.IndentedJSON(http.StatusOK, bills)
}

func procesar1(a int) (r1 int, r2 float32, err error) {

	if a == 10 {
		return 0, 0, errors.New("error 1")
	}
	return a * 3, float32(a), nil
}

type T struct {
	val int
}

type A struct {
	val  int
	val2 bool
}

func (p *T) a() {
	log.Println("type:", reflect.TypeOf(p))
	log.Println("value of val before:", p.val)
	p.val += 1

	log.Println("value of val after:", p.val)
}

func (p T) b() {
	log.Println(reflect.TypeOf(p))
	log.Println("value of val before:", p.val)
	p.val += 1

	log.Println("value of val after:", p.val)
}
