# Equivalent gRPC and cURL commands:

```bash
grpcurl -plaintext -d '{"id": "hello"}' localhost:50100 collection.CollectionService/GetCollection
```

```bash
curl -X GET localhost:8089/v1/collections/hello
```
