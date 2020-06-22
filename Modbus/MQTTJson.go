package ModbusMQTT

import(
	"encoding/json"
	MQTT "github.com/eclipse/paho.mqtt.golang"
	
	)

var opts = MQTT.NewClientOptions().AddBroker("tcp://broker.hivemq.com:1883")

type ReadCoil struct {
	FunctionCode string
	StartAdd uint16
	AnalogState []int
}
type ErrJson struct {
	Errormessage string
}

func ReadCoilPublish(q byte,r bool,a uint16,b []int){
	j := ReadCoil{"DC Current\n", a,b}
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
	j := ErrJson{"You entered it iolncorrectly.Return to the Output Coils menu."}
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
