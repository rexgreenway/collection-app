import { type ReactNode, useEffect } from "react";
import { type SvgIconComponent, CloseRounded } from "@mui/icons-material";

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
        <CloseRounded
          fontSize="large"
          className={styles.CloseButton}
          onClick={close}
        />
      )}
      <div className={`${styles.Modal} ${className}`}>{children}</div>
    </div>
  );
};

const Title = ({
  title,
  subtitle,
  icon: IconComponent,
}: {
  title: string;
  subtitle?: string;
  icon?: SvgIconComponent;
}) => (
  <div className={styles.Title}>
    {IconComponent && <IconComponent />}
    <div>
      <h1>{title}</h1>
      {subtitle && <p>{subtitle}</p>}
    </div>
  </div>
);

const Section = ({
  className,
  children,
  sectionTitle,
}: {
  className: string;
  children: ReactNode;
  sectionTitle?: string;
}) => (
  <section className={`${styles.Section} ${className}`}>
    {sectionTitle && <h3>{sectionTitle}</h3>}
    {children}
  </section>
);

const Line = () => <span className={styles.Line} />;

Modal.Title = Title;
Modal.Section = Section;
Modal.Line = Line;

export default Modal;
