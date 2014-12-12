package boxstrapper

import (
	"io/ioutil"
	"os"
	"fmt"
	"github.com/libgit2/git2go"
	"time"
	"log"
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
		fmt.Println("ensureRepo", err)
		return err
	}
	defer repo.Free()

	err = ioutil.WriteFile(
		this.packagesFile(), 
		[]byte(contents), 
		0666)

	if err != nil {
		fmt.Println("WriteFile", err)
		return err
	}

	idx, err := repo.Index()
	if err != nil {
		fmt.Println("Index", err)
		return err
	}

	name, _ := config_str(repo, "user.name")
	email, _ := config_str(repo, "user.email")

	err = idx.AddByPath("packages.bss")
	if err != nil {
		fmt.Println("AddByPath", err)
		return err
	}

	treeId, err := idx.WriteTree()
	if err != nil {
		fmt.Println("WriteTree", err)
		return err
	}

	tree, err := repo.LookupTree(treeId)
	if err != nil {
		fmt.Println("LookupTree", err)
		return err
	}

	sig := &git.Signature{
		Name: name,
		Email: email,
		When: time.Now(),
	}

	commit, err := repo.CreateCommit("HEAD", sig, sig, reason, tree)
	log.Printf("%s now at revision %s\n", this.BaseDir, commit)
	idx.Write()

	return err
}