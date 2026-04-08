import { RestClient } from "./fetch";
import { LocalStorageClient } from "./local";

export const ApiClientImpl = {
  LOCAL_STORAGE: "local",
  FETCH: "fetch",
} as const;

type ApiClientImpl = (typeof ApiClientImpl)[keyof typeof ApiClientImpl];

export const ApiClientFactory = (impl: ApiClientImpl) => {
  switch (impl) {
    case "local":
      return LocalStorageClient;
    case "fetch":
      return RestClient;
    default:
      throw new Error(`Unsupported API client implementation: ${impl}`);
  }
};
