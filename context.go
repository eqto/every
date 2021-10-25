package every

//Context ..
type Context struct {
	values map[string]interface{}
	hour   uint8
	minute uint8
}

func (c *Context) Hour() int {
	return int(c.hour)
}

func (c *Context) Minute() int {
	return int(c.minute)
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
