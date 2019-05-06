// Copyright © 2019 Shalauddin Ahamad Shuza <shuza.sa@gmail.com>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"errors"
	"fmt"
	"github.com/eclipse/paho.mqtt.golang"
	"github.com/spf13/cobra"
	"mqtt-sh/key"
	"mqtt-sh/utils"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"
)

// subCmd represents the sub command
var subCmd = &cobra.Command{
	Use:   "sub",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		host, _ := cmd.Flags().GetString("address")
		if host == "" {
			host = os.Getenv(key.Host)
		}

		port, _ := cmd.Flags().GetInt("port")
		if port == 0 {
			value := os.Getenv(key.Port)
			if a, err := strconv.Atoi(value); err != nil {
				panic(errors.New("Port number is missing. use -p or MQTT_PORT environment veriable"))
			} else {
				port = a
			}
		}

		clientId, _ := cmd.Flags().GetString("clientId")
		if clientId == "" {
			clientId = os.Getenv(key.ClientId)
			if clientId == "" {
				clientId = utils.NewClientId()
				os.Setenv(key.ClientId, clientId)
			}
		}

		topic, _ := cmd.Flags().GetString("topic")
		if topic == "" {
			topic = os.Getenv(key.Topic)
		}

		qos, _ := cmd.Flags().GetInt("qos")
		if qos < 0 {
			qos, _ = strconv.Atoi(os.Getenv(key.Qos))
		}

		subscribe(host, port, clientId, topic, qos)
	},
}

func init() {
	rootCmd.AddCommand(subCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// subCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// subCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func subscribe(address string, port int, clientId string, topic string, qos int) {
	client := createClient(address, port, clientId)
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	client.Subscribe(topic, byte(qos), func(client mqtt.Client, message mqtt.Message) {
		fmt.Printf("=======		Received	=======\nTopic  ::  %v\nMessage  ::  %v\n", message.Topic(), string(message.Payload()))
	})
	fmt.Println("Broker Address :  ", address)
	fmt.Println("Client ID :  ", clientId)
	fmt.Println("Topic :  ", topic)
	fmt.Println("QOS :  ", qos)
	fmt.Println("Subscribed....")
	<-sigs
}

func createClient(address string, port int, clientId string) mqtt.Client {
	ops := mqtt.NewClientOptions()
	ops.AddBroker(fmt.Sprintf("tcp://%s:%d", address, port))
	ops.SetClientID(clientId)

	client := mqtt.NewClient(ops)
	token := client.Connect()
	for !token.WaitTimeout(3 * time.Second) {
	}
	if err := token.Error(); err != nil {
		panic(err)
	}

	return client
}
