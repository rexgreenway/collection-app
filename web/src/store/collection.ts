import { create } from "zustand";

import type { Collection, CreateCollectionRequest } from "../api/types";
import { ApiClientFactory, ApiClientImpl } from "../api";

const api = ApiClientFactory(ApiClientImpl.SESSION_STORAGE);

type CollectionsState = {
  collections: Collection[];
  fetchCollections: () => Promise<void>;
  createCollection: (req: CreateCollectionRequest) => Promise<void>;
  deleteCollection: (id: string) => Promise<void>;
};

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
}));
