package main

func main() {
	todos := Todos{}
	store := NewStore[Todos]("todos.json")
	store.Load(&todos)
	cmdFlags := NewCmdFlags()
	cmdFlags.Execute(&todos)
	store.Save(todos)
}
