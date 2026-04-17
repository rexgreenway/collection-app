import { RestClient } from "./fetch";
import { SessionStorageClient } from "./session";

export const ApiClientImpl = {
  SESSION_STORAGE: "session",
  FETCH: "fetch",
} as const;

type ApiClientImpl = (typeof ApiClientImpl)[keyof typeof ApiClientImpl];

export const ApiClientFactory = (impl: ApiClientImpl) => {
  switch (impl) {
    case "session":
      return SessionStorageClient;
    case "fetch":
      return RestClient;
    default:
      throw new Error(`Unsupported API client implementation: ${impl}`);
  }
};
