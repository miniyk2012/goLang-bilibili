// https://github.com/aerospike/aerospike-client-go/blob/master/examples/list_map.go

package main

import (
	as "github.com/aerospike/aerospike-client-go"
	"log"
)
const Date string = "2020-02-12"
const NameSpace = "devicePool2"

func scanAllUser() {
	client := newClient()
	recordset, err := client.ScanAll(nil, NameSpace, "user")
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
	recordset, err := client.ScanAll(nil, NameSpace, "version")
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
	recordset, err := client.ScanAll(nil, NameSpace, behavior)
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
		log.Printf("user key=%s", tuid)
		key1, err1 := as.NewKey(NameSpace, "user", tuid)
		panicOnError(err1)
		rec1, err3 := client.Get(readPolicy, key1)
		panicOnError(err3)
		log.Printf("%s:, %v\n", tuid, rec1.Bins)
		//log.Printf("%s:, %v\n", tuid, rec1.Bins["tags"].([]interface{}))
		log.Println("--------------------")
	}
}

func scanVersions(users []string) {
	client := newClient()
	readPolicy := as.NewPolicy()
	for _, tuid := range users {
		log.Printf("version key=%s", tuid)
		key, err1 := as.NewKey(NameSpace, "version", tuid)
		panicOnError(err1)
		rec, err2 := client.Get(readPolicy, key)
		panicOnError(err2)
		log.Printf("%s:, %v\n", tuid, rec.Bins)
		log.Println("--------------------")
	}
}




func scanClick(users []string) {
	client := newClient()
	readPolicy := as.NewPolicy()
	for _, tuid := range users {
		log.Printf("click key=%s", tuid+Date)
		key, err := as.NewKey(NameSpace, "click", tuid+Date)
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
		key, err := as.NewKey(NameSpace, "show", tuid+Date)
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
		key, err := as.NewKey(NameSpace, "cvr", tuid+Date)
		panicOnError(err)

		rec, err := client.Get(readPolicy, key)
		if err != nil {
			continue
		}
		log.Printf("%s:, %v\n", tuid, rec.Bins)
		log.Println("--------------------")
	}
}

func emptyTag(tuid string) {
	log.Printf("empty tuid=%s", tuid)
	client := newClient()
	readPolicy := as.NewPolicy()
	writePolicy := as.NewWritePolicy(0, 0)

	key, err := as.NewKey(NameSpace, "user", tuid)
	panicOnError(err)
	record, err := client.Get(readPolicy, key, "tags")
	panicOnError(err)
	receivedList := record.Bins["tags"].([]interface{})
	log.Printf("before empty tags=%v\n",receivedList)

	// empty tags
	empty_tags := []int{}
	bin := as.NewBin("tags", empty_tags)
	client.PutBins(writePolicy ,key, bin)

	record, err = client.Get(readPolicy, key, "tags")
	panicOnError(err)
	receivedList = record.Bins["tags"].([]interface{})
	log.Printf("after empty tags=%v\n",receivedList)
}
