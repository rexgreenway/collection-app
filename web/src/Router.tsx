import { createHashRouter, RouterProvider } from "react-router";

import App from "./App";

import SettingsModal from "./settings/Settings";
import CreateCollectionModal from "./collections/Create";

const router = createHashRouter([
  {
    path: "/",
    element: <App />,
    errorElement: <h1>Something Went Wrong</h1>,
    children: [
      { path: "settings", element: <SettingsModal /> },
      { path: "create", element: <CreateCollectionModal /> },
    ],
  },
]);

export default function Router() {
  return <RouterProvider router={router} />;
}
