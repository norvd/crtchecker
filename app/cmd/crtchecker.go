package main

import (
	"fmt"
	"log"

	"github.com/norvd/crtchecker/internal/crtfinder"
)

func main() {
	f := crtfinder.New("https://crt.sh")
	certs, err := f.Find("Wildberries")
	if err != nil {
		log.Fatalf("failed to find certs: %v", err)
	}
	fmt.Println(len(certs))
	for i, cert := range certs {
		fmt.Println(i, cert.CommonName)
	}

}
