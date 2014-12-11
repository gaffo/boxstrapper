package boxstrapper

import (
	"io/ioutil"
	"os"
	"fmt"
	"os/exec"
	"strings"
)

type FilesystemStorage struct {
}

func boxstrap_dir() string {
	// NOTE this only works on unix
	return fmt.Sprintf("%s/boxstrap.d", os.Getenv("HOME"))
}

func boxstrap_join(path string) string {
	return fmt.Sprintf("%s/%s", boxstrap_dir(), path)
}

func packages_file() string {
	return boxstrap_join("packages.bss")
}

func ensure_repo() error {
	if _, err := os.Stat(boxstrap_dir()); err != nil {
		err = os.Mkdir(boxstrap_dir(), os.ModePerm)
		if err != nil {
			return err
		}
	}
	if _, err := os.Stat(boxstrap_join(".git")); err != nil {
		cmd := exec.Command("git", "init")
		cmd.Dir = boxstrap_dir()
		err = cmd.Run()
		if err != nil {
			return err
		}
	}
	return nil
}

func (FilesystemStorage) ReadPackages() (string, error) {
	bytes, err := ioutil.ReadFile(packages_file())

	if err != nil {
		return "", err
	}

	return string(bytes), err
}

func (FilesystemStorage) WritePackages(contents string, reason string) error {
	err := ensure_repo()
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(
		packages_file(), 
		[]byte(contents), 
		os.ModePerm)

	if err != nil {
		return err
	}

	cmd := exec.Command("git", "add", "packages.bss")
	cmd.Dir = boxstrap_dir()
	err = cmd.Run()
	if err != nil {
		fmt.Println("Add Err>", err)
		return err
	}
	reason = strings.Replace(reason, `"`, `'`, 0)
	reason = fmt.Sprintf(`"%s"`, reason)
	cmd = exec.Command("git", "commit", "-m", reason)
	cmd.Dir = boxstrap_dir()
	err = cmd.Run()
	if err != nil {
		fmt.Println("Commit Err>", err)
		return err
	}
	return nil
}