package main

import (
	"fmt"
	"github.com/aerospike/aerospike-client-go/examples/shared"
	"log"
	"time"

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

func localDemo() {
	// define a client to connect to
	client := newClient()
	namespace := "test"
	setName := "aerospike"
	key, err := as.NewKey(namespace, setName, "key") // user key can be of any supported type
	panicOnError(err)

	// define some bins
	bins := as.BinMap{
		"bin1": 42, // you can pass any supported type as bin value
		"bin2": "An elephant is a mouse with an operating system",
		"bin3": []interface{}{"Go", 17981},
	}

	// write the bins
	writePolicy := as.NewWritePolicy(0, 0)
	err = client.Put(writePolicy, key, bins)
	panicOnError(err)

	// read it back!
	readPolicy := as.NewPolicy()
	rec, err := client.Get(readPolicy, key)
	panicOnError(err)

	fmt.Printf("%#v\n", *rec)

	// Add to bin1
	err = client.Add(writePolicy, key, as.BinMap{"bin1": 5}) // 对数字的操作
	panicOnError(err)

	rec2, err := client.Get(readPolicy, key)
	panicOnError(err)

	fmt.Printf("value of %s: %v\n", "bin1", rec2.Bins["bin1"])

	// prepend and append to bin2
	err = client.Prepend(writePolicy, key, as.BinMap{"bin2": "Frankly:  "}) // 对字符串的操作
	panicOnError(err)
	err = client.Append(writePolicy, key, as.BinMap{"bin2": "."})
	panicOnError(err)

	rec3, err := client.Get(readPolicy, key)
	panicOnError(err)
	fmt.Printf("value of %s: %v\n", "bin2", rec3.Bins["bin2"])

	// delete bin3
	err = client.Put(writePolicy, key, as.BinMap{"bin3": nil}) // 删除某key的某一列
	rec4, err := client.Get(readPolicy, key)
	panicOnError(err)
	fmt.Printf("bin3 does not exist anymore: %#v\n", *rec4)

	// check if key exists
	exists, err := client.Exists(readPolicy, key)
	panicOnError(err)
	fmt.Printf("key exists in the database: %#v\n", exists)

	// delete the key, and check if key exists
	existed, err := client.Delete(writePolicy, key)
	panicOnError(err)
	fmt.Printf("did key exist before delete: %#v\n", existed)

	exists, err = client.Exists(readPolicy, key)
	panicOnError(err)
	fmt.Printf("key exists in the database: %#v\n", exists)
}
