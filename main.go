package main

func main() {
	todos := Todos{}
	todos.add("Buy milk")
	todos.add("Buy eggs")
	todos.add("Buy sausage")
	todos.toggle(0)
	todos.print()
}
