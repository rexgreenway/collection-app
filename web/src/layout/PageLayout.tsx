import type { ReactNode } from "react";

import styles from "./Layout.module.css";

/**
 * A page layout component that renders main content & an actions panel.
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
    </>
  );
};

export default PageLayout;
