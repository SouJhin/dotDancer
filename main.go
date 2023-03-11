package main

import (
	"dancer/cmd"
)

func main() {
	cmd.Start()
	defer cmd.Clean()
}
