import type { ReactNode } from "react";

import styles from "./CircleButton.module.css";

interface CircleButtonProps {
  onClick: () => void;
  text?: string;
  Icon?: ReactNode;
}

const CircleButton = ({ onClick, text, Icon }: CircleButtonProps) => (
  <button
    className={`${styles.Button} ${Icon && styles.Icon}`}
    onClick={onClick}
  >
    {text ?? Icon}
  </button>
);

export default CircleButton;
