package main

import (
	"encoding/json"
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