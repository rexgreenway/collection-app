import { useMemo } from "react";
import { useNavigate } from "react-router";
import { Add } from "@mui/icons-material";

import { useCollectionStore } from "../store/collection";
import { useViewStore } from "../store/view";

import PageLayout from "../layout/PageLayout";
import CircleButton from "../components/ui/CircleButton";
import BubbleChart from "../components/d3/BubbleChart";

import { SimpleBubble } from "./Bubble";

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

  const bubbles = useMemo(
    () =>
      collections.map((c) => ({
        id: c.id,
        group: c.name,
        radius: 2,
        color: "blue",
        text: "best collection",
        size: 40,
      })),
    [collections],
  );

  const collectionActions = (
    <CircleButton onClick={() => navigate("create")} icon={Add} />
  );

  return (
    <PageLayout actions={collectionActions}>
      {view === "bubble" && (
        <BubbleChart data={bubbles} BubbleComponent={SimpleBubble} />
      )}
      {/* {view === "table" && <CollectionTable />} */}
    </PageLayout>
  );
};

export default Page;
