package main

import (
	"fmt"
)

func main() {
	todos := Todos{}
	todos.add("Buy milk")
	todos.add("Buy eggs")
	todos.add("Buy sausage")
	fmt.Printf("%+v\n\n", todos)
	todos.delete(0)
	fmt.Printf("%+v", todos)

	todos.toggle(2)
	fmt.Printf("%v", todos)
}
