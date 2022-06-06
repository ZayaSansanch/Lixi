package main

import (
	"fmt"

	"github.com/ZayaSansanch/Lixi/data"
	"github.com/ZayaSansanch/Lixi/functions"
)

func main() {
	fmt.Println(data.Green + "Hello, World!" + data.Reset)

	functions.Read()

	fmt.Println(data.Blue, data.Rs, data.Reset)
}
