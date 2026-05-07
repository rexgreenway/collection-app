import { useEffect } from "react";

import { useCollectionStore } from "./store/collection";

import Header from "./components/Header";
import CollectionBubbleChart from "./components/d3/CollectionBubble";

import CreateCollection from "./collections/Create";

import styles from "./App.module.css";

const App = () => {
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
