package panic_recover_test

import (
	"fmt"
	"testing"
)

func TestPanicVxExit(t *testing.T) {
	// defer func() {
	// 	if err := recover(); err != nil {
	// 		fmt.Println("recovered from ", err)
	// 	}
	// }()
	fmt.Println("Start!")
	// panic(errors.New("Something wrong!"))
	// os.Exit(-1)
	fmt.Println("End")
}
