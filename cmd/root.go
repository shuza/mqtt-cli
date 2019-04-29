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
	"os"

	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"time"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "mqtt-sh",
	Short: "Command line communication with MQTT broker",
	Long:  `MQTT-sh is a CLI application to subscribe and publish messages to MQTT broker. First you have to initialize mqtt-sh application with broker credentials like IP, clientID etc. Then you can subscribe or publish messages to any topic.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	//	Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.mqtt-sh.yaml)")

	rootCmd.PersistentFlags().StringP("address", "a", "", "Set MQTT broker host with port example 192.168.0.1:1883")
	rootCmd.PersistentFlags().StringP("clientId", "i",
		fmt.Sprintf("mqtt-sh-%v", time.Now().Nanosecond()),
		"Set your clientID")
	rootCmd.PersistentFlags().StringP("topic", "t", "/*", "Set the topic at which you want to subscribe and publish messages")
	rootCmd.PersistentFlags().IntP("qos", "q", 1, "Set your quality of service ")

}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".mqtt-sh" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".mqtt-sh")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
