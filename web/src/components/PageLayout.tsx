// components/PageLayout.tsx
import type { ReactNode } from "react";
import { Outlet } from "react-router";

const PageLayout = ({ children }: { children: ReactNode }) => {
  return (
    <>
      {children}

      {/* Modal Mount */}
      <Outlet />
    </>
  );
};

export default PageLayout;
