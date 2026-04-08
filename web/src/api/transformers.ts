import {
  type PaginationParams,
  type Collection as pbCollection,
} from "@collection-app/gen/v1";

import type { Collection, ListResponse, Pagination } from "./types";

export const transformResponse = <T, V>(
  response: { data?: T },
  transformFunc: (item: T) => V,
): { data: V } => {
  if (!response.data) {
    throw new Error("Server returned empty response data");
  }
  return { data: transformFunc(response.data) };
};

export const transformPagination = (
  pagination?: PaginationParams,
): Pagination => {
  return {
    page: pagination?.page ?? 1,
    pageSize: pagination?.pageSize ?? 0,
  };
};

export const transformListResponse = <T, V>(
  listResponse: { data: T[]; pagination?: PaginationParams },
  transformFunc: (item: T) => V,
): ListResponse<V> => {
  return {
    data: listResponse.data.map(transformFunc),
    pagination: transformPagination(listResponse.pagination),
  };
};

// Convert a protobuf message to your plain app type
export const protoToCollection = (proto: pbCollection): Collection => {
  return {
    id: proto.id,
    name: proto.name,
  };
};
