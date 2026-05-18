import type { Pagination } from "./common";

export type GetResponse<T> = {
  data: T;
};

export type ListResponse<T> = {
  data: T[];
  pagination: Pagination;
};

// -- Collections

export type CollectionResponse = {
  id: string;
  name: string;
  itemCount: number;
};

// -- Items

export type ItemResponse = {
  id: string;
  name: string;
};
