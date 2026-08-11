import Table from "@mui/material/Table";
import TableBody from "@mui/material/TableBody";
import TableCell from "@mui/material/TableCell";
import TableContainer from "@mui/material/TableContainer";
import TableHead from "@mui/material/TableHead";
import TableRow from "@mui/material/TableRow";

import styles from "./Table.module.css";

const BasicTable = <T extends object>({
  data,
  className,
}: {
  data: T[];
  className?: string;
}) => {
  if (data.length === 0) return null;

  const columns = Object.keys(data[0]) as (keyof T)[];

  return (
    // <TableContainer className={styles.Container}>
    <TableContainer className={`${styles.Container} ${className ?? ""}`}>
      <Table className={styles.Table}>
        <TableHead>
          <TableRow>
            {columns.map((col) => (
              <TableCell key={String(col)} className={styles.HeaderCell}>
                {String(col)}
              </TableCell>
            ))}
          </TableRow>
        </TableHead>

        <TableBody>
          {data.map((row, i) => (
            <TableRow key={i} className={styles.Row}>
              {columns.map((col) => (
                <TableCell key={String(col)} className={styles.Cell}>
                  {String(row[col])}
                </TableCell>
              ))}
            </TableRow>
          ))}
        </TableBody>
      </Table>
    </TableContainer>
  );
};

export default BasicTable;
