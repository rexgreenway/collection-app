import { useNavigate } from "react-router";
import { Switch, ToggleButton, ToggleButtonGroup } from "@mui/material";
import { Settings, DarkMode, LightMode } from "@mui/icons-material";

import { useTheme } from "../contexts";

import Modal from "../components/ui/Modal";
import {
  COLLECTIONS_CONTEXT,
  Settings as CollectionSettings,
} from "../collections";
import { ITEMS_CONTEXT, Settings as ItemsSettings } from "../items";

import styles from "./Settings.module.css";

type SettingsContext = "collections" | "items";

// Should this be a Menu item?? How to link nav and these components?
const SettingsModal = ({ context }: { context: SettingsContext }) => {
  const navigate = useNavigate();

  return (
    <Modal className={styles.Modal} close={() => navigate("..")}>
      <Modal.Title
        icon={Settings}
        title="Settings"
        subtitle={`Change ${context} Settings Here`}
      />

      <Modal.Line />

      <Modal.Section className={styles.Section}>
        <ChangeThemeToggle />
        {/* <ChangeThemeSwitch /> */}
      </Modal.Section>

      <Modal.Line />

      <ContextSettings context={context} />
    </Modal>
  );
};

const ContextSettings = ({ context }: { context: string }) => {
  switch (context) {
    case COLLECTIONS_CONTEXT:
      return <CollectionSettings />;
    case ITEMS_CONTEXT:
      return <ItemsSettings />;
  }
};

const ChangeThemeToggle = () => {
  const { theme, toggleTheme } = useTheme();

  return (
    <div className={styles.Toggle}>
      <h4>Change Theme:</h4>
      <ToggleButtonGroup value={theme} exclusive onChange={toggleTheme}>
        <ToggleButton value="light">
          <LightMode />
        </ToggleButton>
        <ToggleButton value="dark">
          <DarkMode />
        </ToggleButton>
      </ToggleButtonGroup>
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
