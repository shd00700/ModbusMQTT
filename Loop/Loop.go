package ModbusMQTT

import (
	"bufio"
	"fmt"
	mbc "github.com/shd00700/ModbusMQTT/Modbus"
	"log"
	"os"
	"strconv"
	"strings"
	//"time"
	//"time"
	//"reflect"
	//"strings"
)
func Loop() {
	//var arr []string
	var num int
	var add uint16
	var leng uint16



	/*err := mbc.Open()
	if err != nil {
		log.Println("disconnect",err)
	}*/

	defer mbc.Close()

	data, _ := mbc.ReadReg(1, 0, 10)
	log.Println(data)

	input := bufio.NewScanner(os.Stdin)
Main:
	for {
		Scrclr()
		println("-------------------------------------")
		fmt.Println("Main Menu")
		println("-------------------------------------")
		fmt.Print("1 : Output Coils\n","2 : Input Coils\n","3 : Input Registers\n","4 : Holding Registers\n\n\n\n\n")
		fmt.Print("Select number Enter:")
		fmt.Scanln(&num)

		if num == 1{
		OutputCoils:
			for {
				Scrclr()
				println("-------------------------------------")
				fmt.Println("Output Coils")
				println("-------------------------------------")
				fmt.Print("1 : Write Coils\n","2 : Read Coils\n","3 : Go back\n\n\n")
				fmt.Print("Select number Enter:")
				fmt.Scanln(&num)

				if num == 1{

					Scrclr()
					var arr []string
					var leng int

					println("-------------------------------------")
					fmt.Println("Write Coils")
					println("-------------------------------------")
					fmt.Print("\n\n\nStart address: ")
					fmt.Scanln(&add)
					fmt.Print("length values:")
					fmt.Scanln(&leng)

					fmt.Print("Enter data values: ")
					input.Scan()
					err, _ := strconv.Atoi(input.Text())
					arr = strings.Split(input.Text(), " ")
					if err == 1 || err == 0  && leng == len(arr) && add < 11{
						mbc.WriteCoils(1, add, arr)
						Continue()
						fmt.Scanln(&num)
						if num == 1 {
							continue OutputCoils
						}else if num == 2{
							continue Main
						}
					} else {
						Scrclr()
						Error()
						fmt.Scanln(&num)
						if num == 1{
							continue OutputCoils
						}else if num == 2{
							continue Main
						}
					}
				}
				if num == 2{
					Scrclr()
					println("-------------------------------------")
					fmt.Println("Read Coils")
					println("-------------------------------------")
					fmt.Print("\n\n\nStart address: ")
					fmt.Scanln(&add)
					fmt.Print("length values:")
					fmt.Scanln(&leng)

					data, _ := mbc.ReadCoil(1, add, leng)
					fmt.Println("Data values : ", data)
					Continue()
					fmt.Scanln(&num)
					if num == 1 {
						continue OutputCoils
					} else if num == 2 {
						continue Main
					}
				}
				if num == 3{
					continue Main
				} else{
					fmt.Println("Please enter again")
					continue OutputCoils
				}
			}
		}
		if num == 2{
		InputCoils:
			for {
				Scrclr()
				println("-------------------------------------")
				fmt.Println("Input Coils")
				println("-------------------------------------")
				fmt.Print("\n\n\nStart address: ")
				fmt.Scanln(&add)
				fmt.Print("length values:")
				fmt.Scanln(&leng)

				data, _ := mbc.ReadCoilIn(1, add, leng)
				fmt.Println("Data value:",data)
				Continue()
				fmt.Scanln(&num)
				if num == 1 {
					continue InputCoils
				} else if num == 2 {
					continue Main
				}
			}
		}
		if num == 3{
		InputRegisters:
			for {
				Scrclr()
				println("-------------------------------------")
				fmt.Println("Input Registers")
				println("-------------------------------------")
				fmt.Print("\n\n\nStart address: ")
				fmt.Scanln(&add)
				fmt.Print("length values:")
				fmt.Scanln(&leng)
				data, _ := mbc.ReadRegIn(1, add, leng)
				fmt.Println("Data values : ",data)
				Continue()
				fmt.Scanln(&num)
				if num == 1 {
					continue InputRegisters
				} else if num == 2 {
					continue Main
				}
			}
		}

		if num == 4 {
		HoldingRegisters:
			for {

				Scrclr()
				println("-------------------------------------")
				fmt.Println("Holding Registers")
				println("-------------------------------------")
				fmt.Print("1 : Write Registers\n","2 : Read Registers\n","3 : Go back\n\n\n")
				fmt.Print("Select number Enter:")
				fmt.Scanln(&num)
				if num == 1 {
					Scrclr()
					var arr []string
					var leng int
					println("-------------------------------------")
					fmt.Println("Write Registers")
					println("-------------------------------------")
					fmt.Print("\n\n\nStart address: ")
					fmt.Scanln(&add)
					fmt.Print("length values:")
					fmt.Scanln(&leng)
					fmt.Print("Enter Data values: ")
					input.Scan()
					err, _ := strconv.Atoi(input.Text())

					if err > 65535 {
						Scrclr()
						fmt.Println("\n[Max Excess error]")
						Error()
						fmt.Scanln(&num)
						if num == 1{
							continue HoldingRegisters
						}else if num == 2{
							continue Main
						}
					}
					arr = strings.Split(input.Text()," ")
					//bbb := strconv.Itoa(leng)
					if leng == len(arr){
						mbc.WriteRegs(1, add, arr)
						Continue()
						fmt.Scanln(&num)
						if num == 1 {
							continue HoldingRegisters
						} else if num == 2 {
							continue Main
						}
					}else{
						Scrclr()
						fmt.Println("\n[Entered incorrectly length values]")
						Error()
						fmt.Scanln(&num)
						if num == 1{
							continue HoldingRegisters
						}else if num == 2{
							continue Main
						}
					}

				}
				if num == 2{
					Scrclr()
					println("-------------------------------------")
					fmt.Println("Read Registers")
					println("-------------------------------------")
					fmt.Print("\n\n\nStart address: ")
					fmt.Scanln(&add)
					fmt.Print("length values:")
					fmt.Scanln(&leng)
					data, _ := mbc.ReadReg(0, add, leng)
					fmt.Println("Data values : ",data)
					Continue()
					fmt.Scanln(&num)
					if num == 1 {
						continue HoldingRegisters
					} else if num == 2 {
						continue Main
					}
				}
				if num == 3{
					continue Main
				} else{
					continue HoldingRegisters
				}
			}
		}
	}
}
