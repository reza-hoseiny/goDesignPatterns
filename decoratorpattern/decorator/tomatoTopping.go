package decorator

import(
	interfaces "goDesignPatterns/decoratorpattern/interfaces"
)

type TomatoTopping struct { //
   Pizza interfaces.Pizza
}

func (c *TomatoTopping) GetPrice() int { //tomatoTopping is still a Pizaa
    pizzaPrice := c.Pizza.GetPrice()
    return pizzaPrice + 4
}