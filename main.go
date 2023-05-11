package main

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"
	"log"
	
)

func main() {
	files, err := ioutil.ReadDir("config")
	if (err != nil) {
		log.Fatal(err)
	}

	mgr := NewJsonConfigManager()
	for _, f := range files {
		ext := filepath.Ext(f.Name())
		if strings.Compare(ext, ".json") != 0 {
			continue
		}

		fname := filepath.Join("config", f.Name())
		fmt.Println("parse ", fname)

		content, err := ioutil.ReadFile(fname)
		if (err != nil) {
			log.Fatal(err)
		}

		conf := NewJsonConfig(f.Name())
		conf.Parse(content)
		k := strings.TrimSuffix(f.Name(), filepath.Ext(f.Name()))
		fmt.Println(k)
		mgr.AddConf(k, conf)
	}

}