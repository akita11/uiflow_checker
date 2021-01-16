// UIFlow MQTT checker v1.0 by akita11 (akita@ifdl.jp)
// based on MQTT connection checker, https://github.com/ciniml/mqtt-checker
// License: Boost Software License 1.0
// to build executable binary: go build

package main

import (
	"fmt"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

const testServer = "flow.m5stack.com:1883"
const timeoutSeconds = 10

func innerMain() error {
        fmt.Println("M5Stack UIFlow MQTT checker v1.0 (akita@ifdl.jp)")
        fmt.Println("チェック開始 / Check start")
	messageCh := make(chan mqtt.Message)
	var handler mqtt.MessageHandler = func(client mqtt.Client, message mqtt.Message) {
		messageCh <- message
	}

	options := mqtt.NewClientOptions()
	options.AddBroker("tcp://" + testServer)
	client := mqtt.NewClient(options)

//	fmt.Printf("Connecting to %s... \n", testServer)
	if token := client.Connect(); !token.WaitTimeout(time.Second * timeoutSeconds) {
		return fmt.Errorf("エラーError: Connection failed: Timed out")
	} else if err := token.Error(); err != nil {
		return fmt.Errorf("エラーError: Connection failed: %s", err)
	}

	subscription := client.Subscribe("test", 0, handler)
	if !subscription.WaitTimeout(time.Second * timeoutSeconds) {
		return fmt.Errorf("エラーError: Subscription failed: Timed out")
	} else if err := subscription.Error(); err != nil {
		return fmt.Errorf("エラーError: Subscription failed: %s", err)

	}

//	fmt.Println("Publishing some data... ")
	if token := client.Publish("test", 0, false, "test"); !token.WaitTimeout(time.Second * timeoutSeconds) {
		return fmt.Errorf("エラーError: Publish failed: Timed out")
	} else if err := token.Error(); err != nil {
		return fmt.Errorf("エラーError: Publish failed: %s", err)
	}

//	fmt.Println("Waiting for receiving some data... ")
//	timer := time.NewTimer(time.Second * timeoutSeconds)
//	select {
//	case message := <-messageCh:
//                fmt.Printf("データ受信OK: %s", message.Topic())
//	case <-timer.C:
//		return fmt.Errorf("エラー: Subscription data timed out")
//	}

        fmt.Println("チェック終了。途中でエラーが出ていなければOKです Check done. If no error appears, your WiFi is OK for UIFlow.");
	return nil
}

func main() {
	if err := innerMain(); err != nil {
		fmt.Printf("Error: %s\n", err)
	}
	fmt.Println("Enterキーを押すと終了します。 Press Enter to exit.")
	fmt.Scanln()
}
