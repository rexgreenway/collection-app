import type { Collection, Pagination } from "./types";

import type CollectionClient from "./interface";

const STORAGE_KEY = "collections";

function loadCollections(): Collection[] {
  const raw = localStorage.getItem(STORAGE_KEY);
  return raw ? JSON.parse(raw) : [];
}

function saveCollections(collections: Collection[]): void {
  localStorage.setItem(STORAGE_KEY, JSON.stringify(collections));
}

export const LocalStorageClient: CollectionClient = {
  async listCollections(pagination?: Pagination) {
    const collections = loadCollections();

    // Fake pagination
    const page = pagination?.page ?? 1;
    const pageSize = pagination?.pageSize ?? 10;
    const start = (page - 1) * pageSize;
    const end = start + pageSize;

    return {
      data: collections.slice(start, end),
      pagination: {
        page: page,
        pageSize: pageSize,
      },
    };
  },

  async createCollection(collection: Collection) {
    const collections = loadCollections();
    collection.id = crypto.randomUUID();
    collections.push(collection);
    saveCollections(collections);
    return { data: collection };
  },

  async getCollection(id: string) {
    const collections = loadCollections();
    const found = collections.find((c) => c.id === id);
    if (!found) {
      throw new Error(`Collection '${id}' not found`);
    }
    return { data: found };
  },

  async updateCollection(id: string, collection: Collection) {
    const collections = loadCollections();
    const idx = collections.findIndex((c) => c.id === id);
    if (idx === -1) {
      throw new Error(`Collection '${id}' not found`);
    }
    collections[idx] = { ...collections[idx], ...collection, id };
    saveCollections(collections);
    return { data: collections[idx] };
  },

  async deleteCollection(id: string) {
    const collections = loadCollections();
    saveCollections(collections.filter((c) => c.id !== id));
  },
};
