import { useNavigate } from "react-router";
import { Switch, ToggleButton, ToggleButtonGroup } from "@mui/material";
import { Settings, DarkMode, LightMode } from "@mui/icons-material";

import { useTheme } from "../contexts";

import Modal from "../components/ui/Modal";

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

      <h2>{context} Specific Settings go here</h2>
    </Modal>
  );
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
