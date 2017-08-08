package main

import "fmt"
import "sunyata/core/web"

func main() {
	fmt.Println("Hello, 世界")
	fmt.Println("开始")
	var newRoute = web.Route{Controller:"Main", Action:"Index"}
	fmt.Println(newRoute.Controller)
}
