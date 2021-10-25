package every

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

func (c *Context) PutValue(key string, value interface{}) {
	c.values[key] = value
}

func (c *Context) GetValue(key string) interface{} {
	return c.values[key]
}

func NewContext() *Context {
	return &Context{values: make(map[string]interface{})}
}
