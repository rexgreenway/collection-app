import SettingsIcon from "@mui/icons-material/Settings";

import Menu from "./Menu";

import Logo from "../assets/collection-app-v1.svg?react";

import styles from "./Header.module.css";

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
            <Menu.Item name="Edit">
              <Menu.Option name="Undo" />
              <Menu.Option name="Redo" />
            </Menu.Item>
            <Menu.Item name="Options">
              <Menu.Option name="Something" />
            </Menu.Item>
          </Menu>
        </div>

        {/* RIGHT */}
        <div>
          <Menu>
            <Menu.Icon icon={SettingsIcon} />
          </Menu>
        </div>
      </header>
    </>
  );
};

export default Header;
