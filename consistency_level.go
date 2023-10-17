package main

import (
	"fmt"
	"log"
	"time"

	"github.com/gocql/gocql"
)

func insertWithConsistency(cluster *gocql.ClusterConfig, consistency gocql.Consistency, key string) {
	cluster.Consistency = consistency
	session, err := cluster.CreateSession()
	if err != nil {
		log.Fatalf("Failed to create session: %v", err)
	}
	defer session.Close()

	query := `INSERT INTO payment.test (key, value) VALUES (?, ?)`
	if err := session.Query(query, key, "value").Exec(); err != nil {
		log.Printf("Failed to insert with consistency %v: %v", consistency, err)
	} else {
		log.Printf("Successfully inserted with consistency %v", consistency)
	}
}

func main() {
	// Assuming you have nodes running on localhost at ports 9042 and 9043 for two different datacenters
	cluster := gocql.NewCluster("127.0.0.1:9042", "127.0.0.1:9043")
	cluster.Keyspace = "payment"

	// Test with QUORUM
	keyForQuorum := fmt.Sprintf("test_key_quorum_%d", time.Now().UnixNano())
	insertWithConsistency(cluster, gocql.Quorum, keyForQuorum)

	// Test with LOCAL_QUORUM
	keyForLocalQuorum := fmt.Sprintf("test_key_local_quorum_%d", time.Now().UnixNano())
	insertWithConsistency(cluster, gocql.LocalQuorum, keyForLocalQuorum)
}
