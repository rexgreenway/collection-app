import Menu from "./ui/Menu";

import Logo from "../assets/collection-app-v1.svg?react";

import styles from "./Header.module.css";
import Settings from "../settings/Settings";

const Header = () => {
  return (
    <>
      <header className={styles.Header}>
        {/* LEFT */}
        <div className={styles.HeaderLeft}>
          {/* Logo */}
          <Logo className={styles.Logo} />

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
        <div>
          <Settings />
        </div>
      </header>
    </>
  );
};

export default Header;
