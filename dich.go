package dich

import (
	"crypto/md5"
	"fmt"
	"io"
	"os"
)

type Checker struct {
	cashNameAndHash *configJson
}

func NewChecker() *Checker {
	err := os.WriteFile("./dich.json", []byte("{}"), 0644)
	if err != nil {
		panic(err)
	}
	c := NewConfig()
	return &Checker{
		cashNameAndHash: c,
	}
}

func (c *Checker) AddFileList(pathTofile string) {
	sum := c.CreateSumFromFile(pathTofile)
	c.cashNameAndHash.Add(pathTofile, sum)
}
func (c *Checker) RemoveFileList() {}
func (c *Checker) CheckChange(fileName string, hashSum string) bool {
	sum := c.cashNameAndHash.Get(fileName)
	return sum == hashSum
}

func (c *Checker) CreateSumFromFile(path string) string {
	h := md5.New()
	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	_, err = io.Copy(h, f)
	if err != nil {
		panic(err)
	}
	return fmt.Sprintf("%x", h.Sum(nil))
}

func (c *Checker) CreateSum(data string) string {
	h := md5.Sum([]byte(data))
	return fmt.Sprintf("%x", h)
}

func (c *Checker) GetAll() {
	fmt.Println(c.cashNameAndHash)
}
