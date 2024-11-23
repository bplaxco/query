package main

import (
	"fmt"
	"io"
	"log"
	"os"

	"github.com/bplaxco/query/pkg/exec"
	"github.com/bplaxco/query/pkg/parser"
)

func main() {
	rawQuery, err := io.ReadAll(os.Stdin)

	if err != nil {
		log.Fatal(err)
	}

	ctx, err := exec.ExecList(parser.Parse(string(rawQuery)))
	fmt.Printf("#%#v %#v\n", ctx, err)
}
