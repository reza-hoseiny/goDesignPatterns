package main

import (
	_ "goDesignPatterns/decoratorpattern/model"
	_ "goDesignPatterns/decoratorpattern/decorator"
	services "goDesignPatterns/decoratorpattern/services"
	// "fmt"
)



func main() {
	// item := model.Item{
	// 	Name:        "a model name",
	// 	Description: "a model long description",
	// 	Price:       44.44,
	// };

	// fmt.Println("Hello, Gopher!")
	// fmt.Printf("%v %T\n", item, item)
	services.ShowPlainItem()
}
