# CLI Todo List

This is a small project where you can create your own to do list in the terminal.

*PS: This is just a way for me to learn the go language.*

## There is a few commands you need to know to utilize the todolist

To print out your current to do list
````bash
    go run ./ -list
````

To add a new task to your list
````bash
    go run ./ -add "Your task"
````

To edit a specific task
````bash
    go run ./ -edit "id: Your new task"
````

To toggle a specific task
````bash
    go run ./ -toggle id
````

To delete a specific task
````bash
    go run ./ -delete id
````
