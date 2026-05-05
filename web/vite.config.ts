import { defineConfig } from "vite";
import react from "@vitejs/plugin-react";
import svgr from "vite-plugin-svgr";

// https://vite.dev/config/
export default defineConfig({
  plugins: [react(), svgr()],
  // proxies requests to the go backend when running locally
  server: {
    proxy: {
      "/v1/collections": "http://localhost:8089",
    },
  },
});
