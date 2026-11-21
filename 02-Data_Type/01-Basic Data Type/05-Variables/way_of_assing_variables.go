package main
import (
	"fmt"
)

func main(){
	// simple variable declaration
	var name string
	name="sameer"
	fmt.Println(name)

	// another way
	var age=34
	fmt.Println(age)

	// shorthand variable
	isTestDone:=false

	// idiomic way
	var(
		cpu string="i5"
		cores int=4
		genration int=12
	)
	if !isTestDone{
		fmt.Printf("Pc specs are\nCpu: %s\nCores: %d\nGeneration: %d\n",cpu,cores,genration)
		isTestDone=true
		fmt.Println("Tese done")


	}
	
}
