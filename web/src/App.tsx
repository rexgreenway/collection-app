import DarkModeIcon from "@mui/icons-material/DarkMode";
import LightModeIcon from "@mui/icons-material/LightMode";

import { useTheme } from "./contexts";

import CollectionsSection from "./Collection";

import styles from "./App.module.css";

const App = () => {
  // THEME CONTEXT
  const { theme, toggleTheme } = useTheme();
  const title =
    theme === "dark" ? "Switch to Light mode" : "Switch to Dark mode";

  return (
    <div id="app" className={styles.App}>
      {/* Header */}
      <header>
        <h1>Collection App</h1>
      </header>

      <CollectionsSection />

      <footer>
        <div onClick={toggleTheme} title={title}>
          {theme === "dark" ? <LightModeIcon /> : <DarkModeIcon />}
        </div>
      </footer>
    </div>
  );
};

export default App;
