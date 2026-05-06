// src/components/ui/Form.tsx
import type { ReactNode } from "react";

import styles from "./Form.module.css";

interface FormProps {
  onSubmit: (e: React.SubmitEvent<HTMLFormElement>) => void;
  children: ReactNode;
}

const Form = ({ onSubmit, children }: FormProps) => (
  <form className={styles.Form} onSubmit={onSubmit}>
    {children}
  </form>
);

const Field = ({ label, children }: { label: string; children: ReactNode }) => (
  <div>
    <label>
      <h4>{label}</h4>
    </label>
    {children}
  </div>
);

const Actions = ({ children }: { children: ReactNode }) => (
  <div className={styles.Actions}>{children}</div>
);

const Input = (props: React.InputHTMLAttributes<HTMLInputElement>) => (
  <input className={styles.Input} {...props} />
);

const Button = ({
  ...props
}: React.ButtonHTMLAttributes<HTMLButtonElement>) => <button {...props} />;

Form.Field = Field;
Form.Actions = Actions;
Form.Input = Input;
Form.Button = Button;

export default Form;
