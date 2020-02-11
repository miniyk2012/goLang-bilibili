package main

import (
	"fmt"
	"log"
	"time"

	"github.com/aerospike/aerospike-client-go/examples/shared"

	as "github.com/aerospike/aerospike-client-go"
)

type Metrics struct {
	count int
	total int
}

var setMap = make(map[string]Metrics)

func runExample(client *as.Client) {
	log.Println("Scan series: namespace=", *shared.Namespace, " set=", *shared.Set)
	policy := as.NewScanPolicy()
	policy.MaxRetries = 1
	policy.Priority = as.LOW

	nodeList := client.GetNodes()
	begin := time.Now()

	for _, node := range nodeList {
		log.Println("Scan node ", node.GetName())
		recordset, err := client.ScanNode(policy, node, *shared.Namespace, *shared.Set)
		shared.PanicOnError(err)
		fmt.Printf("%v\n", *recordset)

	}
	end := time.Now()
	seconds := float64(end.Sub(begin)) / float64(time.Second)
	log.Println("Elapsed time: ", seconds, " seconds")
}
