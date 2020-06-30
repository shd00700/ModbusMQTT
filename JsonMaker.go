package ModbusMQTT

import (
	"encoding/json"
	//"os"
)

type ReadCoilStruct struct {
	Code_name string
	Start_Address uint16
	Last_Address uint16
}

type WriteCoilStruct struct {
	Code_name string
	Start_Address uint16
	Last_Address uint16
}

type ReadRegStruct struct {
	Code_name string
	Start_Address uint16
	Last_Address uint16
}

type WriteRegStruct struct {
	Code_name string
	Start_Address uint16
	Last_Address uint16
}

type ErrJson struct {
	Errormessage string
}

type ReadRegInStruct struct {
	Alias string
	Address uint16
	Value uint16
}


var ReadReg_alias = []string{"Un_name0","Un_name1","Un_name2","Un_name3","Un_name4","Un_name5","Un_name6",
	"Un_name7","Un_name8","Un_name9","Un_name10","Un_name11","Cn_name12","Un_name13","Un_name14",
	"Un_name15","Un_name16","Un_name17","Un_name18","Un_name19","Un_name20"}

var ReadCoil_alias = []string{"Un_name0","Un_name1","Un_name2","Un_name3","Un_name4","Un_name5","Un_name6",
	"Un_name7","Un_name8","Un_name9","Un_name10","Un_name11","Un_name12","Un_name13","Un_name14",
	"Un_name15","Un_name16","Un_name17","Un_name18","Un_name19","Un_name20"}

/*var arr_alias = []string{"alias0","alias1","alias2","alias3","alias4","alias5","alias6",
	"alias7","alias8","alias9","alias10","alias11","alias12","alias13","alias14",
	"alias15","alias16","alias17","alias18","alias19","alias20"}
*/

func ReadCoilJsonMaker(a uint16,b []int,leng uint16) interface{}{

	d := map[string]int{}
	println("aa",a,b,leng)

	for i:=0;i<int(leng); i++{
		d[ReadCoil_alias[int(a)+i]] = b[i]
	}
	p,err := json.Marshal(d)
	if err!=nil{
		panic(err)
	}
	return p
}

func ReadCoilInJsonMaker(a uint16,b []int,leng uint16) interface{}{
	var p interface{}

	d := map[string]int{}
	println("aa",a,b,leng)

	for i:=0;i<int(leng); i++{
		d[ReadCoil_alias[int(a)+i]] = b[i]
	}
	p,err := json.Marshal(d)
	if err!=nil{
		panic(err)
	}
	return p
}

func ReadRegJsonMaker(a uint16,b []uint16,leng uint16) interface{}{
	var p interface{}

	d := map[string]uint16{}
	println("aa",a,b,leng)

	for i:=0;i<int(leng); i++{
		d[ReadReg_alias[int(a)+i]] = b[i]
	}
	p,err := json.Marshal(d)
	if err!=nil{
		panic(err)
	}
	return p
}

func ReadRegInJsonMaker(a uint16,b []uint16,leng uint16) interface{}{
	var p interface{}
	println(ReadReg_alias[0])
	println(ReadReg_alias[1])
	println(ReadReg_alias[2])
	keys := make([]string, 0)
	strArray := [...]string{"0ndia", "1anada", "2apan", "3Germany", "4Italy","saf","5fsdfas","6sdfsd","7qwre","8hnm","bwnsoi"}
	d := map[string]uint16{}
	var v uint16
	println("aa",a,b,leng)

	for i:=0;i<int(leng); i++{
		d[strArray[int(a)+i]] = b[i]
	}
	//keys := sliceOfKeys(d) // you'll have to implement this
	for _, k := range keys {
		v = d[k]
		// k is the key and v is the value; do your computation here
	}

	p,_ = json.Marshal(v)
	return p
}

func ErrJsonMaker() interface{}{
	j := ErrJson{"You entered it iolncorrectly."}
	p, err := json.Marshal(j)
	if err!=nil{
		panic(err)
	}
	return p
}
