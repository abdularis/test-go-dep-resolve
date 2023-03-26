package main

import (
	mod1 "github.com/abdularis/test-go-dep-resolve/mod-1"
	mod2 "github.com/abdularis/test-go-dep-resolve/mod-2"
	modsuper "github.com/abdularis/test-go-dep-resolve/mod-super"
)

func main() {
	mod1.PrintMod1()
	mod2.PrintMod2()
	modsuper.PrintModSuper()
}
