package functions

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/ZayaSansanch/Lixi/data"
)

func Read() {
	file, err := os.Open("code/code.txt")

	if err != nil {
		log.Fatalf("Error when opening file: %s", err)
	}

	fileScanner := bufio.NewScanner(file)

	for fileScanner.Scan() {
		fmt.Println(fileScanner.Text())
		if fileScanner.Text() == "//" {
			data.Rs++
		}
	}

	if err := fileScanner.Err(); err != nil {
		log.Fatalf("Error while reading file: %s", err)
	}

	file.Close()
}
