import { useMemo } from "react";
import { Add, LegendToggle } from "@mui/icons-material";

import { useModalStore } from "../store/modal";
import { useCollectionStore } from "../store/collection";
import { useViewStore } from "../store/view";

import type { Collection } from "../api/types";

import PageLayout from "../layout/PageLayout";
import CircleButton from "../components/ui/CircleButton";

import BubbleChart from "../components/d3/BubbleChart";
import type { BubbleNode } from "../components/d3/types";
import { SimpleBubble } from "./Bubble";

import BasicTable from "../components/mui/Table";

import styles from "./Collections.module.css";

const Page = () => {
  const { view, setView } = useViewStore();

  const openModal = useModalStore((s) => s.openModal);

  // subscribes to the Zustand store (live data)
  const collections = useCollectionStore((s) => s.collections);

  const collectionData: (Collection & BubbleNode)[] = useMemo(
    () =>
      collections.map((c) => ({
        ...c,
        group: c.name,
        radius: 2,
      })),
    [collections],
  );

  const toggleView = () => {
    switch (view) {
      case "bubble":
        setView("table");
        break;
      case "table":
        setView("bubble");
        break;
    }
  };

  const collectionActions = [
    <CircleButton
      key="create"
      onClick={() => openModal("create-collection")}
      icon={Add}
    />,
    <CircleButton key="toggle-view" onClick={toggleView} icon={LegendToggle} />,
  ];

  return (
    <PageLayout actions={collectionActions}>
      {view === "bubble" && (
        <BubbleChart data={collectionData} element={SimpleBubble} />
      )}
      {view === "table" && (
        <BasicTable data={collectionData} className={styles.Table} />
      )}
    </PageLayout>
  );
};

export default Page;
