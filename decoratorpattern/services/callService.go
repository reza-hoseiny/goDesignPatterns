package services

import (
	model "goDesignPatterns/decoratorpattern/model"
)



 func TestService() {
	item1 := model.Item{
		Name:        "a model name",
		Description: "a model long description",
		Price:       44.44,
	}

	item2 :=  model.Item{
		Name: "Bag Red Velvet", 
		Description: 		"Bag for women", 
		Price: 90.1,
	}

	showPlainItem(item1)
	showPlainItem(item2)

 }
