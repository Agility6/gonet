package main

import (
	"fmt"
	"gonet/http"
	"path/filepath"
)

func main() {
	path, err := filepath.Abs(".")
	if err != nil {
		fmt.Println("filepath abs:", err)
		return
	}
	http.SetDir(path)
	http.Run()
}
