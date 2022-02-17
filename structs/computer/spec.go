package computer

//大写导出结构体
type Spec struct { //exported struct
	Maker string //exported field
	model string //unexported field
	Price int    //exported field
}
