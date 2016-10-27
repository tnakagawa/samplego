// cmd project main.go
package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

func main() {
	var args []string
	console := bufio.NewScanner(os.Stdin)
	fmt.Print("# ")
	for console.Scan() {
		args = regexp.MustCompile("(\"[^\"]+\"|[^\\s\"]+)").FindAllString(console.Text(), -1)
		if len(args) > 0 {
			if args[0] == "exit" {
				fmt.Println("bye!")
				break
			} else if args[0] == "echo" {
				if len(args) > 0 {
					fmt.Printf("%s\n", args[1:])
				} else {
					fmt.Println()
				}
			} else {
				fmt.Printf("unknown command: %s %s\n", args[0], args[1:])
			}
		}
		fmt.Print("# ")
	}
}
