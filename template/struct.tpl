package config

<%= for (n) in names { %>
func New<%=n%>() *<%=n%> {
    return &<%=n%> {}
}

func (c *<%=n%>) LoadFromMap(m map[string]interface{})  error{
    data, err := json.Marshal(m)
	if err == nil{
		err = json.Unmarshal(data, c)
	}
	return err
}
<%} %>