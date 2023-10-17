package main

import (
	"fmt"
	"log"

	"github.com/gocql/gocql"
)

type CassandraConfig struct {
	Port     int
	Keyspace string
}

func NewSession(config CassandraConfig) *gocql.Session {
	cluster := gocql.NewCluster(fmt.Sprintf("127.0.0.1:%d", config.Port))
	cluster.Keyspace = config.Keyspace
	cluster.Consistency = gocql.Quorum
	session, err := cluster.CreateSession()
	if err != nil {
		log.Fatalf("Failed to connect to Cassandra at port %d: %v", config.Port, err)
	}
	return session
}

func InsertRowWithLWT(session *gocql.Session, conversationID string) (bool, error) {
	query := `INSERT INTO payment_id_by_conversation_id (conversation_id, payment_id) VALUES (?, ?) IF NOT EXISTS`
	var applied bool
	var existingConvID string
	var existingPaymentID gocql.UUID
	applied, err := session.Query(query, conversationID, gocql.TimeUUID()).ScanCAS(&existingConvID, &existingPaymentID)
	return applied, err
}

func CloseSession(session *gocql.Session) {
	session.Close()
}
