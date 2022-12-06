package cassandra

import (
	"fmt"

	"github.com/gocql/gocql"
)

var (
	cluster *gocql.ClusterConfig
)

func init() {
	// Connect to Cassandra cluster:
	//Not redeclare cluster
	cluster = gocql.NewCluster("127.0.0.1")
	cluster.Keyspace = "oauth"
	cluster.Consistency = gocql.Quorum

	fmt.Println("cassandra connection successfully created")
}

func GetSession() (*gocql.Session, error) {
	return cluster.CreateSession()
}
