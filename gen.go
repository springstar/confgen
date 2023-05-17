package main

import (
	"github.com/springstar/confgen/config"
	_ "reflect"
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"io/ioutil"
	"path/filepath"
	"strings"
	"os"
	"log"
	"github.com/m-zajac/json2go"
	"github.com/gobuffalo/plush"

)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

type JsonConfig struct {
	name string
	// struct slice, every map is a struct, like ConfAI, ConItem etc
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
	return &JsonConfigManager{
		confs: make(map[string]*JsonConfig),
	}
}

func (m *JsonConfigManager) loadConf(path string) {
	files, err := ioutil.ReadDir(path)
	if (err != nil) {
		log.Fatal(err)
	}

	for _, f := range files {
		ext := filepath.Ext(f.Name())
		if strings.Compare(ext, ".json") != 0 {
			continue
		}

		fname := filepath.Join(path, f.Name())
		content, err := ioutil.ReadFile(fname)
		if (err != nil) {
			log.Fatal(err)
		}

		conf := NewJsonConfig(f.Name())
		conf.Parse(content)
		k := strings.TrimSuffix(f.Name(), filepath.Ext(f.Name()))
		m.addConf(k, conf)
	}
}

func (m *JsonConfigManager) addConf(name string, conf *JsonConfig) {
	if _, ok := m.confs[name]; ok {
		return
	}

	m.confs[name] = conf
}

func (m *JsonConfigManager) findConf(name string, sn int) interface{} {
	if conf, ok := m.confs[name]; ok {
		return conf
	}

	return nil
}

type ConfigGen struct {

}

func genStructs() (names []string, defines []string) {
	files, err := ioutil.ReadDir("config")
	if (err != nil) {
		log.Fatal(err)
	}

	for _, f := range files {
		ext := filepath.Ext(f.Name())
		if strings.Compare(ext, ".json") != 0 {
			continue
		}

		sname, _ := strings.CutSuffix(f.Name(), ".json")
		names = append(names, sname)

		file := filepath.Join("config", f.Name())
		content, _ := ioutil.ReadFile(file)
		parser := json2go.NewJSONParser(sname)
		parser.FeedBytes(content)
		res := parser.String()
		out := strings.Replace(res, "[]", " ", 1)
		defines = append(defines, out)
	}

	return names, defines
	
}

func writeMethods(names []string) {
	content, err := ioutil.ReadFile("template/struct.tpl")
	if (err != nil) {
		log.Fatal(err)
	}

	template := string(content[:])
	ctx := plush.NewContext()
	ctx.Set("names", names)

	s, err := plush.Render(template, ctx)
	if (err != nil) {
		log.Fatal(err)
	}

	file, err := os.Create("config/loader.go")
    if err != nil {
        return
    }
    defer file.Close()

    file.WriteString(s)
}

func writeStructs(str []string) {
	s := strings.Join(str, "\n\n")
	file, err := os.Create("config/structs.go")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer file.Close()
	
	file.WriteString("package config")
	file.WriteString("\n\n\n")
	file.WriteString(s)
}



