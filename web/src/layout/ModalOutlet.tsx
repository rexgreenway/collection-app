import { CreateModal } from "../collections";
import { CreateModal as CreateItemModal } from "../items";
import QuickstartModal from "../help/QuickstartModal";
import SettingsModal from "../settings/SettingsModal";
import { useModalStore } from "../store/modal";

const ModalOutlet = () => {
  const { modal } = useModalStore();

  switch (modal) {
    case "create-collection":
      return <CreateModal />;
    case "create-items":
      return <CreateItemModal />;
    case "settings":
      return <SettingsModal />;
    case "quickstart":
      return <QuickstartModal />;
    default:
      return null;
  }
};

export default ModalOutlet;
