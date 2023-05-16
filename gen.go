package main

import (
	"fmt"
	"encoding/json"
	"io/ioutil"
	"path/filepath"
	"strings"
	"os"
	"log"
	"github.com/m-zajac/json2go"

)

type JsonConfig struct {
	name string
	data []map[string]interface{}
}

func NewJsonConfig(name string) *JsonConfig {
	return &JsonConfig{
		name : name,
	}
}

func (c *JsonConfig) Parse(bytes []byte) error {
	if err := json.Unmarshal(bytes, &c.data); err != nil {
		panic(err)
	}

	return nil
}

type JsonConfigManager struct {
	confs map[string]*JsonConfig
}

func NewJsonConfigManager() *JsonConfigManager {
	return &JsonConfigManager{}
}

func (m *JsonConfigManager) AddConf(name string, conf *JsonConfig) {
	if _, ok := m.confs[name]; !ok {
		return
	}

	m.confs[name] = conf
}

type ConfigGen struct {

}

func genStruct() string {
	files, err := ioutil.ReadDir("config")
	if (err != nil) {
		log.Fatal(err)
	}

	var s []string

	for _, f := range files {
		ext := filepath.Ext(f.Name())
		if strings.Compare(ext, ".json") != 0 {
			continue
		}

		sname, _ := strings.CutSuffix(f.Name(), ".json")

		file := filepath.Join("config", f.Name())
		content, _ := ioutil.ReadFile(file)
		parser := json2go.NewJSONParser(sname)
		parser.FeedBytes(content)
		res := parser.String()
		out := strings.Replace(res, "[]", " ", 1)
		s = append(s, out)
	}

	str := strings.Join(s, "\n\n")
	return str
	
}

func write(s string) {
	file, err := os.Create("config/gen/structs.go")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer file.Close()
	
	file.WriteString("package config")
	file.WriteString("\n\n\n")
	file.WriteString(s)
}



