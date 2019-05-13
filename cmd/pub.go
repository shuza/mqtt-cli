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
	"github.com/shuza/mqtt-cli/key"
	"github.com/shuza/mqtt-cli/utils"
	"github.com/spf13/cobra"
	"os"
	"strconv"
)

// pubCmd represents the pub command
var pubCmd = &cobra.Command{
	Use:   "pub",
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
			if host == "" {
				panic(errors.New("Host address is missing. User -a or MQTT_HOST environment variable"))
			}
		}

		port, _ := cmd.Flags().GetInt("port")
		if port == 0 {
			value := os.Getenv(key.Port)
			if a, err := strconv.Atoi(value); err != nil {
				panic(errors.New("Port number is missing. use -p or MQTT_PORT environment variable"))
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
			if topic == "" {
				panic(errors.New("Topic is missing. User -t or MQTT_TOPIC environment variable"))
			}
		}

		qos, _ := cmd.Flags().GetInt("qos")
		if qos < 0 {
			qos, _ = strconv.Atoi(os.Getenv(key.Qos))
		}

		message, err := cmd.Flags().GetString("message")
		if err != nil {
			panic(errors.New("Required a message. use --message or -m"))
		}

		publish(host, port, clientId, topic, qos, message)
	},
}

func init() {
	rootCmd.AddCommand(pubCmd)
	pubCmd.Flags().StringP("message", "m", "", "Put your message here")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// pubCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// pubCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func publish(address string, port int, clientId string, topic string, qos int, message string) {
	client := createClient(address, port, clientId)
	token := client.Publish(topic, byte(qos), true, message)
	token.Wait()
	if token.Wait() && token.Error() != nil {
		fmt.Println("Error  :  ", token.Error())
	} else {
		fmt.Println("Topic :  ", topic)
		fmt.Println("QOS :  ", qos)
		fmt.Println("Message :  ", message)
		fmt.Println("published successfully")
	}
}
