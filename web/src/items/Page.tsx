import { Add } from "@mui/icons-material";

import { useModalStore } from "../store/modal";
import { useItemStore } from "../store/items";

import PageLayout from "../layout/PageLayout";
import CircleButton from "../components/ui/CircleButton";

import styles from "./Items.module.css";

const Page = () => {
  const openModal = useModalStore((s) => s.openModal);

  // subscribes to the Zustand store (live data)
  const items = useItemStore((s) => s.items);

  const itemsActions = (
    <CircleButton
      key="create"
      onClick={() => openModal("create-items")}
      icon={Add}
    />
  );

  return (
    <PageLayout actions={itemsActions}>
      <h2>ITEMS PAGE</h2>
      <ul>
        {items.map((item) => (
          <li key={item.id}>{item.name}</li>
        ))}
      </ul>
    </PageLayout>
  );
};

export default Page;
