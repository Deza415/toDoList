package main

import (
	"context"
	"log"
	"net"

	"github.com/Deza415/toDoList/proto"
	"google.golang.org/grpc"
)

// INSERT DATABASE HERE
// In-memory todo store
var todos []*proto.Todo
var nextID int32 = 1

// Server struct that implements proto.TodoServiceServer
type server struct {
	proto.UnimplementedTodoServiceServer
}

// Creates a new todo item, copies the title, sets completed to false
// then appends todos to the todos list and returns the created Todo back to the client
func (s *server) CreateTodo(ctx context.Context, req *proto.CreateTodoRequest) (*proto.Todo, error) {
	todo := &proto.Todo{
		Id:        nextID,
		Title:     req.Title,
		Completed: false,
	}
	nextID++
	todos = append(todos, todo)
	return todo, nil
}

// Returns the whole todoList
func (s *server) GetTodos(ctx context.Context, _ *proto.Empty) (*proto.TodoList, error) {
	return &proto.TodoList{Todos: todos}, nil
}

// Loops through the todo list to find matching id and update fields
func (s *server) UpdateTodo(ctx context.Context, req *proto.UpdateTodoRequest) (*proto.Todo, error) {
	for _, t := range todos {
		if t.Id == req.Id {
			t.Title = req.Title
			t.Completed = req.Completed
			return t, nil
		}
	}
	return nil, nil // or return an error
}

// finds todo by it's ID, removes it from list
func (s *server) DeleteTodo(ctx context.Context, req *proto.DeleteTodoRequest) (*proto.Empty, error) {
	for i, t := range todos {
		if t.Id == req.Id {
			todos = append(todos[:i], todos[i+1:]...)
			break
		}
	}
	return &proto.Empty{}, nil
}

func main() {
	//TCP Listener on Port 50051
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	//Register service
	proto.RegisterTodoServiceServer(grpcServer, &server{})

	//start listening for client connections
	log.Println("gRPC server listening on port 50051...")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
