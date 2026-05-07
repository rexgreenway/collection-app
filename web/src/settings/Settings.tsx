import { useNavigate } from "react-router";
import { DarkMode, LightMode } from "@mui/icons-material";

import Modal from "../components/ui/Modal";

import styles from "./Settings.module.css";
import { useTheme } from "../contexts";

// Should this be a Menu item?? How to link nav and these components?
const SettingsModal = () => {
  const navigate = useNavigate();

  const { theme, toggleTheme } = useTheme();
  const title =
    theme === "dark" ? "Switch to Light mode" : "Switch to Dark mode";

  return (
    <Modal className={styles.Modal} close={() => navigate("/")}>
      <h2>Settings</h2>

      <p>Change Theme: </p>
      <div onClick={toggleTheme} title={title}>
        {theme === "dark" ? <LightMode /> : <DarkMode />}
      </div>
    </Modal>
  );
};

export default SettingsModal;
