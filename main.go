package main

import (
	"fmt"

	"version/buildinfo"
)

func main() {
	fmt.Println(buildinfo.Version())
}
