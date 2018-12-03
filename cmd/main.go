package main

import (
	"fmt"
	"github.com/unused/gockethook"
	"os"
	"strings"
)

func usage() {
	fmt.Println("Usage: ROCKET_HOOK=http://<domain>/hooks/<token> gocket message")
}

func main() {
	hookUrl := os.Getenv("ROCKET_HOOK")

	if len(hookUrl) == 0 {
		fmt.Fprintln(os.Stderr, "hook url is missing")
		usage()
		os.Exit(1)
	}

	message := strings.Trim(strings.Join(os.Args[1:], " "), " ")

	if len(message) == 0 {
		fmt.Fprintln(os.Stderr, "Message is empty")
		usage()
		os.Exit(1)
	}

	gockethook.New(hookUrl).SendText(message)
}
