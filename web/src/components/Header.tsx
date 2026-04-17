import { useState, type ReactNode } from "react";
import styles from "./Header.module.css";

const PillBox = ({ children }: { children: ReactNode | ReactNode[] }) => {
  const processedChildren = Array.isArray(children)
    ? children.flatMap((child, index) => [
        child,
        index < children.length - 1 && <p key={`separator-${index}`}>I</p>,
      ])
    : children;
  return <div className={styles.PillBox}>{processedChildren}</div>;
};

type DropdownMenuState = {
  menu: string;
  isOpen: boolean;
};

const DropdownMenu = ({
  menu,
  children,
}: {
  menu: string;
  children: ReactNode | ReactNode[];
}) => {
  return (
    <div className={styles.DropdownMenu}>
      <h3>{menu}</h3>
      {children}
    </div>
  );
};

const Header = () => {
  const [menuState, setMenuState] = useState<DropdownMenuState>({
    menu: "file",
    isOpen: false,
  });

  const openMenu = (menu: string) => {
    setMenuState({ menu, isOpen: true });
  };

  return (
    <>
      <header className={styles.Header}>
        <div></div>
        {/* LEFT */}
        <div className={styles.HeaderLeft}>
          {/* Logo */}
          <h2>Collection App</h2>

          {/* Main Pill */}
          <PillBox>
            <h5 onClick={() => openMenu("file")}>File</h5>
            <h5 onClick={() => openMenu("edit")}>Edit</h5>
            <h5 onClick={() => openMenu("options")}>Options</h5>
          </PillBox>
        </div>

        {/* RIGHT */}
        <div>
          {/* End Pill */}
          <PillBox>
            {/* Replace with clickable dropdown*/}
            <h5 onClick={() => openMenu("settings")}>Settings</h5>
          </PillBox>
        </div>
      </header>
      {/* Dropdown */}
      {menuState.isOpen && (
        <DropdownMenu menu={menuState.menu}>
          <p>Test 1</p>
          <p>Test 2</p>
        </DropdownMenu>
      )}
    </>
  );
};

export default Header;
