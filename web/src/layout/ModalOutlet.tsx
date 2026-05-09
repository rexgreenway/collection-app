import { CreateModal } from "../collections";
import QuickstartModal from "../help/QuickstartModal";
import SettingsModal from "../settings/SettingsModal";
import { useModalStore } from "../store/modal";

const ModalOutlet = () => {
  const { modal } = useModalStore();

  switch (modal) {
    case "create-collection":
      return <CreateModal />;
    case "settings":
      return <SettingsModal />;
    case "quickstart":
      return <QuickstartModal />;
    default:
      return null;
  }
};

export default ModalOutlet;
