package mod2

import (
	"fmt"

	modsuper "github.com/abdularis/test-go-dep-resolve/mod-super/v2"
)

func PrintMod2() {
	fmt.Println("------ mod 2 start -------")
	modsuper.PrintModSuper()
	fmt.Println("------ mod 2 end -------")
}
