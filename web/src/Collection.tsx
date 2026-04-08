import { useEffect, useState } from "react";

import { ApiClientFactory, ApiClientImpl } from "./api";
import BubbleCanvas from "./components/d3/BubbleCanvas";

import type { Collection } from "./api/types";

const CollectionsSection = () => {
  // get LOCAL STORAGE
  const api = ApiClientFactory(ApiClientImpl.LOCAL_STORAGE);

  // List Collections
  const [collections, setCollections] = useState<Collection[]>([]);

  const generateRandomName = () => {
    const adjectives = ["quick", "lazy", "happy", "bright", "wild"];
    const nouns = ["fox", "bear", "eagle", "tiger", "wolf"];
    const adj = adjectives[Math.floor(Math.random() * adjectives.length)];
    const noun = nouns[Math.floor(Math.random() * nouns.length)];
    return `${adj}-${noun}`;
  };

  const createOne = () => {
    api
      .createCollection({ name: generateRandomName() })
      .then((resp) => setCollections((prev) => [...prev, resp.data]))
      .catch((error) => {
        console.log(error);
      });
  };

  const listEm = () => {
    api
      .listCollections()
      .then((response) => {
        setCollections(response.data);
      })
      .catch((error) => {
        console.log(error);
      });
  };

  useEffect(listEm, [api]);

  return (
    <>
      <section id="collection">
        <div>
          <h2>Collections</h2>
          <button onClick={createOne}>Create One.</button>
          <button onClick={listEm}>Get Em.</button>
        </div>
        <div>
          {collections.map((c) => (
            <p key={c.id}>
              {c.id} - {c.name}
            </p>
          ))}
        </div>
      </section>

      <BubbleCanvas>
        {...collections.map((c) => {
          // Radius of collections should be the size of the collection??
          return { group: c.name, radius: 10 };
        })}
      </BubbleCanvas>
    </>
  );
};

export default CollectionsSection;
