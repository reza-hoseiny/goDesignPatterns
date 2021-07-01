package decorators

import (
	model "goDesignPatterns/decoratorpattern/model"
)


type ShowItem interface {
	Show() model.Item
}