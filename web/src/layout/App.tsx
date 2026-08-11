import { Outlet } from "react-router";

import Header from "./Header";
import Footer from "./Footer";
import ModalOutlet from "./ModalOutlet";

import DemoBanner from "../demo/DemoBanner";

const App = () => {
  const demo = true;

  return (
    <div id="app">
      {demo && <DemoBanner />}
      <Header />
      <Outlet />
      <Footer />
      <ModalOutlet />
    </div>
  );
};

export default App;
