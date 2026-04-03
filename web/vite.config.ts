import { defineConfig } from "vite";
import react from "@vitejs/plugin-react";

// https://vite.dev/config/
export default defineConfig({
  plugins: [react()],
  // proxys requests to the go backend
  server: {
    proxy: {
      "/v1/collections": "http://localhost:8089",
    },
  },
});
