import { useEffect, useMemo } from "react";
import { Outlet, useNavigate } from "react-router";
import { Add } from "@mui/icons-material";

import { useCollectionStore } from "./store/collection";

import Header from "./components/Header";
import CircleButton from "./components/ui/CircleButton";
import CollectionBubbleChart from "./components/d3/CollectionBubble";

import styles from "./App.module.css";

const App = () => {
  const navigate = useNavigate();

  const collections = useCollectionStore((s) => s.collections);
  const fetchCollections = useCollectionStore((s) => s.fetchCollections);

  // Render collections on page
  useEffect(() => {
    fetchCollections();
  }, [fetchCollections]);

  const bubbleData = useMemo(
    () => collections.map((c) => ({ group: c.name, radius: 2 })),
    [collections],
  );

  return (
    <>
      <div id="app" className={styles.App}>
        <Header />

        <footer className={styles.Footer}>
          <CircleButton onClick={() => navigate("/create")} icon={Add} />
        </footer>
      </div>

      {/* This sits outside normal document flow as this 
      is the background Collection View */}
      {/* This could possibly sit somewhere else with react router now introduced */}
      <CollectionBubbleChart>{bubbleData}</CollectionBubbleChart>

      {/* This sits outside normal document flow as this is where modals are mounted */}
      <Outlet />
    </>
  );
};

export default App;
