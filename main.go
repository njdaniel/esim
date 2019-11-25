// Copyright © 2019 NAME HERE <EMAIL ADDRESS>
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

package main

import (
	"fmt"
	main2 "github.com/njdaniel/esim/services/server"
	"log"
	"math"
	"time"

	flag "github.com/spf13/pflag"
	"github.com/spf13/viper"
)

const (
	RadToDeg = 180 / math.Pi
	DegToRad = math.Pi / 180
)

func init() {
	log.Println("Checking for config file...")
	var configFile *string
	configFile = flag.StringP("config", "c", "", "location of config file")
	flag.Parse()
	viper.SetConfigFile(*configFile)
	err := viper.ReadInConfig()
	if err != nil {
		log.Printf("could not read config file: %v \n", err)
	} else {
		fmt.Printf("Config %v read in", *configFile)
	}
}

func main() {
	// cmd.Execute()

	// TODO: create space
	// TODO: create ship and add to space
	// TODO: ctrl ship via cli
	// TODO: print out ship traits
	// TODO: PRIORITY leave program running
	fmt.Println("Starting esim...")

	// srv := grpc.NewServer()
	// // var universe universeServer
	// l, err := net.Listen("tcp", ":4040")
	// if err != nil {
	// 	log.Fatalf("could not listen to :4040: %v", err)
	// }
	// log.Fatal(srv.Serve(l))

	space := make(main2.Space)
	//var wg sync.WaitGroup

	// Update state every 1/s =========================================
	go func() {
		for range time.Tick(time.Second) {
			//TODO: Update the position of each object in space
			for _, x := range space {
				x.UpdateLocation()
			}
			//time.Sleep(time.Second * 5)
		}
	}()

	// Wait for input by cli ==========================================
	fmt.Print("Enter cmd: \n")
	c := ""
	for {

		fmt.Scanln(&c)
		// fmt.Println(c)
		if c == "create" {
			fmt.Println("Creating...")
			ship := main2.Ship{
				Name:        "rifter",
				MaxVelocity: 180,
			}
			space["rifter"] = &ship
			fmt.Println("done.")
		} else if c == "full" {
			if s, ok := space["rifter"]; ok {
				s.FullSpeed()
				fmt.Printf("current speed %f \n", s.CurVelocity)
			} else {
				fmt.Println("doesnt exist")
			}
		} else if c == "location" {
			fmt.Printf("Current location %v \n", space["rifter"].Location)
		} else if c == "describe" {
			//fmt.Printf("Ship stats: %v", space["rifter"])
			space["rifter"].Print()
		} else if c == "exit" {
			break
		} else {
			fmt.Printf("command not recognized: %s \n", c)
		}
		//fmt.Printf("space: %v \n", space)
	}
}
