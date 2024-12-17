package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Print("hi")

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		msg := scanner.Text()
		handleMessage(msg)
	}
}

func handleMessage(_ any) {

}
