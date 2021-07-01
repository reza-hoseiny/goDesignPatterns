package main

import (
	_ "goDesignPatterns/decoratorpattern/model"
	_ "goDesignPatterns/decoratorpattern/decorator"
	services "goDesignPatterns/decoratorpattern/services"
	"fmt"
)



func main() {
	fmt.Println("Hello, Gopher!")
	

	// 
	// fmt.Printf("%v %T\n", item, item)
	services.TestService()
}
