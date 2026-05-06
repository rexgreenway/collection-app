import { useEffect } from "react";
import DarkModeIcon from "@mui/icons-material/DarkMode";
import LightModeIcon from "@mui/icons-material/LightMode";

import { useTheme } from "./contexts";

import { useCollectionStore } from "./store/collection";

import Header from "./components/Header";
import CollectionBubbleChart from "./components/d3/CollectionBubble";

import CreateCollection from "./collections/Create";

import styles from "./App.module.css";

const App = () => {
  // THEME CONTEXT
  const { theme, toggleTheme } = useTheme();
  const title =
    theme === "dark" ? "Switch to Light mode" : "Switch to Dark mode";

  const collections = useCollectionStore((s) => s.collections);
  const fetchCollections = useCollectionStore((s) => s.fetchCollections);

  // Render collections on page
  useEffect(() => {
    fetchCollections();
  }, [fetchCollections]);

  return (
    <>
      <div id="app" className={styles.App}>
        <Header />

        <footer className={styles.Footer}>
          <div onClick={toggleTheme} title={title}>
            {theme === "dark" ? <LightModeIcon /> : <DarkModeIcon />}
          </div>

          {/* This is just the Button... This feels wrong to have this here */}
          <CreateCollection />
        </footer>
      </div>

      {/* This sits outside normal document flow as this 
      is the background Collection View */}
      <CollectionBubbleChart>
        {...collections.map((c) => {
          // Radius of collections should be the size of the collection??
          return { group: c.name, radius: 2 };
        })}
      </CollectionBubbleChart>

      {/* This sits outside normal document flow as this 
      is where modals are mounted */}
      <div id="modal-root" />
    </>
  );
};

export default App;
