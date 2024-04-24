package cassandra

import (
	"fmt"
	"log"

	"github.com/gocql/gocql"
)

func ConnectDatabase(url string, keyspace string) *gocql.Session {
	cc := gocql.NewCluster(url)

	cc.Keyspace = keyspace
	fmt.Println(cc)
	s, err := cc.CreateSession()
	if err != nil {
		log.Fatal(err)
	}
	return s
}
