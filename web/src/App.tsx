import DarkModeIcon from "@mui/icons-material/DarkMode";
import LightModeIcon from "@mui/icons-material/LightMode";

import { useTheme } from "./contexts";

import Header from "./components/Header";
import CollectionsSection from "./Collection";

import styles from "./App.module.css";

const App = () => {
  // THEME CONTEXT
  const { theme, toggleTheme } = useTheme();
  const title =
    theme === "dark" ? "Switch to Light mode" : "Switch to Dark mode";

  return (
    <>
      <div id="app" className={styles.App}>
        <Header />

        <footer>
          <div onClick={toggleTheme} title={title}>
            {theme === "dark" ? <LightModeIcon /> : <DarkModeIcon />}
          </div>
        </footer>
      </div>
      <CollectionsSection />
    </>
  );
};

export default App;
