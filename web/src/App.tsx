import { useEffect, useState } from "react";

import { type Collection } from "@collection-app/gen/v1";

import { useTheme } from "./contexts";
import { RestClient } from "./api/fetch";

// import Bubble from "./components/Bubble";
import BubbleCanvas from "./components/BubbleCanvas";

import styles from "./App.module.css";

const App = () => {
  // THEME CONTEXT
  const { theme, toggleTheme } = useTheme();
  const title =
    theme === "dark" ? "Switch to Light mode" : "Switch to Dark mode";

  // List Collections
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

  const generateRandomArray = () => {
    const length = Math.floor(Math.random() * 10) + 1;
    return Array.from({ length }, () => Math.ceil(Math.random() * 10 + 4));
  };

  const [foo, setFoo] = useState(generateRandomArray());

  return (
    <div id="app" className={styles.App}>
      {/* Header */}
      <header>
        <h1>Collection App</h1>
        <button onClick={toggleTheme} title={title}>
          BUTTON
        </button>
      </header>

      <div>{foo}</div>

      <BubbleCanvas>
        {/* {...collections.map((c) => {
          // Radius of collections should be the size of the collection??
          return { group: c.name, radius: 1 };
        })} */}
        {...foo.map((c) => {
          // Radius of collections should be the size of the collection??
          return { group: `test-${c}`, radius: c };
        })}
      </BubbleCanvas>

      <button onClick={() => setFoo(generateRandomArray())}>
        Add data point
      </button>

      {/* List Collections */}
      {/* <section>
        <div>
          <h2>Collections</h2>
          <button onClick={listEm}>Get Em.</button>
        </div>
      </section> */}
      {/* THe list of collections itself */}
      {/* <section>
        {collections.map((c) => (
          <p key={c.id}>{c.name}</p>
        ))}
      </section> */}

      <footer>
        <h3>FOOTER - links - etc...</h3>
      </footer>
    </div>
  );
};

export default App;
