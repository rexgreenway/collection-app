import type {
  Collection,
  GetCollectionResponse,
  ListCollectionsResponse,
} from "@collection-app/gen/v1";

import type CollectionClient from "./interface";

export const RestClient: CollectionClient = {
  async listCollections() {
    const res = await fetch("/collections");
    return await res.json();
  },
  async createCollection(collection: Collection) {
    const res = await fetch("/collections", {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify(collection),
    });
    return await res.json();
  },
  async getCollection(id: string) {
    const res = await fetch(`/collections/${id}`);
    return await res.json();
  },
  async updateCollection(id: string, collection: Collection) {
    const res = await fetch(`/collections/${id}`, {
      method: "PATCH",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify(collection),
    });
    return await res.json();
  },
  async deleteCollection(id: string) {
    await fetch(`/collections/${id}`, { method: "DELETE" });
  },
};
