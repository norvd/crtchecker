package main

import (
	"fmt"

	"github.com/norvd/crtchecker/internal/crtfinder"
)

func main() {
	f := crtfinder.New("https://crt.sh")
	result := f.Find("")
	fmt.Println(string(result))

}
