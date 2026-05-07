import { create } from "zustand";

type View = "bubble" | "table";

interface ViewStore {
  view: View;
  setView: (view: View) => void;
}

export const useViewStore = create<ViewStore>((set) => ({
  view: "bubble",
  setView: (view) => set({ view }),
}));
