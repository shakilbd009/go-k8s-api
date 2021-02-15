package cassandra

import "github.com/gocql/gocql"

var (
	session     *gocql.Session
	cassandraIP = "192.168.0.100"
	keyspace    = "k8s"
)

func init() {
	cluster := gocql.NewCluster(cassandraIP)
	cluster.Keyspace = keyspace
	cluster.Consistency = gocql.Quorum
	var err error
	session, err = cluster.CreateSession()
	if err != nil {
		//panic(err)
	}
}

func GetSession() *gocql.Session { return session }
