package boxstrapper

import (
	"io/ioutil"
	"os"
	"fmt"
	"github.com/libgit2/git2go"
)

type FilesystemStorage struct {
	BaseDir string
}

func NewFilesystemStorage(basedir string) (*FilesystemStorage) {
	storage := new(FilesystemStorage)
	if basedir == "" {
		storage.BaseDir = boxstrap_dir()
	} else {
		storage.BaseDir = basedir
	}
	return storage
}

func boxstrap_dir() string {
	// NOTE this only works on unix
	return fmt.Sprintf("%s/boxstrap.d", os.Getenv("HOME"))
}

func (this *FilesystemStorage) path(path string) string {
	return fmt.Sprintf("%s/%s", this.BaseDir, path)
}

func (this *FilesystemStorage) packagesFile() string {
	return this.path("packages.bss")
}

func (this *FilesystemStorage) ensureRepo() (* git.Repository, error) {
	if _, err := os.Stat(this.BaseDir); err != nil {
		err = os.Mkdir(this.BaseDir, os.ModePerm)
		if err != nil {
			return nil, err
		}
	}
	if _, err := os.Stat(this.path(".git")); err != nil {
		repo, err := git.InitRepository(this.BaseDir, false)
		if err != nil {
			return nil,err
		}
		return repo, nil
	}
	return git.OpenRepository(this.BaseDir)
}

func (this *FilesystemStorage) ReadPackages() (string, error) {
	bytes, err := ioutil.ReadFile(this.packagesFile())

	if err != nil {
		return "", err
	}

	return string(bytes), err
}

func config_str(repo *git.Repository, key string) (string, error) {
	config, err := repo.Config()
	if err != nil {
		return "", err
	}
	return config.LookupString(key)
}

func (this *FilesystemStorage) WritePackages(contents string, reason string) error {
	repo, err := this.ensureRepo()
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(
		this.packagesFile(), 
		[]byte(contents), 
		os.ModePerm)

	if err != nil {
		return err
	}

	idx, err := repo.Index()
	if err != nil {
		return err
	}
	idx.AddByPath("packages.bss")
	name, _ := config_str(repo, "user.name")
	email, _ := config_str(repo, "user.email")
	fmt.Println("user:", name)
	fmt.Println("email:", email)

	return nil
}