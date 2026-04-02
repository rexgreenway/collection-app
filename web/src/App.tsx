import { useEffect, useState } from "react";
import { RestClient } from "./api/fetch";
import type { Collection } from "@collection-app/gen/v1";

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
      <h1>Collections</h1>
      <button onClick={listEm}>Get Em.</button>
      {collections.map(c => <p key={c.id}>{c.name}</p>)}
    </main>
  );
}

export default App;
