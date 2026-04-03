import { useTheme } from "./contexts";

import styles from "./App.module.css";
import { useEffect, useState } from "react";
import { RestClient } from "./api/fetch";
import type { Collection } from "@collection-app/gen/v1";

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

  return (
    <div id="app" className={styles.App}>
      {/* Header */}
      <header>
        <h1>Collection App</h1>
        <button onClick={toggleTheme} title={title}>
          BUTTON
        </button>
      </header>

      {/* List Collections */}
      <section>
        <div>
          <h2>Collections</h2>
          <button onClick={listEm}>Get Em.</button>
        </div>
      </section>
      {/* THe list of collections itself */}
      <section>
        {collections.map((c) => (
          <p key={c.id}>{c.name}</p>
        ))}
      </section>

      <footer>
        <h3>FOOTER - links - etc...</h3>
      </footer>
    </div>
  );
};

export default App;
