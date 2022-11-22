package main

import (
	"fmt"
	"path"
	"runtime"
)

func getCurrentAbPathByCaller() string {
	var abPath string
	_, filename, _, ok := runtime.Caller(0)
	if ok {
		abPath = path.Dir(filename)
	}
	return abPath
}

func main() {
	fmt.Println(getCurrentAbPathByCaller())
}
