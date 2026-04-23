import { createContext, useContext, useState, type ReactNode } from "react";

// 1. Shared context so children can communicate with the parent
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

// 2. Parent component provides context
const Menu = ({ children }: { children: ReactNode }) => {
  const [openMenu, setOpenMenu] = useState<string | null>(null);
  return (
    <MenuContext.Provider value={{ openMenu, setOpenMenu }}>
      <nav>{children}</nav>
    </MenuContext.Provider>
  );
};

// 3. Sub-components consume context
const Item = ({ name, children }: { name: string; children: ReactNode }) => {
  const { openMenu, setOpenMenu } = useMenuContext();
  const isOpen = openMenu === name;
  return (
    <div>
      <button onClick={() => setOpenMenu(isOpen ? null : name)}>{name}</button>
      {isOpen && <div>{children}</div>}
    </div>
  );
};

const Option = ({ children }: { children: ReactNode }) => {
  return <div>{children}</div>;
};

// 4. Attach sub-components as static properties
Menu.Item = Item;
Menu.Option = Option;

export default Menu;
