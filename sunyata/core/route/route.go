package route

import (
	"reflect"
	"sync"
)

type ControllerRoute struct {
	Controller interface{}
	Action string
}

type ControllerHandler struct {
	mu              sync.RWMutex
	routeEntry     map[string]*ControllerRoute
}

var DefaultControllerHandler = &defaultControllerHandler

var defaultControllerHandler ControllerHandler
const ContentTypeHeader = "Content-Type"
func MapRoutes(){
	//RegisterRoute("/TestJson", &controller.MainController{}, "TestJson")
	//RegisterRoute("/{s+}/TestRegex", &controller.MainController{}, "TestParam")
}


func RegisterRoute(pattern string, controller interface{}, action string) {
	DefaultControllerHandler.mu.Lock()
	defer DefaultControllerHandler.mu.Unlock()

	if pattern == "" {
		panic("handler: invalid pattern " + pattern)
	}
	if controller == nil {
		panic("controller: nil controller")
	}

	if DefaultControllerHandler.routeEntry == nil {
		DefaultControllerHandler.routeEntry = make(map[string]*ControllerRoute)
	}
	var newRouteEntry = &ControllerRoute{Controller:&controller, Action:action}
	DefaultControllerHandler.routeEntry[pattern] = newRouteEntry
}

/*
func DispatchRoute(w http.ResponseWriter, r *http.Request) {
	var urlPath = r.URL.Path
	var routeEntry = DefaultControllerHandler.routeEntry[urlPath]
	var t = reflect.ValueOf(routeEntry.Controller)
	var actionString = routeEntry.Action
	inputs := make([]reflect.Value, 1)
	inputs[0] = reflect.ValueOf(r)
	var response = t.MethodByName(actionString).Call(inputs)[0].Interface().(re.Response)
	w.Header().Set(ContentTypeHeader, response.ContentType)
	fmt.Fprintf(w, response.ContentString)
}
*/

func Invoke(controller interface{}, name string, args ...interface{}) string {
	inputs := make([]reflect.Value, len(args))
	for i, _ := range args {
		inputs[i] = reflect.ValueOf(args[i])
	}
	var result string = reflect.ValueOf(controller).MethodByName(name).Call(inputs)[0].String()
	return result
}
