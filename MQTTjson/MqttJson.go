package ModbusMQTT

import(
	"encoding/json"
	MQTT "github.com/eclipse/paho.mqtt.golang"
)

var opts = MQTT.NewClientOptions().AddBroker("tcp://broker.hivemq.com:1883")

type ReadCoil struct {
	FunctionCode string
	StartAdd uint16
	DigitalState []int
}
type WriteCoils struct {
	FunctionCode string
	StartAdd uint16
	DigitalState []string
}

type ReadReg struct {
	FunctionCode string
	StartAdd uint16
	AnalogState []uint16
}

type WriteRegs struct {
	FunctionCode string
	StartAdd uint16
	AnalogState []string
}

type ErrJson struct {
	Errormessage string
}

func ReadCoilPublish(q byte,r bool,a uint16,b []int){
	j := ReadCoil{"ReadCoil", a,b}
	p, err := json.Marshal(j)
	if err!=nil{
		panic(err)
	}

	c := MQTT.NewClient(opts)
	if token := c.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}
	token := c.Publish("test/topic12/1",q,r,p); token.Wait()
}

func ReadCoilInPublish(q byte,r bool,a uint16,b []int){
	j := ReadCoil{"ReadCoilIn", a,b}
	p, err := json.Marshal(j)
	if err!=nil{
		panic(err)
	}

	c := MQTT.NewClient(opts)
	if token := c.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}
	token := c.Publish("test/topic12/1",q,r,p); token.Wait()
}

func ReadRegPublish(q byte,r bool,a uint16,b []uint16){
	j := ReadReg{"ReadRegister", a,b}
	p, err := json.Marshal(j)
	if err!=nil{
		panic(err)
	}

	c := MQTT.NewClient(opts)
	if token := c.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}
	token := c.Publish("test/topic12/1",q,r,p); token.Wait()
}

func ReadRegInPublish(q byte,r bool,a uint16,b []uint16){
	j := ReadReg{"ReadInputRegister", a,b}
	p, err := json.Marshal(j)
	if err!=nil{
		panic(err)
	}

	c := MQTT.NewClient(opts)
	if token := c.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}
	token := c.Publish("test/topic12/1",q,r,p); token.Wait()
}

func WriteCoilsPublish(q byte,r bool,a uint16,b []string){
	j := WriteCoils{"WriteCoils", a,b}
	p, err := json.Marshal(j)
	if err!=nil{
		panic(err)
	}

	c := MQTT.NewClient(opts)
	if token := c.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}
	token := c.Publish("test/topic12/1",q,r,p); token.Wait()
}

func WriteRegsPublish(q byte,r bool,a uint16,b []string){
	j := WriteRegs{"WriteRegisters", a,b}
	p, err := json.Marshal(j)
	if err!=nil{
		panic(err)
	}

	c := MQTT.NewClient(opts)
	if token := c.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}
	token := c.Publish("test/topic12/1",q,r,p); token.Wait()
}

func ErrPublish(){
	j := ErrJson{"You entered it iolncorrectly."}
	p, err := json.Marshal(j)
	if err!=nil{
		panic(err)
	}

	c := MQTT.NewClient(opts)
	if token := c.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}
	token := c.Publish("test/topic12/1",0,true,p); token.Wait()
}
