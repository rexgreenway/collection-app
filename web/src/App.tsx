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
        <button onClick={toggleTheme} title={title}>
          Change Theme
        </button>
      </header>

      <CollectionsSection />

      <footer>
        <h3>FOOTER - links - etc...</h3>
      </footer>
    </div>
  );
};

export default App;
