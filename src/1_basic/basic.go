package main

import "fmt"

/**
 GO 定义变量
 */
func variableZeroValue() {
	/**
	变量名在前，类型在后
	 */
	var a int
	var s string

	/**
	print the initial value in go,
	the init in cpp is random, in java is "bull".
	but in go all sys data type has it's init value.
	 */

	 /**
	 if you use the "println", you can just print the value by the name of variable.
	 in this way you can just print the a, and it won't print the s.
	  */

	fmt.Println(a,s)
	fmt.Println("====================")
	/**
	if you want to print the inital value of a string, you need ti use the "printf".

	"%q" is mean "quotation marks",if we use "%s" to print the initial value of a string ,it will print nothing.
	 */

	fmt.Printf("%d   %q\n", a, s)
}

func variableInitialVal() {
	var a, b int = 3,4
	var s string = "abc"

	fmt.Println(a, s, b)
}

/**
GO can auto judge the type of variable ,which according to the value type that you assign to.
 */
func variableTypeDeduction() {
	// Assignment statement can assign different type of variable.
	var a, b, c, s = 3, 4, true, "def"
	fmt.Println(a, b, c, s)

}

func variableShorter() {
	//we can use a shorter Assignment by use the ":=".
	// but the ":=" can only just and must use in initial value assign.
	a, b, c, s := 3, 4, true, "def"
	// if you re-assign the variable you can just use "="
	b = 5
	fmt.Println(a, b, c, s)
}

//we can define the variable out of method, but we must use the key word "var".
var aa = 3
//if we use the ":=", the IDE will report a mistake.
//c := 88

// all the out-method variable is not mean these variable are Global variable, it's mean the variable is
// a "package variable", these variable is "Global" in this package.

//ths out-method variable can define in this way

var (
	dd = 8
	vv = "jjj"
	bb = true
)

func main() {
	//fmt.Println("广哥牛逼啊！  ")

	//variableZeroValue()

	//variableInitialVal()

	//variableTypeDeduction()

	//variableShorter()

	fmt.Println(aa, dd, vv, bb)
}
