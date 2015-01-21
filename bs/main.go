package main

import (
	"fmt"
	"github.com/gaffo/boxstrapper"
	"os"
)

func main() {
	driver := boxstrapper.UbuntuDriver{}
	storage := boxstrapper.NewFilesystemStorage("")
	if os.Args[1] != "ap" {
		fmt.Println("Error, only ap is supported")
		return
	}
	boxstrapper.Ap(driver, storage, os.Args[2:])
}
