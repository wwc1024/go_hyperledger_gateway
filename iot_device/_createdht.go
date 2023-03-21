package iot_device

import (
	"fmt"

	"go_gateway/public"

	dht "github.com/d2r2/go-dht"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/channel"
	"github.com/robfig/cron/v3"
)

func getDeviceDate() (float32, float32) {
	temperature, humidity, _, err :=
		dht.ReadDHTxxWithRetry(dht.DHT11, 4, false, 1)
	if err != nil || temperature == -1 {
		fmt.Println("get Dates failed!")
	}
	return temperature, humidity
}

//float32 转 String工具类，保留6位小数
// func FloatToString(input_num float32) string {
// 	// to convert a float number to a string
// 	return strconv.FormatFloat(float64(input_num), 'f', 6, 64)
// }

func GetDate() (string, string) {
	//得到温湿度
	temperatureNum, humidityNum := getDeviceDate()
	for temperatureNum == -1 {
		temperatureNum, humidityNum = getDeviceDate()
	}
	temp := fmt.Sprint(int(temperatureNum))
	hum := fmt.Sprint(int(humidityNum))
	fmt.Println("Dates are ", temp, ",", hum)
	return temp, hum
}

//执行create的合约
func CreateContent() string {
	channelclient := FabsdkConfFile()
	//设置第几台机器 "dht2"
	temp, hum := GetDate()
	name := "dht" + fmt.Sprint(public.DhtId)
	fmt.Printf("---------%v当前采样频率%v----------\n", name, public.NowSample)
	req := channel.Request{
		ChaincodeID: "mycc",      //合约名字
		Fcn:         "CreateDht", //方法
		Args:        [][]byte{[]byte(name), []byte(temp), []byte(hum)},
	}
	res, err := channelclient.Execute(req)
	if err != nil {
		fmt.Println(err)
	}
	return fmt.Sprint(res.TransactionID)
}

func DhtCronStart() {
	c := cron.New()
	fmt.Println("开启create循环:每10m执行一次")

	str := "@every " + fmt.Sprint(public.NowSample) + "m"
	c.AddFunc(str, func() {
		// TestCmd()
		transactionID := CreateContent()
		fmt.Printf("tx: %v\n", transactionID)
		//存储到tx到数据库，前端读取
	})
	c.Start()

}
