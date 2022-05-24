package main

import (
	"fmt"
	"bufio"
	"os"
)

func main() {

	var aString string = "This is Go!"
	fmt.Println(aString)
	fmt.Printf("The variable's type is %T\n", aString)

	reader := bufio.NewReader(os.Stdin);

	fmt.Println("Enter text: ");
	input,_ := reader.ReadString('\n')

	fmt.Println(input);

}
