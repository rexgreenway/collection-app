import { useEffect, useMemo, useState, type ReactNode } from "react";

import type { Theme } from "./types";

import { ThemeContext } from "./ThemeContext";

const ThemeProvider = ({ children }: { children: ReactNode }) => {
  const deviceDarkTheme = useMemo(
    () => window.matchMedia("(prefers-color-scheme: dark)").matches,
    [],
  );

  // Sets to device theme on initial render:
  useEffect(() => {
    const r = document.querySelector(":root");
    r!.setAttribute("data-theme", deviceDarkTheme ? "dark" : "light");
  }, [deviceDarkTheme]);

  const [theme, setTheme] = useState<Theme>(deviceDarkTheme ? "dark" : "light");

  const toggleTheme = () => {
    const newTheme = theme === "light" ? "dark" : "light";

    const r = document.querySelector(":root");
    r!.setAttribute("data-theme", newTheme);

    setTheme(newTheme);
  };

  return (
    <ThemeContext.Provider value={{ theme, toggleTheme }}>
      {children}
    </ThemeContext.Provider>
  );
};

export default ThemeProvider;
