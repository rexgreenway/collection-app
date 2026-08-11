import { useItemStore } from "../store/items";
import type { LoaderFunctionArgs } from "react-router";

const itemsLoader = async ({ params }: LoaderFunctionArgs) => {
  const collectionId = params.id!;
  await useItemStore.getState().fetchItems(collectionId);
  return null;
};

export default itemsLoader;
