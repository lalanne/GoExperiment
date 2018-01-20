package main

import "fmt"

func main() {
	service := NewHello_PortType("http://schemas.xmlsoap.org/soap/envelope/", false)
	amIthere, err := service.SayHello(&areYouThere{})
	if err != nil {
		panic(err)
	}

	fmt.Printf("Alive?: %t\n", amIthere.Return_)

	/*stations, err := service.GetStations(&getStations{})
	if err != nil {
		panic(err)
	}
	fmt.Printf("Stations: %v\n", stations.Return_)*/
}
