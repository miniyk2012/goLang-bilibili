// https://github.com/aerospike/aerospike-client-go/blob/master/examples/list_map.go

package main

import (
	as "github.com/aerospike/aerospike-client-go"
	"log"
)
const Date string = "2020-02-11"

func scanAllUser() {
	client := newClient()
	recordset, err := client.ScanAll(nil, "devicePool2", "user")
	panicOnError(err)
	log.Printf("%s:\n", "user")
	var count = 0
	for res := range recordset.Results() {
		if res.Err != nil {
			// handle error; or close the recordset and break
			panicOnError(res.Err)
		}
		// process record
		//log.Printf("%v\n", res.Record.Bins["tags"].([]interface{}))
		log.Printf("%v\n", res.Record.Bins)
		count++
		if count > 10 {
			break
		}
	}
}

func scanAllVersion() {
	client := newClient()
	recordset, err := client.ScanAll(nil, "devicePool2", "version")
	panicOnError(err)
	log.Printf("%s:\n", "version")
	var count = 0
	for res := range recordset.Results() {
		if res.Err != nil {
			// handle error; or close the recordset and break
			panicOnError(res.Err)
		}
		// process record
		//log.Printf("%v\n", res.Record.Bins["tags"].([]interface{}))
		log.Printf("%v\n", res.Record.Bins)
		count++
		if count > 100 {
			break
		}
	}
}

func scanAllBehaviors(behavior string) {
	client := newClient()
	recordset, err := client.ScanAll(nil, "devicePool2", behavior)
	panicOnError(err)
	var count = 0
	log.Printf("%s:\n", behavior)
	for res := range recordset.Results() {
		if res.Err != nil {
			// handle error; or close the recordset and break
			panicOnError(res.Err)
		}
		// process record
		log.Printf("%v\n", res.Record.Bins)
		count++
		if count > 10 {
			break
		}
	}
}

func scanUsers(users []string) {
	client := newClient()
	readPolicy := as.NewPolicy()
	for _, tuid := range users {
		key1, err1 := as.NewKey("devicePool2", "user", tuid)
		panicOnError(err1)
		key2, err2 := as.NewKey("devicePool2", "version", tuid)
		panicOnError(err2)

		rec1, err3 := client.Get(readPolicy, key1)
		panicOnError(err3)
		log.Printf("%s:, %v\n", tuid, rec1.Bins)
		//log.Printf("%s:, %v\n", tuid, rec1.Bins["tags"].([]interface{}))
		log.Println()
		rec2, err4 := client.Get(readPolicy, key2)
		panicOnError(err4)
		log.Printf("%s:, %v\n", tuid, rec2.Bins)
		log.Println("--------------------")
	}
}



func scanClick(users []string) {
	client := newClient()
	readPolicy := as.NewPolicy()
	for _, tuid := range users {
		log.Printf("click key=%s", tuid+Date)
		key, err := as.NewKey("devicePool2", "click", tuid+Date)
		panicOnError(err)

		rec, err := client.Get(readPolicy, key)
		if err != nil {
			continue
		}
		log.Printf("%s:, %v\n", tuid, rec.Bins)
		log.Println("--------------------")
	}

}

func scanShow(users []string) {
	client := newClient()
	readPolicy := as.NewPolicy()
	for _, tuid := range users {
		log.Printf("show key=%s", tuid+Date)
		key, err := as.NewKey("devicePool2", "show", tuid+Date)
		panicOnError(err)

		rec, err := client.Get(readPolicy, key)
		if err != nil {
			continue
		}
		log.Printf("%s:, %v\n", tuid, rec.Bins)
		log.Println("--------------------")
	}
}

func scanCvr(users []string) {
	client := newClient()
	readPolicy := as.NewPolicy()
	for _, tuid := range users {
		log.Printf("cvr key=%s", tuid+Date)
		key, err := as.NewKey("devicePool2", "cvr", tuid+Date)
		panicOnError(err)

		rec, err := client.Get(readPolicy, key)
		if err != nil {
			continue
		}
		log.Printf("%s:, %v\n", tuid, rec.Bins)
		log.Println("--------------------")
	}
}