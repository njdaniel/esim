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
	"log"
	"net"

	"github.com/njdaniel/esim/cmd"
	grpc "google.golang.org/grpc"
)

func main() {
	cmd.Execute()
	// TODO: create space
	// TODO: create ship and add to space
	// TODO: ctrl ship via cli
	// TODO: print out ship traits
	// TODO: PRIORITY leave program running
	fmt.Println("Starting esim...")
	srv := grpc.NewServer()
	// var universe universeServer
	l, err := net.Listen("tcp", ":4040")
	if err != nil {
		log.Fatalf("could not listen to :4040: %v", err)
	}
	log.Fatal(srv.Serve(l))
}

type universeServer struct{}
