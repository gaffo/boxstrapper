package main

import (
	"os"
	"fmt"
	. "github.com/gaffo/boxstrapper"
)

func usage() {
	fmt.Println(`USAGE: bs COMMAND
Where commands are:
	ap package[:group,...]...
	recover

DETAILS:

ap: 
	Adds a list of packages with optional groups (which are just tags for later
    knowing why you added the package). Ap will install each of hte packages using
    your package manager and then commit to ~/boxstrap.d an updated packages.bss
    file. This file can be later used on a new machine to recover all your installs.

recover:
	Rebuilds the machine to the best of it's ability using info in ~/boxstrap.d.
	Currently you will need to pull the repo yourself.

MORE INFO:
	Can be found at http://github.com/gaffo/boxstrapper`)
}

func main() {
	 if len(os.Args) < 2 {
	 	usage()
	 	return
	 }

	 driver := UbuntuDriver{}
	 storage := NewFilesystemStorage("")

	 switch os.Args[1] {
	 case "ap":
	 	Ap(driver, storage, os.Args[2:])
	 	return
	 case "recover":
	 	Recover(driver, storage)
	 	return
	 default:
	 	usage()
	 	return
	 }
}