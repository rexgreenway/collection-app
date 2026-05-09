import { useMemo } from "react";
import { Add } from "@mui/icons-material";

import { useCollectionStore } from "../store/collection";
import { useViewStore } from "../store/view";

import PageLayout from "../layout/PageLayout";
import CircleButton from "../components/ui/CircleButton";
import BubbleChart from "../components/d3/BubbleChart";

import { SimpleBubble } from "./Bubble";
import { useModalStore } from "../store/modal";

const Page = () => {
  const view = useViewStore((s) => s.view);

  const openModal = useModalStore((s) => s.openModal);

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
    <CircleButton onClick={() => openModal("create-collection")} icon={Add} />
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
