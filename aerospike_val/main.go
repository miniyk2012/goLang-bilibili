package main

import (
	as "github.com/aerospike/aerospike-client-go"
	"os"
	"log"
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
	people_type := os.Args[1]
	tuid := os.Args[2]
	log.Printf("type=%s, tuid=%s", people_type, tuid)
	users := []string{tuid}
	switch people_type {
	case "cvr":
		scanCvr(users)
	case "user":
		scanUsers(users)
	case "click":
		scanClick(users)
	case "show":
		scanShow(users)
	case "version":
		scanVersions(users)
	case "emptyTag":
		emptyTag(tuid)
	default:
		log.Printf("%s not support", people_type)
	}
	//scanAllUser()
	//log.Println("---------------------------------------------------------------------------")
	//scanAllVersion()
	//log.Println("---------------------------------------------------------------------------")
	//scanAllBehaviors("click")
	//log.Println("---------------------------------------------------------------------------")
	//scanAllBehaviors("show")
	//log.Println("---------------------------------------------------------------------------")
	//scanAllBehaviors("cvr")
}
