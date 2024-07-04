package main

import (
	"fmt"
	"io"
	"os"

	"mon.example/evaluator"
	"mon.example/lexer"
	"mon.example/object"
	"mon.example/parser"
)

// "fmt"
// "os"
// "os/user"

// "mon.example/repl"


func main() {
	// user, err := user.Current()
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Printf("Hello %s, welcome to the langugae\n", user.Username)
	// fmt.Println("feel free to type in command")
	// repl.Start(os.Stdin, os.Stdout)
	byt, err := os.ReadFile("./hello.gs")
	if err != nil {
		panic(fmt.Sprintf("unable to read file: %+v", err))
	}
	l := lexer.New(string(byt))
	p := parser.New(l)
	program := p.ParseProgram()
	env := object.NewEnvironment()
	evaluated := evaluator.Eval(program, env)
	if evaluated != nil {
		io.WriteString(os.Stdout, evaluated.Inspect())
		io.WriteString(os.Stdout, "\n")
	}
}
