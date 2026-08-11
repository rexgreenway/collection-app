# Collection App

A self-hosted application for curating and tracking collections.

# Development

## Protos

The Application uses Protocol Buffers to define Types and API Contracts.

Proto files can be found in the [proto directory](./proto/).

### Generation

This repository uses `buf` to manage dependencies and generate code from proto
files.

- Generate code from proto files:

```bash
buf generate
```

## Server

- Running Locally:

```bash
ENVIRONMENT=dev go run cmd/main.go
```

## Frontend

- Running Locally:

```bash
cd dev
pnpm dev
```

# Equivalent gRPC and cURL commands:

```bash
grpcurl -plaintext -d '{"id": "hello"}' localhost:50100 collection.CollectionService/GetCollection
```

```bash
curl -X GET localhost:8089/v1/collections
curl -X GET localhost:8089/v1/collections/ID_GOES_HERE
curl -X POST -d '{ "name": "new_collection" }' localhost:8089/v1/collections
```

## Helpful Docs:

## GRPC Gateway README:

https://github.com/grpc-ecosystem/grpc-gateway#readme

## Google Annotations Info:

https://github.com/googleapis/googleapis/blob/master/google/api/README.md

## Google gRPC Transcoding Docs

https://docs.cloud.google.com/endpoints/docs/grpc/transcoding
https://github.com/googleapis/googleapis/blob/7ae842846c3fa71ea909e6ad04d3d0b9b06756e9/google/api/http.proto#L43

## Tutorial:

https://www.speakeasy.com/openapi/frameworks/grpc-gateway
https://github.com/speakeasy-api/speakeasy-grpc-gateway-example
