package main

import (
	"net/http"

	"github.com/plkumar/learn-go/controllers"
)

func main() {
	//fmt.Println("Hello World")
	controllers.RegisterControllers()
	http.ListenAndServe(":3000", nil)
}
