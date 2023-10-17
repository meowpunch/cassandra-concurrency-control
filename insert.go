package main

import (
	"fmt"
	"log"
	"sync"

	"github.com/gocql/gocql"
)

func insertRow(session *gocql.Session, conversationID string, dcName string, wg *sync.WaitGroup) {
	defer wg.Done()

	query := `INSERT INTO payment_id_by_conversation_id (conversation_id, payment_id) VALUES (?, ?) IF NOT EXISTS`
	var applied bool
	var existingConvID string
	var existingPaymentID gocql.UUID
	applied, err := session.Query(query, conversationID, gocql.TimeUUID()).ScanCAS(&existingConvID, &existingPaymentID)
	if err != nil {
		log.Printf("Failed to execute LWT in data center %s for conversation_id %s: %v", dcName, conversationID, err)
		return
	}

	if !applied {
		log.Printf("LWT failed in data center %s for conversation_id %s: row with the same conversation_id already exists (existing payment_id: %s).", dcName, conversationID, existingPaymentID)
	} else {
		log.Printf("LWT succeeded in data center %s for conversation_id %s: row was inserted.", dcName, conversationID)
	}
}

func insertData(port int, dcName string, wg *sync.WaitGroup) {
	defer wg.Done()

	// Connect to Cassandra
	cluster := gocql.NewCluster(fmt.Sprintf("127.0.0.1:%d", port))
	cluster.Keyspace = "payment"
	cluster.Consistency = gocql.Quorum
	session, err := cluster.CreateSession()
	if err != nil {
		log.Fatalf("Failed to connect to Cassandra at port %d: %v", port, err)
	}
	defer session.Close()

	log.Printf("Successfully created session for data center: %s", dcName)

	var innerWg sync.WaitGroup
	for i := 0; i < 10; i++ {
		conversationID := fmt.Sprintf("conv_id_%d", 1)
		innerWg.Add(1)
		go insertRow(session, conversationID, dcName, &innerWg)
	}
	innerWg.Wait()
}

func main() {
	var wg sync.WaitGroup

	// Start two goroutines for concurrent inserts
	wg.Add(2)
	go insertData(9042, "dc1", &wg) // Port for cassandra-dc1-node1
	go insertData(9043, "dc2", &wg) // Port for cassandra-dc2-node1

	wg.Wait()
}
