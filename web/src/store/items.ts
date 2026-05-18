import { create } from "zustand";

import type { ItemResponse, CreateItemRequest } from "../api/types";
import { ApiClientFactory, ApiClientImpl } from "../api";

const api = ApiClientFactory(ApiClientImpl.SESSION_STORAGE);

interface ItemsState {
  items: ItemResponse[];
  fetchItems: (collectionId: string) => Promise<void>;
  createItem: (collectionId: string, req: CreateItemRequest) => Promise<void>;
  deleteItem: (collectionId: string, id: string) => Promise<void>;
}

export const useItemStore = create<ItemsState>((set) => ({
  items: [],

  fetchItems: async (collectionId) => {
    const res = await api.listItems(collectionId);
    set({ items: res.data });
  },

  createItem: async (collectionId, req) => {
    const res = await api.createItem(collectionId, req);
    set((state) => ({ items: [...state.items, res.data] }));
  },

  deleteItem: async (collectionId, id) => {
    await api.deleteItem(collectionId, id);
    set((state) => ({
      items: state.items.filter((item) => item.id !== id),
    }));
  },
}));
