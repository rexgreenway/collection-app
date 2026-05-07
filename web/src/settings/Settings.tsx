import { useState } from "react";
import {
  Settings as SettingsIcon,
  DarkMode,
  LightMode,
} from "@mui/icons-material";

import Menu from "../components/ui/Menu";
import Modal from "../components/ui/Modal";

import styles from "./Settings.module.css";
import { useTheme } from "../contexts";

// Should this be a Menu item?? How to link nav and these components?
const Settings = () => {
  const { theme, toggleTheme } = useTheme();
  const title =
    theme === "dark" ? "Switch to Light mode" : "Switch to Dark mode";

  const [modalOpen, setModalOpen] = useState(false);

  return (
    <>
      <Menu>
        <Menu.Icon icon={SettingsIcon} onClick={() => setModalOpen(true)} />
      </Menu>

      {modalOpen && (
        <Modal className={styles.Modal} close={() => setModalOpen(false)}>
          <h2>Settings</h2>

          <p>Change Theme: </p>
          <div onClick={toggleTheme} title={title}>
            {theme === "dark" ? <LightMode /> : <DarkMode />}
          </div>
        </Modal>
      )}
    </>
  );
};

export default Settings;
