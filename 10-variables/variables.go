package main

import "fmt"

func main() {
	var num1 uint32 = 55;
	fmt.Printf("num1 :: Datatype - %T and Value - %d", num1, num1);

	fmt.Println();

	var num2 = 8998.034;
	fmt.Printf("num2 :: Datatype - %T and Value with precision 1 - %5.1f", num2, num2);

	fmt.Println();

	num3 := "Our Green Earth";
	fmt.Println("num3 :: ", num3, " and it is declared using Walrus Operator.");

	var a, b int = 3, 4
	fmt.Println("value of a = ", a, " and for b = ", b)

	c, d := 44, 45
	fmt.Println("value of c = ", c, " and for d = ", d)

	var e, f = 3442, 1223
	fmt.Println("value of e = ", e, " and for f = ", f)
}
