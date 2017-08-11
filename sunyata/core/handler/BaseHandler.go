package handler

import (
	"fmt"
	"net/http"
	"reflect"
	"sunyata/core/action"
	"sync"
)

type ControllerHandler struct {
	mu              sync.RWMutex
	controllerEntry map[string]interface{}
	actionEntry     map[string]string
}

var DefaultControllerHandler = &defaultControllerHandler

var defaultControllerHandler ControllerHandler

func RegisterRoute(pattern string, controller interface{}, action string) {
	DefaultControllerHandler.mu.Lock()
	defer DefaultControllerHandler.mu.Unlock()

	if pattern == "" {
		panic("handler: invalid pattern " + pattern)
	}
	if controller == nil {
		panic("controller: nil controller")
	}

	if DefaultControllerHandler.controllerEntry == nil {
		DefaultControllerHandler.controllerEntry = make(map[string]interface{})
		DefaultControllerHandler.actionEntry = make(map[string]string)
	}
	DefaultControllerHandler.controllerEntry[pattern] = controller
	DefaultControllerHandler.actionEntry[pattern] = action
}

func DispatchRoute(w http.ResponseWriter, r *http.Request) {
	var urlPath = r.URL.Path
	var controller = DefaultControllerHandler.controllerEntry[urlPath]
	var actionString = DefaultControllerHandler.actionEntry[urlPath]
	var t = reflect.ValueOf(controller)
	inputs := make([]reflect.Value, 1)
	inputs[0] = reflect.ValueOf(r)
	var response = t.MethodByName(actionString).Call(inputs)[0].Interface().(action.Response)
	fmt.Fprintf(w, response.ContentString)
}

func Invoke(controller interface{}, name string, args ...interface{}) string {
	inputs := make([]reflect.Value, len(args))
	for i, _ := range args {
		inputs[i] = reflect.ValueOf(args[i])
	}
	var result string = reflect.ValueOf(controller).MethodByName(name).Call(inputs)[0].String()
	return result
}
