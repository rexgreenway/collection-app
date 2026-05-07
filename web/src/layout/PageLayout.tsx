// components/PageLayout.tsx
import type { ReactNode } from "react";
import { Outlet } from "react-router";

import styles from "./Layout.module.css";

/**
 * A page layout component that renders main content, an actions panel, and a modal outlet.
 *
 * @param children - The main content to render in the page body.
 * @param actions - Action elements displayed in the actions panel.
 */
const PageLayout = ({
  children,
  actions,
}: {
  children: ReactNode;
  actions: ReactNode;
}) => {
  return (
    <>
      {/* Main */}
      <div id="main" className={styles.Main}>
        {children}
      </div>

      {/* ACTIONS */}
      <div id="actions" className={styles.Actions}>
        {actions}
      </div>

      {/* Modal Mount */}
      <Outlet />
    </>
  );
};

export default PageLayout;
