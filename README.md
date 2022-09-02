# Before runing this we need to make sure that following steps are performed
#### aerospike-server-community
#### aerospike-tools
## To run aerospike locally use this command 
#### sudo service aerospike start
## Get aerospike client for Go
#### go get github.com/aerospike/aerospike-client-go

## Now we are good run the program using below command
##### go run <name> -h <host> -p <port> -n <namespace> -s <set>
#### go run main.go -h localhost -p 3000 -n test -s myset


