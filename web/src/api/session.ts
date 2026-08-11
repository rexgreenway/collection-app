import type {
  CollectionResponse,
  Pagination,
  CreateCollectionRequest,
  UpdateCollectionRequest,
  ItemResponse,
  CreateItemRequest,
  UpdateItemRequest,
} from "./types";

import type CollectionClient from "./interface";

const COLLECTIONS_KEY = "collections";
const ITEMS_PREFIX = "items:";

function loadCollectionMap(): Record<string, CollectionResponse> {
  const raw = sessionStorage.getItem(COLLECTIONS_KEY);
  return raw ? JSON.parse(raw) : {};
}

function saveCollectionMap(
  collections: Record<string, CollectionResponse>,
): void {
  sessionStorage.setItem(COLLECTIONS_KEY, JSON.stringify(collections));
}

function loadCollectionItemsMap(
  collectionId: string,
): Record<string, ItemResponse> {
  const raw = sessionStorage.getItem(ITEMS_PREFIX + collectionId);
  return raw ? JSON.parse(raw) : {};
}

function saveCollectionItems(
  collectionId: string,
  items: Record<string, ItemResponse>,
): void {
  sessionStorage.setItem(ITEMS_PREFIX + collectionId, JSON.stringify(items));
}

function removeCollectionItems(collectionId: string): void {
  sessionStorage.removeItem(ITEMS_PREFIX + collectionId);
}

// --- Pagination helper ---

function paginate<T>(
  items: T[],
  pagination?: Pagination,
): { data: T[]; pagination: Pagination } {
  const page = pagination?.page ?? 1;
  const pageSize = pagination?.pageSize ?? 10;
  const start = (page - 1) * pageSize;
  return {
    data: items.slice(start, start + pageSize),
    pagination: { page, pageSize },
  };
}

export const SessionStorageClient: CollectionClient = {
  // ---------- COLLECTIONS -----------

  async listCollections(pagination?: Pagination) {
    const map = loadCollectionMap();
    return paginate(Object.values(map), pagination);
  },

  async createCollection(req: CreateCollectionRequest) {
    const collectionMap = loadCollectionMap();

    const id = `DEMO-${Object.keys(collectionMap).length + 1}`;

    const collection: CollectionResponse = {
      id: id,
      name: req.name,
      itemCount: 0,
    };

    collectionMap[id] = collection;
    saveCollectionMap(collectionMap);
    saveCollectionItems(id, {}); // initialize empty item store
    return { data: collection };
  },

  async getCollection(id: string) {
    const map = loadCollectionMap();
    const found = map[id];
    if (!found) throw new Error(`Collection '${id}' not found`);

    const items = loadCollectionItemsMap(id);
    return {
      data: { ...found, itemCount: Object.keys(items).length },
    };
  },

  async updateCollection(id: string, req: UpdateCollectionRequest) {
    const map = loadCollectionMap();
    if (!map[id]) throw new Error(`Collection '${id}' not found`);

    map[id] = { ...map[id], ...req, id };
    saveCollectionMap(map);
    return { data: map[id] };
  },

  async deleteCollection(id: string) {
    const map = loadCollectionMap();
    delete map[id];
    saveCollectionMap(map);
    removeCollectionItems(id); // clean up associated items
  },

  // ---------- ITEMS -----------

  async listItems(collectionId: string, pagination?: Pagination) {
    const items = loadCollectionItemsMap(collectionId);
    return paginate(Object.values(items), pagination);
  },

  async createItem(collectionId: string, req: CreateItemRequest) {
    const items = loadCollectionItemsMap(collectionId);
    const id = `${collectionId}-${Object.keys(items).length + 1}`;

    const item: ItemResponse = { id, name: req.name };
    items[id] = item;
    saveCollectionItems(collectionId, items);
    return { data: item };
  },

  async createItems(collectionId: string, reqs: CreateItemRequest[]) {
    const items = loadCollectionItemsMap(collectionId);
    const createdItems: ItemResponse[] = [];

    reqs.forEach((req, i) => {
      const id = `${collectionId}-${Object.keys(items).length + i + 1}`;
      const item: ItemResponse = { id, name: req.name };
      items[id] = item;
      createdItems.push(item);
    });

    saveCollectionItems(collectionId, items);
    return paginate(createdItems, { page: 1, pageSize: createdItems.length });
  },

  async getItem(collectionId: string, id: string) {
    const items = loadCollectionItemsMap(collectionId);
    const found = items[id];
    if (!found) throw new Error(`Item '${id}' not found`);
    return { data: found };
  },

  async updateItem(collectionId: string, id: string, req: UpdateItemRequest) {
    const items = loadCollectionItemsMap(collectionId);
    if (!items[id]) throw new Error(`Item '${id}' not found`);
    items[id] = { ...items[id], ...req };
    saveCollectionItems(collectionId, items);
    return { data: items[id] };
  },

  async deleteItem(collectionId: string, id: string) {
    const items = loadCollectionItemsMap(collectionId);
    delete items[id];
    saveCollectionItems(collectionId, items);
  },
};
