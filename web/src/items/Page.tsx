import { useNavigate } from "react-router";
import { Add } from "@mui/icons-material";

import PageLayout from "../components/PageLayout";
import CircleButton from "../components/ui/CircleButton";

import styles from "./Items.module.css";

export const itemsLoader = async () => {
  // await useItemStore.getState().fetchCollections();
  return null;
};

const Page = () => {
  const navigate = useNavigate();

  return (
    <PageLayout>
      <h2>ITEMS PAGE</h2>

      <div className={styles.ActionButtons}>
        <CircleButton onClick={() => navigate("create")} icon={Add} />
      </div>
    </PageLayout>
  );
};

export default Page;
