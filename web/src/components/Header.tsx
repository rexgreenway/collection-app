import Menu from "./Menu";

import styles from "./Header.module.css";

const Header = () => {
  return (
    <>
      <header className={styles.Header}>
        {/* LEFT */}
        <div className={styles.HeaderLeft}>
          {/* Logo */}
          <h2>Collection App</h2>

          <Menu>
            <Menu.Item name="File">
              <Menu.Option optionName="Save" />
            </Menu.Item>
            <Menu.Item name="Edit">
              <Menu.Option optionName="Undo" />
              <Menu.Option optionName="Redo" />
            </Menu.Item>
            <Menu.Item name="Options">
              <Menu.Option optionName="Something" />
            </Menu.Item>
          </Menu>
        </div>

        {/* RIGHT */}
        <div>
          {/* Replace with Logo */}
          <h5>Settings</h5>
        </div>
      </header>
    </>
  );
};

export default Header;
