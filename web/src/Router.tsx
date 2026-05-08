import { createHashRouter, Navigate, RouterProvider } from "react-router";

import App from "./layout/AppLayout";

import CollectionsPage, {
  CreateModal,
  collectionsLoader,
  COLLECTIONS_CONTEXT,
} from "./collections";
import ItemsPage, { itemsLoader, ITEMS_CONTEXT } from "./items";

import SettingsModal from "./settings/SettingsModal";

const router = createHashRouter([
  {
    path: "/",
    element: <App />,
    errorElement: <h1>Something Went Wrong</h1>,
    children: [
      { path: "test", element: <h1>TESTING</h1> },

      { index: true, element: <Navigate to="collections" replace /> },

      {
        path: "collections",
        element: <CollectionsPage />,
        loader: collectionsLoader,
        children: [
          {
            path: "settings",
            element: <SettingsModal context={COLLECTIONS_CONTEXT} />,
          },

          // Collection Actions
          { path: "create", element: <CreateModal /> },
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
              {
                path: "settings",
                element: <SettingsModal context={ITEMS_CONTEXT} />,
              },

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
