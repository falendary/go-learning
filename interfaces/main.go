package main

import (
	"fmt"
	"os"
)

//That means:
//
//Any type that has a Greet() string method automatically satisfies this interface.
//
//There is no implements keyword.
//
//No inheritance.
//
//Just behavior matching.

type Greeter interface {
	Greet() string
}

/*
Logger defines behavior, not implementation.

Any type that has a Log(string) method
automatically satisfies this interface.
*/
type Logger interface {
	Log(msg string) error
}
type ConsoleLogger struct{}

func (c ConsoleLogger) Log(msg string) error {
	_, err := fmt.Println(msg)

	return err
}

/*
FileLogger writes messages into a file.
*/
type FileLogger struct {
	file *os.File
}

/*
Constructor for FileLogger.
*/
func NewFileLogger(path string) (*FileLogger, error) {
	f, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return nil, err
	}
	return &FileLogger{file: f}, nil
}

func (f *FileLogger) Log(msg string) error {
	_, err := f.file.WriteString(msg + "\n")
	return err
}

func (f *FileLogger) Close() error {
	return f.file.Close()
}

type OrderService struct {
	logger Logger
}

func NewOrderService(logger Logger) *OrderService {
	return &OrderService{logger: logger}
}

func (s *OrderService) CreateOrder(id string) error {
	return s.logger.Log("creating order: " + id)
}

func main() {

	// --- Use ConsoleLogger ---
	consoleLogger := ConsoleLogger{}
	consoleService := NewOrderService(consoleLogger)

	consoleService.CreateOrder("A100")

	// --- Use FileLogger ---
	fileLogger, err := NewFileLogger("orders.log")
	if err != nil {
		panic(err)
	}
	defer fileLogger.Close()

	fileService := NewOrderService(fileLogger)
	fileService.CreateOrder("B200")

	fmt.Println("Done")
}
