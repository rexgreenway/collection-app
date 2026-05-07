import { Outlet } from "react-router";

import Header from "./Header";
import Footer from "./Footer";

const App = () => {
  return (
    <div id="app">
      <Header />

      <Outlet />

      <Footer />
    </div>
  );
};

export default App;
