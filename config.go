package dich

import (
	"encoding/json"
	"io/ioutil"
	"os"

	jsoniter "github.com/json-iterator/go"
)

type configJson struct {
	file jsoniter.Any
}

func NewConfig() *configJson {
	file, err := os.Open("./dich.json")
	if err != nil {
		panic(err)
	}
	byteValue, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	jsons := jsoniter.Get(byteValue)
	return &configJson{
		file: jsons,
	}
}

func (c *configJson) Add(key string, value interface{}) error {
	var m map[string]interface{}
	err := json.Unmarshal([]byte(c.file.Get().ToString()), &m)
	if err != nil {
		return err
	}
	m[key] = value
	b, err := jsoniter.Marshal(m)
	if err != nil {
		return err
	}
	c.file = jsoniter.Get(b)
	err = os.WriteFile("./dich.json", []byte(c.file.ToString()), 0644)
	if err != nil {
		return err
	}
	return nil
}
func (c *configJson) Remove(key string) error {
	var m map[string]interface{}
	err := json.Unmarshal([]byte(c.file.Get().ToString()), &m)
	if err != nil {
		return err
	}
	delete(m, key)
	b, err := jsoniter.Marshal(m)
	if err != nil {
		return err
	}
	c.file = jsoniter.Get(b)
	err = os.WriteFile("./dich.json", []byte(c.file.ToString()), 0644)
	if err != nil {
		return err
	}
	return nil
}
func (c *configJson) Get(key string) string {
	return c.file.Get(key).ToString()
}
