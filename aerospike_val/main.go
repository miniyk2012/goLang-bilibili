package main

import (
	"fmt"
	"github.com/aerospike/aerospike-client-go/examples/shared"

	. "github.com/aerospike/aerospike-client-go"
)

// This is only for this example.
// Please handle errors properly.
func panicOnError(err error) {
	if err != nil {
		panic(err)
	}
}

func newClient() *Client {
	client, err := NewClient("localhost", 3000)
	panicOnError(err)
	return client
}

func localDemo() {
	// define a client to connect to
	client := newClient()
	namespace := "test"
	setName := "aerospike"
	key, err := NewKey(namespace, setName, "key") // user key can be of any supported type
	panicOnError(err)

	// define some bins
	bins := BinMap{
		"bin1": 42, // you can pass any supported type as bin value
		"bin2": "An elephant is a mouse with an operating system",
		"bin3": []interface{}{"Go", 17981},
	}

	// write the bins
	writePolicy := NewWritePolicy(0, 0)
	err = client.Put(writePolicy, key, bins)
	panicOnError(err)

	// read it back!
	readPolicy := NewPolicy()
	rec, err := client.Get(readPolicy, key)
	panicOnError(err)

	fmt.Printf("%#v\n", *rec)

	// Add to bin1
	err = client.Add(writePolicy, key, BinMap{"bin1": 5})  // 对数字的操作
	panicOnError(err)

	rec2, err := client.Get(readPolicy, key)
	panicOnError(err)

	fmt.Printf("value of %s: %v\n", "bin1", rec2.Bins["bin1"])

	// prepend and append to bin2
	err = client.Prepend(writePolicy, key, BinMap{"bin2": "Frankly:  "})  // 对字符串的操作
	panicOnError(err)
	err = client.Append(writePolicy, key, BinMap{"bin2": "."})
	panicOnError(err)

	rec3, err := client.Get(readPolicy, key)
	panicOnError(err)
	fmt.Printf("value of %s: %v\n", "bin2", rec3.Bins["bin2"])

	// delete bin3
	err = client.Put(writePolicy, key, BinMap{"bin3": nil})  // 删除某key的某一列
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

func prodTest() {
	client := newClient()
	fmt.Printf("%v\n", *client)

}
func main() {
	//localDemo()
	//prodTest()
	runExample(shared.Client)
}
