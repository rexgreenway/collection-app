import { create } from "zustand";

import type { CollectionResponse, CreateCollectionRequest } from "../api/types";
import { ApiClientFactory, ApiClientImpl } from "../api";

const api = ApiClientFactory(ApiClientImpl.SESSION_STORAGE);

interface CollectionsState {
  collections: CollectionResponse[];
  fetchCollections: () => Promise<void>;
  createCollection: (req: CreateCollectionRequest) => Promise<void>;
  deleteCollection: (id: string) => Promise<void>;
  getCollection: (id: string) => Promise<CollectionResponse>;
}

export const useCollectionStore = create<CollectionsState>((set) => ({
  collections: [],

  fetchCollections: async () => {
    const res = await api.listCollections();
    set({ collections: res.data });
  },

  createCollection: async (req) => {
    const res = await api.createCollection(req);
    set((state) => ({ collections: [...state.collections, res.data] }));
  },

  deleteCollection: async (id) => {
    await api.deleteCollection(id);
    set((state) => ({
      collections: state.collections.filter((c) => c.id !== id),
    }));
  },

  getCollection: async (id) => {
    const res = await api.getCollection(id);
    return res.data;
  },
}));
