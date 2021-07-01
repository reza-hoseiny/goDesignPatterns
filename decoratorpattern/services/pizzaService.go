package services

import(
	decorator "goDesignPatterns/decoratorpattern/decorator"
	components "goDesignPatterns/decoratorpattern/components"
	"fmt"
)

func MakePizza(){

	// fmt.Println("Hello, Pizza services!")

	pizza := &components.VeggeMania{}

    //Add cheese topping
    pizzaWithCheese := &decorator.CheeseTopping{
        Pizza: pizza,
    }
	price := pizzaWithCheese.GetPrice()

	fmt.Printf("Price of veggeMania with cheese topping is  %v %T\n", price, price)
	//Add tomato topping
    pizzaWithCheeseAndTomato := &decorator.TomatoTopping{
        Pizza: pizzaWithCheese,
    }

	price = pizzaWithCheeseAndTomato.GetPrice()

	fmt.Printf("Price of veggeMania with tomato and cheese topping is %v %T\n", price, price)

	

}