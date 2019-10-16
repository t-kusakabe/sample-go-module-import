package main

import (
	"fmt"

	"github.com/t-kusakabe/sample-go-module-import/subdir"
)

func main() {
	fmt.Println(subdir.SubDirFunc())
}
