import { fromJson } from "@bufbuild/protobuf";

import {
  GetCollectionResponseSchema,
  ListCollectionsResponseSchema,
} from "@collection-app/gen/v1";

import type CollectionClient from "./interface";

import { type Collection } from "./types";
import {
  protoToCollection,
  transformListResponse,
  transformResponse,
} from "./transformers";

export const RestClient: CollectionClient = {
  async listCollections() {
    const res = await fetch("/v1/collections");
    if (!res.ok) {
      throw new Error(`Failed to list Collections: ${res.status}`);
    }
    const proto = fromJson(ListCollectionsResponseSchema, await res.json());
    return transformListResponse(proto, protoToCollection);
  },

  async createCollection(collection: Collection) {
    const res = await fetch("/v1/collections", {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify(collection),
    });
    if (!res.ok) {
      throw new Error(`Failed to create Collection: ${res.status}`);
    }
    const proto = fromJson(GetCollectionResponseSchema, await res.json());
    return transformResponse(proto, protoToCollection);
  },

  async getCollection(id: string) {
    const res = await fetch(`/v1/collections/${id}`);
    if (!res.ok) {
      throw new Error(`Failed to get Collection '${id}': ${res.status}`);
    }
    const proto = fromJson(GetCollectionResponseSchema, await res.json());
    return transformResponse(proto, protoToCollection);
  },

  async updateCollection(id: string, collection: Collection) {
    const res = await fetch(`/v1/collections/${id}`, {
      method: "PATCH",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify(collection),
    });
    if (!res.ok) {
      throw new Error(`Failed to update Collection '${id}': ${res.status}`);
    }
    const proto = fromJson(GetCollectionResponseSchema, await res.json());
    return transformResponse(proto, protoToCollection);
  },

  async deleteCollection(id: string) {
    const res = await fetch(`/v1/collections/${id}`, { method: "DELETE" });
    if (!res.ok) {
      throw new Error(`Failed to delete Collection '${id}': ${res.status}`);
    }
  },
};
