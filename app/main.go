package main

import (
	"bufio"
	"fmt"
	"os"
)

// Ensures gofmt doesn't remove the "fmt" import in stage 1 (feel free to remove this!)
var _ = fmt.Print

func main() {
	// TODO: Uncomment the code below to pass the first stage
	for {
		fmt.Print("$ ")
		command, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error reading input:", err)
			os.Exit(1)
		}
		command = command[:len(command)-1]
		if command == "exit" {
			return
		}
		i := 0
		for ; i < len(command); i++ {
			if command[i] == ' ' {
				break
			}
		}
		if i == len(command) {
			if command == "echo" {
				fmt.Println()
			} else if command != "exit" {
				fmt.Printf("%s: command not found\n", command)
			}
			continue
		}
		str := command[i+1:]
		command = command[:i]

		if command == "echo" {
			fmt.Println(str)
		}
		if command == "type" {
			if str == "echo" || str == "type" || str == "exit" {
				fmt.Printf("%s is a shell builtin\n", str)
			} else {
				fmt.Printf("%s: not found\n", str)
			}
		}
		if command == "exit" {
			return
		}
		if command != "exit" && command != "echo" && command != "type" {
			fmt.Printf("%s: command not found\n", command)
		}
	}
}
