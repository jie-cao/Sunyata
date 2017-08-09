package handler

import (
	"fmt"
	"net/http"
	"reflect"
	"sunyata/controller"
	"sync"
)

type BaseHandler struct {
	mu          sync.RWMutex
	actionEntry map[string]string
}

func DispatchRoute(w http.ResponseWriter, r *http.Request) {
	var baseHandler = controller.MainController{}
	var t = reflect.ValueOf(&baseHandler)
	inputs := make([]reflect.Value, 0)
	var result string = t.MethodByName("Index").Call(inputs)[0].String()
	fmt.Println(result)
	fmt.Fprintf(w, result) //这个写入到w的是输出到客户端的
}

func Invoke(controller interface{}, name string, args ...interface{}) string {
	inputs := make([]reflect.Value, len(args))
	for i, _ := range args {
		inputs[i] = reflect.ValueOf(args[i])
	}
	var result string = reflect.ValueOf(controller).MethodByName(name).Call(inputs)[0].String()
	return result
}
