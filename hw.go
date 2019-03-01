package main  // 1. Must declare package
import "fmt"  //2. import package, also  called  preprocessor command which tell go to compile files lying in package fmt

func main () /*3.main function where execution starts */{
//4. This my friend is a comment !!
fmt.Println("Hello, World!")  // another function in go where fmt package call Println method
}
