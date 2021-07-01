package decorators

import (
	model "goDesignPatterns/decoratorpattern/model"
)

type ShowItemView struct {
	item model.Item
}

func NewShowItemView(name, description string, price float64) ShowItemView {
	return ShowItemView{
		item: model.Item{
			Name:        name,
			Description: description,
			Price:       price,
		},
	}
}

func (i ShowItemView) Show() model.Item {
	return i.item
}