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
	"fmt"
	"github.com/spf13/cobra"
	"mqtt-sh/db"
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
		db.Client = &db.CacheClient{}
		db.Client.Init()
		defer db.Client.Close()

		address, _ := cmd.Flags().GetString("address")
		if address == "" {
			address = db.Client.Get(db.Address)
		}

		clientId, _ := cmd.Flags().GetString("clientId")
		if clientId == "" {
			clientId = db.Client.Get(db.ClientId)
		}

		topic, _ := cmd.Flags().GetString("topic")
		if topic == "" {
			topic = db.Client.Get(db.Topic)
		}

		qos, _ := cmd.Flags().GetInt("qos")
		if qos < 0 {
			qos, _ = strconv.Atoi(db.Client.Get(db.Qos))
		}

		db.Client.Close()

		message, _ := cmd.Flags().GetString("message")

		publish(address, clientId, topic, qos, message)
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

func publish(address string, clientId string, topic string, qos int, message string) {
	client := createClient(address, clientId)
	token := client.Publish(topic, byte(qos), true, message)
	if err := token.Error(); err != nil {
		fmt.Println("Error  :  ", err)
	} else {
		fmt.Println("published successfully")
	}
}
