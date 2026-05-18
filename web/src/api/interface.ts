import type {
  ListResponse,
  GetResponse,
  Pagination,
  CreateCollectionRequest,
  UpdateCollectionRequest,
  CollectionResponse,
  ItemResponse,
  CreateItemRequest,
  UpdateItemRequest,
} from "./types";

export default interface CollectionClient {
  // -> COLLECTIONS
  listCollections(
    pagination?: Pagination,
  ): Promise<ListResponse<CollectionResponse>>;

  createCollection(
    req: CreateCollectionRequest,
  ): Promise<GetResponse<CollectionResponse>>;

  getCollection(id: string): Promise<GetResponse<CollectionResponse>>;

  updateCollection(
    id: string,
    req: UpdateCollectionRequest,
  ): Promise<GetResponse<CollectionResponse>>;

  deleteCollection(id: string): Promise<void>;

  // -> ITEMS
  listItems(
    collectionId: string,
    pagination?: Pagination,
  ): Promise<ListResponse<ItemResponse>>;

  createItem(
    collectionId: string,
    req: CreateItemRequest,
  ): Promise<GetResponse<ItemResponse>>;

  createItems(
    collectionId: string,
    reqs: CreateItemRequest[],
  ): Promise<ListResponse<ItemResponse>>;

  getItem(collectionId: string, id: string): Promise<GetResponse<ItemResponse>>;

  updateItem(
    collectionId: string,
    id: string,
    req: UpdateItemRequest,
  ): Promise<GetResponse<ItemResponse>>;

  deleteItem(collectionId: string, id: string): Promise<void>;
}
