import { useEffect, useState } from "react";
import { RestClient } from "./api/fetch";

import type { Collection } from "@collection-app/gen/v1";

// import collectionAppLogo from "./assets/collection-app-v1.svg";

function App() {
  const [collections, setCollections] = useState<Collection[]>([]);

  const listEm = () => {
    RestClient.listCollections()
      .then((response) => {
        setCollections(response.data);
      })
      .catch((error) => {
        console.log(error);
      });
  };

  useEffect(listEm, []);

  return (
    <main>
      {/* <img src={collectionAppLogo} alt="Collection App" /> */}
      <h1>Collections</h1>
      <button onClick={listEm}>Get Em.</button>
      {collections.map((c) => (
        <p key={c.id}>{c.name}</p>
      ))}
    </main>
  );
}

export default App;
