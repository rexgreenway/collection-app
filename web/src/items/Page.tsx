import { useNavigate } from "react-router";
import { Add } from "@mui/icons-material";

import PageLayout from "../layout/PageLayout";
import CircleButton from "../components/ui/CircleButton";

import styles from "./Items.module.css";

export const itemsLoader = async () => {
  // await useItemStore.getState().fetchCollections();
  return null;
};

const Page = () => {
  const navigate = useNavigate();

  const itemsActions = (
    <CircleButton onClick={() => navigate(".")} icon={Add} />
  );

  return (
    <PageLayout actions={itemsActions}>
      <h2>ITEMS PAGE</h2>
    </PageLayout>
  );
};

export default Page;
