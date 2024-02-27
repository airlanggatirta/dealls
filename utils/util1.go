package utils

import (
	"fmt"
)

// varDump will print out any number of variables given to it
// e.g. varDump("test", 1234)
func VarDump(myVar ...interface{}) {
	fmt.Printf("%v\n", myVar)
}

// dd will print out variables given to it (like varDump()) but
// will also stop execution from continuing.
func Dd(myVar ...interface{}) {
	VarDump(myVar...)
	//os.Exit(1)
}
