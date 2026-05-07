import {
  createContext,
  useContext,
  useEffect,
  useRef,
  useState,
  type ReactNode,
} from "react";

import type { SvgIconComponent } from "@mui/icons-material";

import styles from "./Menu.module.css";

type MenuContextType = {
  openMenu: string | null;
  setOpenMenu: (name: string | null) => void;
};

const MenuContext = createContext<MenuContextType | null>(null);

const useMenuContext = () => {
  const ctx = useContext(MenuContext);
  if (!ctx) throw new Error("Menu.* components must be used within <Menu>");
  return ctx;
};

const Menu = ({ children }: { children: ReactNode }) => {
  const [openMenu, setOpenMenu] = useState<string | null>(null);

  const ref = useRef<HTMLElement>(null);

  useEffect(() => {
    const handleClickOutside = (e: MouseEvent) => {
      if (ref.current && !ref.current.contains(e.target as Node)) {
        setOpenMenu(null);
      }
    };
    document.addEventListener("mousedown", handleClickOutside);
    return () => document.removeEventListener("mousedown", handleClickOutside);
  }, []);

  const processedChildren = Array.isArray(children)
    ? children.flatMap((child, index) => [
        child,
        index < children.length - 1 && <p key={`separator-${index}`}>I</p>,
      ])
    : children;

  return (
    <MenuContext.Provider value={{ openMenu, setOpenMenu }}>
      <nav ref={ref} className={styles.Menu}>
        {processedChildren}
      </nav>
    </MenuContext.Provider>
  );
};

const Item = ({ name, children }: { name: string; children: ReactNode }) => {
  const { openMenu, setOpenMenu } = useMenuContext();
  const isOpen = openMenu === name;

  return (
    <div className={styles.Item}>
      <p
        className={`${styles.Text} ${isOpen ? styles.DropdownOpen : ""}`}
        onClick={() => setOpenMenu(isOpen ? null : name)}
      >
        {name}
      </p>
      {isOpen && <div className={styles.Dropdown}>{children}</div>}
    </div>
  );
};

const Option = ({ name, onClick }: { name: string; onClick?: () => void }) => (
  <p className={styles.Text} onClick={onClick}>
    {name}
  </p>
);

const Icon = ({
  icon: IconComponent,
  onClick,
}: {
  icon: SvgIconComponent;
  onClick?: () => void;
}) => (
  <div className={styles.Icon}>
    <IconComponent onClick={onClick} />
  </div>
);

Menu.Item = Item;
Menu.Option = Option;
Menu.Icon = Icon;

export default Menu;
