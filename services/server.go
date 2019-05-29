package main

import (
	"context"
	"fmt"
	"github.com/njdaniel/esim/services/models"
	"log"
	"net"
	"time"

	grpc "google.golang.org/grpc"
)

func main() {
	fmt.Println("Starting esim...")

	srv := grpc.NewServer()
	// var universe universeServer
	l, err := net.Listen("tcp", ":4040")
	if err != nil {
		log.Fatalf("could not listen to :4040: %v", err)
	}
	log.Fatal(srv.Serve(l))

	space := make(models.Space)
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
			ship := models.Ship{
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

type simServer struct{}

func (sim simServer) Describe(ctx context.Context) {

}
