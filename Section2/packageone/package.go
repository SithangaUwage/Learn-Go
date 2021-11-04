package packageone

import "fmt"

var PackageVar = "Package Var"

func PrintMe(myVar, blockVar string) {
	fmt.Println(myVar, blockVar, PackageVar)
}
