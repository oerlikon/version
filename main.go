package main

import (
	"fmt"

	"version/buildinfo"
)

func main() {
	fmt.Println("version", buildinfo.Version())
}
