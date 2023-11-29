package main

import (
	restclient "enigmacamp.com/be-enigma-laundry/rest-client"
)

func main() {
	//delivery.NewServer().Run()
	//restclient.SampleRestClientBuiltIn()
	restclient.SampleRestClientResty()
}
