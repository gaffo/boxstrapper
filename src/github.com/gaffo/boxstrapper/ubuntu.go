package boxstrapper

import (
	"fmt"
	"os/exec"
	"io/ioutil"
)

type UbuntuDriver struct {
}

func (UbuntuDriver) AddPackage(packageName string) error {
	fmt.Println("Installing Package:", packageName)
	cmd := exec.Command("sudo", "apt-get", "install", "-y", packageName)
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Println("ERR From PIPE>", err)
		return err
	}
	if err := cmd.Start(); err != nil {
		fmt.Println("ERR FROM START>", err)
		return err
	}
	str, err := ioutil.ReadAll(stdout)
	if err != nil {
		fmt.Println("ERR FROM READALL>", err)
		return err
	}
	fmt.Println(">", string(str))
	if err := cmd.Wait(); err != nil {
		fmt.Println("Error Installing Package", err)
		return err
	}
	return nil
}