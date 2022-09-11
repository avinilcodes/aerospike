package main

import (
	as "github.com/aerospike/aerospike-client-go/v6"
)

const (
	Host = "localhost"
	Port = 3000
)

func main() {
	client, err := as.NewClient(Host, Port)
	if err != nil {
		panic(err)
	}
	defer client.Close()
	policy := as.NewWritePolicy(0, 0)

	key, _ := as.NewKey("test", "myset", "mykey")
	key1, _ := as.NewKey("test", "myset", "mykey1")
	key2, _ := as.NewKey("test", "myset", "mykey2")

	bin := as.NewBin("mybin", "myvalue")

	client.PutBins(policy, key2, bin)
	bin1 := as.NewBin("name", "John")
	bin2 := as.NewBin("age", 25)

	client.PutBins(policy, key2, bin1, bin2)
	bin3 := as.NewBin("electronics", "Phone")
	client.PutBins(policy, key, bin3)

	bin4 := as.NewBin("electronics", "Tv")
	client.PutBins(policy, key1, bin4)

	bin5 := as.NewBin("electronics", "WM")
	client.PutBins(policy, key2, bin5)

	bin6 := as.NewBin("date", "25/06/2022")
	client.PutBins(policy, key2, bin6)
	bin7 := as.NewBin("Salary", 1000.0)
	client.PutBins(policy, key2, bin7)
	bin8 := as.NewBin("Currency", "USD")
	client.PutBins(policy, key2, bin8)
}
