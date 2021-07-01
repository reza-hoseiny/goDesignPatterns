package main

import (
	_ "goDesignPatterns/decoratorpattern/model"
	decorator "goDesignPatterns/decoratorpattern/decorator"
	"fmt"
)

func ShowPlainItem() {
	item := decorator.NewShowItemView("Bag Red Velvet", "Bag for women", 90.1).Show()
 
	fmt.Printf("=====Show Plain Item=====\n")
	fmt.Printf("Item name: %s\n", item.Name)
	fmt.Printf("Item desc: %s\n", item.Description)
	fmt.Printf("Item price: %.2f\n", item.Price)
	fmt.Printf("%v %T\n", item, item)
 }

func main() {
	// item := model.Item{
	// 	Name:        "a model name",
	// 	Description: "a model long description",
	// 	Price:       44.44,
	// };

	// fmt.Println("Hello, Gopher!")
	// fmt.Printf("%v %T\n", item, item)
	ShowPlainItem()
}
