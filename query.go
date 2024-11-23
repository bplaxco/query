package main

import (
	"fmt"
	"io"
	"log"
	"os"

	"github.com/bplaxco/query/pkg/parser"
)

func main() {
	rawQuery, err := io.ReadAll(os.Stdin)

	if err != nil {
		log.Fatal(err)
	}

	for _, command := range parser.Parse(string(rawQuery)) {
		fmt.Printf("%#v\n", command)
	}
}
