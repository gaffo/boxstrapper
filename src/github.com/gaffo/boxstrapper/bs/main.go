package main

import (
	"os"
	"fmt"
	. "github.com/gaffo/boxstrapper"
)

func usage() {
	fmt.Println("USAGE: bs ap PACKAGES...")
}

func main() {
	 if len(os.Args) < 2 {
	 	usage()
	 	return
	 }

	 driver := UbuntuDriver{}

	 switch os.Args[1] {
	 case "ap":
	 	Ap(driver, os.Args[2:])
	 	return
	 default:
	 	usage()
	 	return
	 }
}