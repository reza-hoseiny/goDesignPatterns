package main

import (
	model "goDesignPatterns/decoratorpattern/model"
	"fmt"
)

func main() {
	item := model.Item{
		Name:        "name",
		Description: "description",
		Price:       44.44,
	};

	fmt.Println("Hello, Gopher!")
	fmt.Printf("%v %T\n", item, item)
}
