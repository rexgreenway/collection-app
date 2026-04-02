import type {
  Collection,
  GetCollectionResponse,
  ListCollectionsResponse,
} from "@collection-app/gen/v1";

import type CollectionClient from "./interface";

export const RestClient: CollectionClient = {
  async listCollections() {
    const res = await fetch("/v1/collections");
    if (!res.ok) {
      throw new Error(`Failed to list collections: ${res.status}`);
    }
    return (await res.json()) as ListCollectionsResponse;
  },
  async createCollection(collection: Collection) {
    const res = await fetch("/v1/collections", {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify(collection),
    });
    if (!res.ok) {
      throw new Error(`Failed to create collection: ${res.status}`);
    }
    return (await res.json()) as Collection;
  },
  async getCollection(id: string) {
    const res = await fetch(`/v1/collections/${id}`);
    if (!res.ok) {
      throw new Error(`Failed to get collections '${id}': ${res.status}`);
    }
    return (await res.json()) as GetCollectionResponse;
  },
  async updateCollection(id: string, collection: Collection) {
    const res = await fetch(`/v1/collections/${id}`, {
      method: "PATCH",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify(collection),
    });
    if (!res.ok) {
      throw new Error(`Failed to update collection '${id}': ${res.status}`);
    }
    return (await res.json()) as Collection;
  },
  async deleteCollection(id: string) {
    await fetch(`/v1/collections/${id}`, { method: "DELETE" });
  },
};
