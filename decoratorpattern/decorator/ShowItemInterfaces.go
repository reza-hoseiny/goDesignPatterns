package decorator

import (
	model "goDesignPatterns/decoratorpattern/model"
)

type ShowItem interface {
	Show() model.Item
}