package main

import "github.com/SanyaWarvar/todo-app"

func main() {
	srv := new(todo.Server)
	srv.Run("8000")
}
