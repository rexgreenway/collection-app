import type { SvgIconComponent } from "@mui/icons-material";

import styles from "./CircleButton.module.css";

interface CircleButtonProps {
  onClick: () => void;
  text?: string;
  icon?: SvgIconComponent;
}

const CircleButton = ({
  onClick,
  text,
  icon: IconComponent,
}: CircleButtonProps) => (
  <button
    className={`${styles.Button} ${IconComponent && styles.Icon}`}
    onClick={onClick}
  >
    {text ?? (IconComponent && <IconComponent />)}
  </button>
);

export default CircleButton;
