package main

import (
	"os"
	"log"
)

func main(){
	pid := os.Getegid()
	log.Printf("Sub program: %d", pid)

	if len(os.Args) >0 {
		for i :=1; i>0; i++ {
		}
	}
}
