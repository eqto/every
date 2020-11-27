package every

//Context ..
type Context struct {
	values map[string]interface{}
}

//PutValue ..
func (c *Context) PutValue(key string, value interface{}) {
	c.values[key] = value
}

//GetValue ..
func (c *Context) GetValue(key string) interface{} {
	return c.values[key]
}

//NewContext ..
func NewContext() *Context {
	return &Context{values: make(map[string]interface{})}
}
