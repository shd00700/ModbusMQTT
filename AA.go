package LOOP

import(
	"fmt"

)

func Loop() {
	var num int



Main:
	for {
		println("-------------------------------------")
		fmt.Println("Main Menu")
		println("-------------------------------------")
		fmt.Print("1 : Output Coils\n", "2 : Input Coils\n", "3 : Input Registers\n", "4 : Holding Registers\n\n\n\n\n")
		fmt.Print("Select number Enter:")
		fmt.Scanln(&num)
		if num == 1 {
			fmt.Println("1번")
			continue Main
		}
		if num == 2 {
			fmt.Println("2번")
			continue Main
		} else {
			continue Main
		}
	}
}
