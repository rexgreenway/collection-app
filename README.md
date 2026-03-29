# Run Server Locally

```bash
ENVIRONMENT=dev go run cmd/main.go
```

# Equivalent gRPC and cURL commands:

```bash
grpcurl -plaintext -d '{"id": "hello"}' localhost:50100 collection.CollectionService/GetCollection
```

```bash
curl -X GET localhost:8089/v1/collections/hello
```

# GRPC Gateway README:

https://github.com/grpc-ecosystem/grpc-gateway#readme
