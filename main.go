package main

import (
	// _ "goDesignPatterns/decoratorpattern/model"
	// _ "goDesignPatterns/decoratorpattern/decorator"
	"fmt"
	services "goDesignPatterns/decoratorpattern/services"
)

func main() {
	fmt.Println("Hello, Gopher!")
	// services.TestService()
	services.MakePizza()
}
