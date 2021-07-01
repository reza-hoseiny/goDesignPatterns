package services

import (
	model "goDesignPatterns/decoratorpattern/model"
	decorator "goDesignPatterns/decoratorpattern/decorator"
	"fmt"
)

func showPlainItem(input_item model.Item) {
	item := decorator.NewShowItemViewFromPlainItem(input_item).Show()
 
	fmt.Printf("=====Show Plain Item=====\n")
	fmt.Printf("Item name: %s\n", item.Name)
	fmt.Printf("Item desc: %s\n", item.Description)
	fmt.Printf("Item price: %.2f\n", item.Price)
	fmt.Printf("%v %T\n", item, item)
 }
