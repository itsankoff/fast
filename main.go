package main

import (
	"fmt"
	"os"

	"github.com/itsankoff/fast/api/fast"
)

func bandwidth() {
	sourceCount := os.Getenv("GOMAXPROCS")

	fastAPI := fast.New(true)
	sources, err := fastAPI.GetDownloadURLs(sourceCount)
}

func main() {
	fmt.Println("fast")
}
