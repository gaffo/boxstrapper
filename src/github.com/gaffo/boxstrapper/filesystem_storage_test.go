package boxstrapper_test

import (
	"testing"
	. "github.com/gaffo/boxstrapper"
  	"github.com/stretchr/testify/assert"
  	"io/ioutil"
  	"os"
	"github.com/libgit2/git2go"
	"log"
)

func cT() {
	_ = os.RemoveAll("tmp")
}

func fileExists(file string) bool {
	if _, err := os.Stat(file); err != nil {
		return false
	}
	return true
}

func fileContents(file string) string {
	bytes, err := ioutil.ReadFile(file)
	if err != nil {
		return ""
	}
	return string(bytes)
}

func walkRepo(base string, fun git.RevWalkIterator) {
	log.Print("Walking")
	repo, err := git.OpenRepository(base)
	if err != nil {
		log.Print("Open", err)
		return
	}
	walk, err := repo.Walk()
	if err != nil {
		log.Print("Walk ", err)
		return
	}

	err = walk.PushRange("HEAD..HEAD")
	if err != nil {
		log.Print("PushRange ", err)
		return
	}

	walk.Iterate(fun)
}

func listCommitMessages(base string) []string {
	commits := make([]string, 0, 32)
	walkRepo(base, func (commit *git.Commit) bool {
		log.Print("Iterate Message")
		commits = append(commits, commit.Message())
		return true
	})
	return commits
}

func listCommitFiles(base string) [][]string {
	files := make([][]string, 0, 32)
	walkRepo(base, func (commit *git.Commit) bool {
		log.Print("Iterate File")
		tree, err := commit.Tree()
		if err != nil {
			log.Print("commit.Tree ", err)
			return true
		}
		fileList := make([]string, 0, 32)
		tree.Walk(func(str string, te *git.TreeEntry) int {
			fileList = append(fileList, te.Name)
			return 0
		})

		files = append(files, fileList)
		return true
	})
	return files
}

func TestReadPackages_EmptyRepo(t *testing.T) {
	assert := assert.New(t)
	storage := NewFilesystemStorage("/nonexistent")

	data, err := storage.ReadPackages()
	assert.NotNil(err)
	assert.Equal("", data)
}

func TestReadPackages_RepoWithPackagefile(t *testing.T) {
	assert := assert.New(t)
	defer cT()
	_ = os.MkdirAll("tmp", os.ModePerm)
	_ = ioutil.WriteFile("tmp/packages.bss", []byte(`contents`), os.ModePerm)

	storage := NewFilesystemStorage("tmp")

	data, err := storage.ReadPackages()
	assert.Nil(err)
	assert.Equal(`contents`, data)
}

func TestWritePackages_NoRepo(t *testing.T) {
	assert := assert.New(t)
	defer cT()

	storage := NewFilesystemStorage("tmp")
	err := storage.WritePackages("packages", "reason")
	assert.Nil(err)

	assert.True(fileExists("tmp"))
	assert.True(fileExists("tmp/packages.bss"))
	assert.Equal("packages", fileContents("tmp/packages.bss"))
	assert.Equal([]string{"reason"}, listCommitMessages("tmp"))
	assert.Equal([][]string{{"packages.bss"}}, listCommitFiles("tmp"))
}