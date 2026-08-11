import { useLocation } from "react-router";
import { Switch } from "@mui/material";
import { Settings, DarkMode, LightMode } from "@mui/icons-material";

import { useTheme } from "../contexts";

import { useModalStore } from "../store/modal";

import Toggle from "../components/mui/Toggle";
import Modal from "../components/ui/Modal";

import {
  COLLECTIONS_PATH,
  Settings as CollectionSettings,
} from "../collections";
import { ITEMS_PATH, Settings as ItemsSettings } from "../items";

import styles from "./Settings.module.css";

// Should this be a Menu item?? How to link nav and these components?
const SettingsModal = () => {
  const location = useLocation();

  return (
    <Modal className={styles.Modal} close={useModalStore((s) => s.closeModal)}>
      <Modal.Title
        icon={Settings}
        title="Settings"
        // subtitle={`Change ${context} Settings Here`}
      />

      <Modal.Line />

      <Modal.Section className={styles.Section}>
        <ChangeThemeToggle />
        {/* <ChangeThemeSwitch /> */}
      </Modal.Section>

      <Modal.Line />

      <ContextSettings context={location.pathname} />
    </Modal>
  );
};

const ContextSettings = ({ context }: { context: string }) => {
  switch (context.split("/").pop()) {
    case COLLECTIONS_PATH:
      return <CollectionSettings />;
    case ITEMS_PATH:
      return <ItemsSettings />;
  }
};

const ChangeThemeToggle = () => {
  const { theme, toggleTheme } = useTheme();

  return (
    <div className={styles.Toggle}>
      <h4>Change Theme:</h4>
      <Toggle value={theme} exclusive onChange={toggleTheme}>
        <Toggle.Option value="light">
          <LightMode />
        </Toggle.Option>
        <Toggle.Option value="dark">
          <DarkMode />
        </Toggle.Option>
      </Toggle>
    </div>
  );
};

const ChangeThemeSwitch = () => {
  const { theme, toggleTheme } = useTheme();

  return (
    <div className={styles.Toggle}>
      <h4>Change Theme:</h4>
      <LightMode />
      <Switch checked={theme === "dark"} onChange={toggleTheme} />
      <DarkMode />
    </div>
  );
};

export default SettingsModal;
