package main

import (
	"os"
	"fmt"
	// "boxstrapper"
)

func usage() {
	fmt.Println("bs ap PACKAGES...")
}

func main() int {
	 if len(os.Args) < 2 {
	 	usage()
	 	return 1
	 }

	 driver := UbuntuDriver{}

	 switch os.Args[1] {
	 case "ap":
	 	return boostrapper.Ap(driver)
	 default:
	 	usage()
	 	return 1
	 }
}