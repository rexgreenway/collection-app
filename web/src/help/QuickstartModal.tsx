import { Help } from "@mui/icons-material";

import { useModalStore } from "../store/modal";

import Modal from "../components/ui/Modal";

const QuickstartModal = () => {
  return (
    <Modal close={useModalStore((s) => s.closeModal)}>
      <Modal.Title icon={Help} title="QuickStart" />

      <Modal.Section>
        <p>Help messages</p>
      </Modal.Section>
    </Modal>
  );
};

export default QuickstartModal;
