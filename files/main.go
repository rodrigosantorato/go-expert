package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	f, err := os.Create("ilovego.txt")
	if err != nil {
		panic(err)
	}
	_, err = f.Write([]byte("The Go programming language is an open source project to make programmers more productive.\n\nGo is expressive, concise, clean, and efficient. Its concurrency mechanisms make it easy to write programs that get the most out of multicore and networked machines, while its novel type system enables flexible and modular program construction. Go compiles quickly to machine code yet has the convenience of garbage collection and the power of run-time reflection. It's a fast, statically typed, compiled language that feels like a dynamically typed, interpreted language."))
	if err != nil {
		panic(err)
	}

	content, err := os.ReadFile("ilovego.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(content))

	bigFile, err := os.Open("ilovego.txt")
	if err != nil {
		panic(err)
	}
	reader := bufio.NewReader(bigFile)
	buffer := make([]byte, 10)
	for {
		n, err := reader.Read(buffer)
		if err != nil {
			break
		}
		fmt.Println(string(buffer[:n]))
	}
}
