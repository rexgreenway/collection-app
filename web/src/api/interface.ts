import type {
  Collection,
  ListResponse,
  GetResponse,
  Pagination,
  CreateCollectionRequest,
  UpdateCollectionRequest,
} from "./types";

export default interface CollectionClient {
  listCollections(pagination?: Pagination): Promise<ListResponse<Collection>>;
  createCollection(
    req: CreateCollectionRequest,
  ): Promise<GetResponse<Collection>>;
  getCollection(id: string): Promise<GetResponse<Collection>>;
  updateCollection(
    id: string,
    req: UpdateCollectionRequest,
  ): Promise<GetResponse<Collection>>;
  deleteCollection(id: string): Promise<void>;
}
