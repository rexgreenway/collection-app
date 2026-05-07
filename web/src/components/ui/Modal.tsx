import { type ReactNode, useEffect } from "react";
import CloseRoundedIcon from "@mui/icons-material/CloseRounded";

import styles from "./Modal.module.css";

interface ModalProps {
  close: () => void;
  allowClose?: boolean;
  children?: ReactNode | ReactNode[];
  className?: string;
}

const Modal = ({
  close,
  allowClose = true,
  children,
  className,
}: ModalProps) => {
  // Add ability to close with Escape Key
  if (allowClose) {
    useEffect(() => {
      const handleEscKey = (event: KeyboardEvent) => {
        if (event.key === "Escape") {
          close();
        }
      };
      document.addEventListener("keydown", handleEscKey);
      return () => {
        document.removeEventListener("keydown", handleEscKey);
      };
    }, [close]);
  }

  return (
    <div id="modal" className={styles.PageBackground}>
      {/* DON'T KNOW IF I WANT THIS CLOSE FUNCTIONALITY */}
      {/* MAYBE A CANCEL INSTEAD */}
      {allowClose && (
        <CloseRoundedIcon
          fontSize="large"
          className={styles.CloseButton}
          onClick={close}
        />
      )}
      <div className={`${styles.Modal} ${className}`}>{children}</div>
    </div>
  );
};

export default Modal;
