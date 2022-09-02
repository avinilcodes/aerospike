Install these ->
aerospike-server-community
aerospike-tools
then -> 
sudo service aerospike start
then -> 
go get github.com/aerospike/aerospike-client-go
cd $GOPATH/src/github.com/aerospike/aerospike-client-go/examples
go run <name> -h <host> -p <port> -n <namespace> -s <set>
