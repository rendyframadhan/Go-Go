package main

import "fmt"

func main() {
	fmt.Println("Hello Golang")

	// String
	fmt.Println("======== Learning String ========")
	fmt.Println("String : ", "asd")
	fmt.Println("Total Char String : ", len("asd"))
	fmt.Println("Get Char String (Byte): ", "asdf"[1])

	// Integer
	fmt.Println("======== Learning Integer ========")
	var nilai32 int32 = 1234512326
	var nilai64 int64 = int64(nilai32)
	var nilai16 int16 = int16(nilai32)
	fmt.Println("Test Integer 32 : ", nilai32)
	fmt.Println("Test Integer 64 : ", nilai64)
	fmt.Println("Test Integer 16 : ", nilai16)
	fmt.Println("One : ", 1)
	fmt.Println("One Point Two (Decimal) : ", 1.2)

	// Boolean
	fmt.Println("======== Learning Boolean ========")
	fmt.Println("Boolean : ", true)

	// Declare Constant
	fmt.Println("======== Declare Constant ========")
	const INDEX = 1
	const PAYMENT_TYPE = "Debit"
	fmt.Println("Constant Integer : ", INDEX)
	fmt.Println("Constant String : ", PAYMENT_TYPE)

	// How to declare in Simplified version using :=
	fmt.Println("======== Declare Simplified Version ========")
	simplifiedVarString := "asd"
	simplifiedVarInteger := 1
	simplifiedVarDecimal := 2.5

	fmt.Println(simplifiedVarString)
	fmt.Println(simplifiedVarInteger)
	fmt.Println(simplifiedVarDecimal)

	// Override defined Variable
	fmt.Println("======== Override And Define Default Variable ========")
	var defaultString string // <= will be empty string cause the value not defined
	var defaultNumber int    // <= will be Zero cause the value not defined
	fmt.Println(defaultString)
	fmt.Println(defaultNumber)
	defaultString = "Override The Name"
	defaultNumber = 123
	fmt.Println(defaultString)
	fmt.Println(defaultNumber)

	// Combine variable
	fmt.Println("======== Declare variable into one ========")
	var (
		hello    = "Hello "
		test     = "123"
		decimal  = 3.5
		number   = 1
		isItTrue = true
	)

	fmt.Println(hello + test)
	fmt.Println(hello+test, decimal, number)
	fmt.Println(decimal)
	fmt.Println(number)
	fmt.Println(isItTrue)

	// Arithmetic Operational
	fmt.Println("======== Arithmetic Operational ========")
	var a = 1
	var b = 2
	fmt.Println("Addition : ", a+b)
	fmt.Println("Addition Decimal : ", decimal+float64(number))
	fmt.Println("Subtraction Decimal : ", float64(a)-decimal)
	fmt.Println("Multiplication : ", a*b)
	fmt.Println("Division : ", a/b)
}
