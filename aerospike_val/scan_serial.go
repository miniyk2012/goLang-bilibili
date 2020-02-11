package main

import (
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
		recordset, err := client.ScanNode(policy, node, "test", "demo")
		shared.PanicOnError(err)
	L:
		for {
			select {
			case rec := <-recordset.Records:
				if rec == nil {
					break L
				}
				//log.Println(rec.Key)
				metrics, exists := setMap[rec.Key.SetName()]

				if !exists {
					metrics = Metrics{}
				}
				metrics.count++
				metrics.total++
				setMap[rec.Key.SetName()] = metrics

			case err := <-recordset.Errors:
				// if there was an error, stop
				shared.PanicOnError(err)
			}
		}
		for k, v := range setMap {
			log.Println("Node ", node, " set ", k, " count: ", v.count)
			v.count = 0
		}

	}
	end := time.Now()
	seconds := float64(end.Sub(begin)) / float64(time.Second)
	log.Println("Elapsed time: ", seconds, " seconds")
}
