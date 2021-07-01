package decorator

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

func NewShowItemViewFromPlainItem(input_item model.Item) ShowItemView {
	return ShowItemView{
		item: model.Item{
			Name:        input_item.Name,
			Description: input_item.Description,
			Price:       input_item.Price,
		},
	}
}

func (i ShowItemView) Show() model.Item {
	return i.item
}