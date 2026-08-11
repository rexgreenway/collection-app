import { useCollectionStore } from "../store/collection";

// Fetches Data from the API ahead of render
const collectionsLoader = async () => {
  await useCollectionStore.getState().fetchCollections();
  return null;
};

export default collectionsLoader;
