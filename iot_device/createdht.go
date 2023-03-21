package iot_device

import (
	"fmt"
	"go_gateway/public"
	"strconv"
	"time"

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
func FloatToString(input_num float32) string {
	// to convert a float number to a string
	return strconv.FormatFloat(float64(input_num), 'f', 6, 64)
}

func GetDate() (string, string) {
	//得到温湿度
	retrytime := 3
	temperatureNum, humidityNum := getDeviceDate()
	for temperatureNum == -1 && retrytime >= 0 {
		temperatureNum, humidityNum = getDeviceDate()
		retrytime--
	}
	retrytime = 0
	temp := FloatToString(temperatureNum)
	hum := FloatToString(humidityNum)
	fmt.Println("Dates are ", temp, ",", hum)
	return temp, hum
}

//执行create的合约
func CreateContent() string {
	channelclient := FabsdkConfFile()
	//设置第几台机器 "dht2"
	now := time.Now().Format("2006-01-02 15:04:05")
	name := fmt.Sprint(public.DhtId)
	temp, hum := GetDate()
	countdhtnumber := fmt.Sprint(Countdhtnumber)
	req := channel.Request{
		ChaincodeID: "mycc",      //合约名字
		Fcn:         "CreateDht", //方法
		Args:        [][]byte{[]byte(countdhtnumber), []byte(name), []byte(now), []byte(temp), []byte(hum)},
	}
	if temp == "-1.000000" {
		return "传感器不稳定"
	} else {
		res, err := channelclient.Execute(req)
		if err != nil {
			panic(err)
		}
		return "End Createcontent:" + fmt.Sprint(res.TransactionID)
	}
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
