# Cassandra Concurrency Control POC

This project demonstrates concurrency control in a Cassandra cluster using Lightweight Transactions (LWT). The setup uses Docker Compose to create a local multi-node, multi-datacenter Cassandra cluster.

## Prerequisites

- Docker
- Docker Compose

## Get Started

### 1. Clone the Repository:

```bash
git clone [repository-url]
cd [repository-directory]
```

### 2. Start the Cassandra Cluster:

Navigate to the directory containing the `docker-compose.yml` file and run:

```bash
docker-compose up -d
```

This will start a 2-datacenter Cassandra cluster with 4 nodes in each datacenter.

### 3. Verify Cluster Setup with Nodetool:

To check the status of the nodes in the cluster, you can use the `nodetool` command:

```bash
docker exec -it cassandra-dc1-node1 nodetool status
```

```commandline
Datacenter: dc1
===============
Status=Up/Down
|/ State=Normal/Leaving/Joining/Moving
--  Address     Load        Tokens  Owns (effective)  Host ID                               Rack 
UN  172.21.0.4  127.67 KiB  16      100.0%            9ce73c5c-b589-41df-93d6-c05c9180f96a  rack1
UN  172.21.0.2  150.2 KiB   16      100.0%            d3454bb8-7b35-45a6-b35f-fcccded6cfe8  rack1
UN  172.21.0.3  121.28 KiB  16      100.0%            c261eefd-3093-47a2-8eaa-601311a4260a  rack1

Datacenter: dc2
===============
Status=Up/Down
|/ State=Normal/Leaving/Joining/Moving
--  Address     Load        Tokens  Owns (effective)  Host ID                               Rack 
UN  172.21.0.5  121.24 KiB  16      100.0%            06b81e50-ab5c-4165-a900-718c39626f0b  rack1
UN  172.21.0.7  137.27 KiB  16      100.0%            99a24c67-52cb-41f2-ab57-d94ef5e2275a  rack1
UN  172.21.0.6  121.21 KiB  16      100.0%            6b388013-8060-471b-94d4-8c67c5f49f89  rack1

```

### 4. Create the Tables:

You can execute the create_tables.cql script to create the necessary tables:

**Create tables**
```bash
cqlsh -f scripts/create_tables.cql
````

**Check tables:**
```bash
cqlsh -e "DESCRIBE KEYSPACE payment"
```

**Reset tables:**
```bash
cqlsh -f scripts/reset_tables.cql
```

### 7. Testing Concurrency:

To test concurrency control, you can execute the batch.cql script:

```bash
go run insert.go
```
```bash
cqlsh -e "USE payment; SELECT * from payment_id_by_conversation_id;"
```
```bash
docker exec -it cassandra-dc1-node1 nodetool setlogginglevel ROOT TRACE
```
You should get error.

## Notes

- This setup is intended for local testing and development purposes only.
- Adjust the replication factor in `create_tables.cql` based on your requirements.

## Contributing

If you'd like to contribute, please fork the repository and use a feature branch. Pull requests are warmly welcome.

## Links

- [Cassandra Docker Image Documentation](https://hub.docker.com/_/cassandra)
- [Cassandra Lightweight Transactions](https://docs.datastax.com/en/cql-oss/3.3/cql/cql_using/useInsertLWT.html)
