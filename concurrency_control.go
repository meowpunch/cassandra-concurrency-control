package main

import (
	"fmt"
	"github.com/gocql/gocql"
	"log"
	"sync"
)

func InsertConcurrently(session *gocql.Session, dcName string, wg *sync.WaitGroup) {
	defer wg.Done()

	var innerWg sync.WaitGroup
	for i := 0; i < 10; i++ {
		conversationID := fmt.Sprintf("conv_id_%d", i)
		innerWg.Add(1)
		go func(id string) {
			defer innerWg.Done()
			applied, err := InsertRowWithLWT(session, id)
			if err != nil {
				log.Printf("Failed to execute LWT in data center %s for conversation_id %s: %v", dcName, id, err)
				return
			}
			if !applied {
				log.Printf("LWT failed in data center %s for conversation_id %s.", dcName, id)
			} else {
				log.Printf("LWT succeeded in data center %s for conversation_id %s.", dcName, id)
			}
		}(conversationID)
	}
	innerWg.Wait()
}

func main() {
	var wg sync.WaitGroup

	dc1Config := CassandraConfig{Port: 9042, Keyspace: "payment"}
	dc2Config := CassandraConfig{Port: 9043, Keyspace: "payment"}

	dc1Session := NewSession(dc1Config)
	defer CloseSession(dc1Session)

	dc2Session := NewSession(dc2Config)
	defer CloseSession(dc2Session)

	wg.Add(2)
	go InsertConcurrently(dc1Session, "dc1", &wg)
	go InsertConcurrently(dc2Session, "dc2", &wg)

	wg.Wait()
}
