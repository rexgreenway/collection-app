import type {
  Collection,
  GetCollectionResponse,
  ListCollectionsResponse,
} from "@collection-app/gen/v1";

export default interface CollectionClient {
  listCollections(): Promise<ListCollectionsResponse>;
  createCollection(collection: Collection): Promise<GetCollectionResponse>;
  getCollection(id: string): Promise<GetCollectionResponse>;
  updateCollection(
    id: string,
    collection: Collection,
  ): Promise<GetCollectionResponse>;
  deleteCollection(id: string): Promise<void>;
}
