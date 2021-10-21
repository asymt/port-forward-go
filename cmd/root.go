/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"io"
	"log"
	"net"
	"os"
	"strconv"

	"github.com/spf13/viper"
)

type ForwardItem struct {
	bindAddress string
	localPort   int16
	remoteHost  string
	remotePort  int16
}

var (
	cfgFile            string
	version            = "v1.0.0"
	commandForwardItem ForwardItem
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:     "port-forward",
	Short:   "port-forward cli",
	Version: version,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		if commandForwardItem.remoteHost == "" {
			log.Fatalln("remote host is not defined")
		}
		if commandForwardItem.remotePort == 0 {
			log.Fatalln("remote port is not defined")
		}
		server()
	},
}

func server() {
	sAddress := commandForwardItem.bindAddress + ":" + strconv.Itoa(int(commandForwardItem.localPort))
	dAddress := commandForwardItem.remoteHost + ":" + strconv.Itoa(int(commandForwardItem.remotePort))
	lis, err := net.Listen("tcp", sAddress)
	if err != nil {
		log.Printf("%v", err)
		return
	}
	log.Printf("port-forward started! %s->%s", sAddress, dAddress)
	for {
		log.Printf("%s ready to accept……", sAddress)
		sconn, err := lis.Accept()
		if err != nil {
			log.Panicf("create connection error:%v", err)
			continue
		}
		dconn, err := net.Dial("tcp", dAddress)
		if err != nil {
			log.Panicf("connect to %s fail:%v", dAddress, err)
		}
		log.Printf("forward data:%s<->%s", sconn.LocalAddr(), dconn.RemoteAddr())
		go func() {
			_, err = io.Copy(dconn, sconn)
			if err != nil {
				log.Panicf("send data to %v fail：%v", dAddress, err)
			}
		}()

		go func() {
			_, err = io.Copy(sconn, dconn)
			if err != nil {
				log.Panicf("recive data from %v fail：%v", dAddress, err)
			}
		}()
	}
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	cobra.OnInitialize(initConfig)
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	//rootCmd.PersistentFlags().StringVarP(&cfgFile, "config","c", "", "config file (default is $HOME/.port-forward.yaml)")
	rootCmd.PersistentFlags().StringVarP(&commandForwardItem.bindAddress, "bind-address", "b", "0.0.0.0", "bind address")
	rootCmd.PersistentFlags().StringVarP(&commandForwardItem.remoteHost, "remote-host", "r", "", "remote host")
	rootCmd.PersistentFlags().Int16VarP(&commandForwardItem.localPort, "local-port", "p", 9001, "local port")
	rootCmd.PersistentFlags().Int16VarP(&commandForwardItem.remotePort, "remote-port", "P", 0, "remote port")

	rootCmd.AddCommand(versionCmd)
	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	//rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of port-forward CLI",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("%s, version %s\n", rootCmd.Short, version)
	},
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	}
	viper.SetConfigType("json")
	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}
