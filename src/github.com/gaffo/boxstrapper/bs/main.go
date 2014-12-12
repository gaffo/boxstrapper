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
	bs watch FILEPATH...

DETAILS:

ap: 
	Adds a list of packages with optional groups (which are just tags for later
    knowing why you added the package). Ap will install each of hte packages using
    your package manager and then commit to ~/boxstrap.d an updated packages.bss
    file. This file can be later used on a new machine to recover all your installs.

    examples:
    	bs ap vim
    	bs ap emacs:dev blender:games gimp:games,default

    creates:
    	packages.bss: PACKAGENAME: groups

recover:
	Rebuilds the machine to the best of it's ability using info in ~/boxstrap.d.
	Currently you will need to pull the repo yourself.

	example:
		bs recover

	creates:
		NONE

watch:
	Tells boxstrap to watch a file. Any time boxstrap is run, the file will be checked against
	the version cached in boxstrap.d/files to the file on disk. If the file has changed, boxstrap will
	check in the new revision of the file

	WARNING: Currently only use this for small files as they will be versioned in boxstrap.d

	example:
		bs watch ~/.zshrc

	creates:
		watch(FILEPATH): groups

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
	 case "watch":
	 	Watch(driver, storage, os.Args[2:])
	 	return
	 default:
	 	usage()
	 	return
	 }
}