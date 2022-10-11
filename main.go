package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"

	"go.neonxp.dev/jsonrpc2/rpc"
	"go.neonxp.dev/jsonrpc2/transport"
)

var todols = make(map[int]*Todo)

type Todo struct {
	Title string `json:"title"`
	Body  string `json:"body"`
	Done  bool   `json:"done"`
}

func main() {
	server := rpc.New(
		rpc.WithLogger(rpc.StdLogger),
		rpc.WithTransport(&transport.HTTP{Bind: ":8080"}),
	)

	server.Register("addTodo", rpc.H(AddTodo))
	server.Register("deleteTodo", rpc.H(DeleteTodo))
	server.Register("updateTodo", rpc.H(UpdateTodo))

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill)
	defer cancel()

	if err := server.Run(ctx); err != nil {
		log.Fatal(err)
	}
}

func AddTodo(ctx context.Context, todo *Todo) (*Todo, error) {
	fmt.Println(ctx)
	counter := 10

	todols = map[int]*Todo{
		counter: todo,
	}

	fmt.Println(todols)
	return todols[counter], nil
}

func DeleteTodo(ctx context.Context, todo *Todo) (string, error) {
	return "", nil
}

func UpdateTodo(ctx context.Context, todo *Todo) (*Todo, error) {
	return todo, nil
}
