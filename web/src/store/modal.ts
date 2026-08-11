import { create } from "zustand";

type ModalType =
  | "settings"
  | "create-collection"
  | "create-items"
  | "quickstart"
  | null;

interface ModalState {
  modal: ModalType;
  modalProps: Record<string, unknown>;
  openModal: (modal: ModalType, props?: Record<string, unknown>) => void;
  closeModal: () => void;
}

export const useModalStore = create<ModalState>((set) => ({
  modal: null,
  modalProps: {},
  openModal: (modal, props = {}) => set({ modal, modalProps: props }),
  closeModal: () => set({ modal: null, modalProps: {} }),
}));
