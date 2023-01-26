// package test

// import (
// 	"fmt"
// 	"testing"

// 	"github.com/kinbiko/jsonassert"
// 	// "log"
// 	// "net/http/httptest"
// 	// "testing"
// 	// "github.com/stretchr/testify/assert"
// )

// type printer struct{}

// func (p *printer) Errorf(format string, args ...interface{}) {
// 	fmt.Println(fmt.Sprintf(format, args...))
// }

// var t *printer

// func TestHBill(t *testing.T) {
// 	ja := jsonassert.New(t)
// 	expTestScore := "46%"
// 	var2 := 34
// 	ja.Assertf(
// 		`{ "name": "Jayne Cobb", "age": 36, "averageTestScore": "88%" }`,
// 		`{ "name": "Jayne Cobb", "age": 36, "averageTestScore": "%s" }`, expTestScore,
// 		`{ "name": "Jayne Cobb", "age": 36, "averageTestScore": "%v" }`, var2,
// 	)
// }
