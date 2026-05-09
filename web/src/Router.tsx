import { createHashRouter, Navigate, RouterProvider } from "react-router";

import App from "./layout/App";

import CollectionsPage, {
  collectionsLoader,
  COLLECTIONS_PATH,
} from "./collections";
import ItemsPage, { itemsLoader, ITEMS_PATH } from "./items";

const router = createHashRouter([
  {
    path: "/",
    element: <App />,
    errorElement: <h1>Something Went Wrong</h1>,
    children: [
      { path: "test", element: <h1>TESTING</h1> },

      { index: true, element: <Navigate to={COLLECTIONS_PATH} replace /> },

      {
        path: COLLECTIONS_PATH,
        element: <CollectionsPage />,
        loader: collectionsLoader,
        children: [],
      },

      {
        path: `${COLLECTIONS_PATH}/:id`,
        children: [
          { index: true, element: <Navigate to={ITEMS_PATH} replace /> },
          {
            path: ITEMS_PATH,
            element: <ItemsPage />,
            loader: itemsLoader,
            children: [],
          },
        ],
      },
    ],
  },
]);

export default function Router() {
  return <RouterProvider router={router} />;
}
