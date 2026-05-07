import { useLocation, useNavigate } from "react-router";
import { Settings } from "@mui/icons-material";

import Menu from "../components/ui/Menu";

import Logo from "../assets/collection-app-v1.svg?react";

import styles from "./Layout.module.css";

const Header = () => {
  const navigate = useNavigate();
  const location = useLocation();

  return (
    <>
      <header className={styles.Header}>
        {/* LEFT */}
        <div className={styles.HeaderLeft}>
          {/* Logo */}
          <Logo className={styles.Logo} onClick={() => navigate("/")} />

          <Menu>
            <Menu.Item name="File">
              <Menu.Option name="Save" />
              <Menu.Option name="Export" />
            </Menu.Item>
            {/* <Menu.Item name="Edit">
              <Menu.Option name="Undo" />
              <Menu.Option name="Redo" />
            </Menu.Item> */}
            <Menu.Item name="Help">
              <Menu.Option name="Quick Start" />
            </Menu.Item>
          </Menu>
        </div>

        {/* RIGHT */}
        <div className={styles.HeaderRight}>
          <Menu>
            <Menu.Icon
              icon={Settings}
              onClick={() => navigate(`${location.pathname}/settings`)}
            />
          </Menu>
        </div>
      </header>
    </>
  );
};

export default Header;
