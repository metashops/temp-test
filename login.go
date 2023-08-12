func MyMain()  {
	fmt.Println("我是Main")
	fmt.Println(MyD())
	fmt.Println(MyF())
	fmt.Println(MyFC())
	fmt.Println(MyGC())
}

// D 开发
func MyD() {
	fmt.Println("是我D函数")
}

// E 开发
func MyD() {
	fmt.Println("是我E函数")
}

// F 开发
func MyF() {
fmt.Println("是我F函数")
}

// FC 开发
func MyF() {
	fmt.Println("是我FC函数")
}


// GC 开发
func MyGC() {
fmt.Println("是我GC函数")
}
