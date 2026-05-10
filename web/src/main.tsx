import { StrictMode } from "react";
import { createRoot } from "react-dom/client";
import { StyledEngineProvider } from "@mui/material";

import { ThemeProvider } from "./contexts";

import Router from "./Router.tsx";

import "./index.css";

createRoot(document.getElementById("root")!).render(
  <StrictMode>
    <ThemeProvider>
      <StyledEngineProvider injectFirst>
        <Router />
      </StyledEngineProvider>
    </ThemeProvider>
  </StrictMode>,
);
