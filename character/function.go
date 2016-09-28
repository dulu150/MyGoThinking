package main

import "fmt"

func TestMutiPara(input string, args ...interface{}) (ret int, err string) {

	// 输入参数测试
	fmt.Println(input)

	// 可变参数测试
	var returnvalue int
	for i, v := range args {
		switch v.(type) {
		case int:
			fmt.Println("this is a int:", v)
		default:
			fmt.Println("I don't know what type you are")
		}
		returnvalue = i
	}

	// 多参数返回值测试
	return returnvalue, "I'm nickel"
}
func main() {
	fmt.Println("hello the world")

	// 可变参数、多返回值函数测试
	ret, err := TestMutiPara("I'm test you", 1, "I'm nothing")
	fmt.Println("The return value is:", ret)
	fmt.Println("The return err inf is:", err)
}
