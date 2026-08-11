import styles from "./DemoBanner.module.css";

const DemoBanner = () => (
  <div className={styles.Banner}>
    <h4>Warning</h4>
    <p>
      This is a DEMO! Session storage is used and no collections created here
      will be saved.
    </p>
    <h4>Warning</h4>
  </div>
);

export default DemoBanner;
