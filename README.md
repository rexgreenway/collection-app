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

# Helpful Docs:

## GRPC Gateway README:

https://github.com/grpc-ecosystem/grpc-gateway#readme

## Tutorial:

https://www.speakeasy.com/openapi/frameworks/grpc-gateway
https://github.com/speakeasy-api/speakeasy-grpc-gateway-example
