import {
  ToggleButton,
  ToggleButtonGroup,
  type ToggleButtonGroupProps,
} from "@mui/material";
import { type ReactNode } from "react";

import styles from "./Toggle.module.css";

interface ToggleProps extends ToggleButtonGroupProps {
  children: ReactNode;
}

const Toggle = ({ children, ...rest }: ToggleProps) => (
  <ToggleButtonGroup {...rest}>{children}</ToggleButtonGroup>
);

const Option = ({
  children,
  value,
}: {
  children: ReactNode;
  value: string;
}) => (
  <ToggleButton
    value={value}
    classes={{
      root: styles.Button,
      selected: styles.ButtonSelected,
    }}
  >
    {children}
  </ToggleButton>
);

Toggle.Option = Option;

export default Toggle;
