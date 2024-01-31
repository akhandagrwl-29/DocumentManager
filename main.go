package main

import (
	"DocumentManager/errors"
	"DocumentManager/processor"
	"os"
)

func main() {
	if len(os.Args) > 1 {
		processor.ProcessFile(os.Args[1])
	} else {
		panic(errors.ErrNoInput)
	}
}
