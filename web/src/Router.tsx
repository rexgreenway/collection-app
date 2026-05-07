import { createHashRouter, Navigate, RouterProvider } from "react-router";

import App from "./App";

import CreateCollectionModal from "./collections/CreateModal";
import CollectionsPage, { collectionsLoader } from "./collections/Page";
import ItemsPage, { itemsLoader } from "./items/Page";

import SettingsModal from "./settings/SettingsModal";

const router = createHashRouter([
  {
    path: "/",
    element: <App />,
    errorElement: <h1>Something Went Wrong</h1>,
    children: [
      { index: true, element: <Navigate to="collections" replace /> },

      {
        path: "collections",
        element: <CollectionsPage />,
        loader: collectionsLoader,
        children: [
          {
            path: "settings",
            element: <SettingsModal context="collections" />,
          },

          // Collection Actions
          { path: "create", element: <CreateCollectionModal /> },
        ],
      },

      // SHOULD ADD LOADERS TO CERTAIN ONES HERE -> ITEMS
      {
        path: "collections/:id",
        children: [
          { index: true, element: <Navigate to="items" replace /> },
          {
            path: "items",
            element: <ItemsPage />,
            loader: itemsLoader,
            children: [
              { path: "settings", element: <SettingsModal context="items" /> },

              // Items Actions
              // { path: "create", element: <CreateItemModal /> },
            ],
          },
        ],
      },
    ],
  },
]);

export default function Router() {
  return <RouterProvider router={router} />;
}
