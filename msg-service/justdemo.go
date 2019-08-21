package main

import "bx.com/msg-service/workers"

func main() {
	workers.InitManager()
	workers.Start()
}
