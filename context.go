package every

//Context ..
type Context struct {
	Debug func(d ...interface{})
	Info  func(i ...interface{})
	Error func(e ...interface{})
}
