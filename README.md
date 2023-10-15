# Cassandra Concurrency Control POC

This project demonstrates concurrency control in a Cassandra cluster using batch statements with Lightweight Transactions (LWT). The setup uses Docker Compose to create a local multi-node, multi-datacenter Cassandra cluster.

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

### 4. Accessing Cassandra:

To access the Cassandra node, use the `cqlsh` command:

```bash
docker exec -it cassandra-dc1-node1 cqlsh
```

### 6. Create the Tables:

You can execute the setup.cql script to create the necessary tables:

```bash
cqlsh -f scripts/setup.cql
````

### 7. Testing Concurrency:

To test concurrency control, you can execute the batch.cql script:

```bash
cqlsh -f scripts/batch.cql
```
```commandline
Error from server: code=2200 [Invalid query] message="Batch with conditions cannot span multiple tables"
```

You should get error.

## Notes

- This setup is intended for local testing and development purposes only.
- Adjust the replication factor in `setup.cql` based on your requirements.

## Contributing

If you'd like to contribute, please fork the repository and use a feature branch. Pull requests are warmly welcome.

## Links

- [Cassandra Docker Image Documentation](https://hub.docker.com/_/cassandra)
- [Cassandra Lightweight Transactions](https://docs.datastax.com/en/cql-oss/3.3/cql/cql_using/useInsertLWT.html)
