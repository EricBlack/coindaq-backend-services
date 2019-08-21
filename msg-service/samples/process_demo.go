package main

import (
	"os"
	"log"
	"os/exec"
)

func main() {
	pid := os.Getegid()
	log.Printf("Main program: %d", pid)

	cmd1 := exec.Command("go", "run", "run_process.go", "test1")
	cmd1Out, err := cmd1.Output()
	if err != nil {
		panic(err)
	}
	log.Printf("Output: %s", cmd1Out)

	cmd2 := exec.Command("go", "run", "run_process.go", "test2")
	cmd2Out, err := cmd2.Output()
	if err != nil {
		panic(err)
	}
	log.Printf("Output: %s", cmd2Out)

	for i :=1; i>0; i++ {
	}
}
