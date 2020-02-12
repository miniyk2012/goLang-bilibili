package main

import (
	as "github.com/aerospike/aerospike-client-go"
)

// This is only for this example.
// Please handle errors properly.
func panicOnError(err error) {
	if err != nil {
		panic(err)
	}
}

func newClient() *as.Client {
	client, err := as.NewClient("localhost", 3000)
	panicOnError(err)
	return client
}


func main() {
	//localDemo()
	//runExample(shared.Client)

	//scanAllUser()
	//log.Println("---------------------------------------------------------------------------")
	//scanAllVersion()
	//log.Println("---------------------------------------------------------------------------")
	//scanAllBehaviors("click")
	//log.Println("---------------------------------------------------------------------------")
	//scanAllBehaviors("show")
	//log.Println("---------------------------------------------------------------------------")
	//scanAllBehaviors("cvr")
	//scanUsers([]string{"--ktYSAvQWekDQUFGC21Sg"})
	//scanClick([]string{"--DWyUnie5wxFwJdeMIuag"})
	scanCvr([]string{"-02r0bcUQ92QCYU7dM_S1A"})
	//scanShow([]string{"-ggHtd993qKScx93HieoRQ"})
}
