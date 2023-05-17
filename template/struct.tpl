package config

import (
    "encoding/json"
)

type loader func(map[string]interface{}) error

var (
    loaders map[string]loader
)   

func LoadConf(name string, m map[string]interface{}) {
    f := loaders[name]
    f(m)
}



<%= for (n) in names { %>
func load<%=n%>(m map[string]interface{}) error {
    obj := new<%=n%>()
    return obj.loadFromMap(m)
}

func new<%=n%>() *<%=n%> {
    return &<%=n%> {}
}

func (c *<%=n%>) loadFromMap(m map[string]interface{})  error{
    data, err := json.Marshal(m)
	if err == nil{
		err = json.Unmarshal(data, c)
	}
	return err
}
<%} %>

func InitLoaders() {
loaders = map[string]loader {
    <%= for (n) in names { %>
    "<%=n%>": load<%=n%>,
    <%} %>
}
}