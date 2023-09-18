package main

import (
	"log"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("not enough arguments! Must be 2 or more")

		return
	}

	envDir := os.Args[1]

	env, err := ReadDir(envDir)
	if err != nil {
		log.Fatalf("read dir %v", err)

		return
	}

	returnCode := RunCmd(os.Args[2:], env)

	os.Exit(returnCode)
}
