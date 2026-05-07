import { Outlet } from "react-router";

import Header from "./components/Header";

import styles from "./App.module.css";

const App = () => {
  return (
    <div id="app" className={styles.App}>
      <Header />

      <Outlet />
    </div>
  );
};

export default App;
