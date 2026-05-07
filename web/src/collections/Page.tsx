import { useMemo } from "react";
import { useNavigate } from "react-router";
import { Add } from "@mui/icons-material";

import { useCollectionStore } from "../store/collection";
import { useViewStore } from "../store/view";

import PageLayout from "../components/PageLayout";
import CollectionBubbleChart from "../components/d3/CollectionBubble";
import CircleButton from "../components/ui/CircleButton";

import styles from "./Collections.module.css";

// Fetches Data from the API ahead of render
export const collectionsLoader = async () => {
  await useCollectionStore.getState().fetchCollections();
  return null;
};

const Page = () => {
  const navigate = useNavigate();

  const view = useViewStore((s) => s.view);

  // subscribes to the Zustand store (live data)
  const collections = useCollectionStore((s) => s.collections);

  const bubbleData = useMemo(
    () => collections.map((c) => ({ group: c.name, radius: 2 })),
    [collections],
  );

  return (
    <PageLayout>
      {/* WANT THIS TO BE AN OUTLET DEPENDING ON VIEW */}
      {view === "bubble" && (
        <CollectionBubbleChart>{bubbleData}</CollectionBubbleChart>
      )}
      {/* {view === "table" && <CollectionTable />} */}

      <div className={styles.ActionButtons}>
        <CircleButton onClick={() => navigate("create")} icon={Add} />
      </div>
    </PageLayout>
  );
};

export default Page;
