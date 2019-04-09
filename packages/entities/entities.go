package entities

//定义一个User用户类型
type user struct {
	Name  string
	Email string
}

//定义一个admin 类型
type Admin struct {
	user //嵌入类型示未公开的 首字母小写是为公开
	Rights int
}
