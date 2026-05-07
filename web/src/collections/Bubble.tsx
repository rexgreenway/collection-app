import type { BubbleComponent } from "../components/d3/types";

import styles from "./Bubble.module.css";

export const SimpleBubble = ({ node }: BubbleComponent) => {
  const r = 50;

  return (
    <g
      className={styles.Bubble}
      key={node.group}
      transform={`translate(${node.x}, ${node.y})`}
    >
      <circle
        className={styles.Circle}
        r={r * node.radius}
        style={node.color ? { fill: node.color } : undefined}
      />
    </g>
  );
};

export const ComplexBubble = ({ node, color }: BubbleComponent) => {
  const r = 50;
  const small_r = 30;

  return (
    <g key={node.group} transform={`translate(${node.x}, ${node.y})`}>
      {/* Collection -> Group with the circles inside, parent group controls the positions */}
      <g>
        <circle r={r * node.radius} fill={color} />
        {/* Add */}
        <circle r={small_r} cx={r * node.radius} fill="red" />
        {/* Inspect */}
        <circle
          r={small_r}
          cx={(r * node.radius) / Math.sqrt(2)}
          cy={(r * node.radius) / Math.sqrt(2)}
          fill="green"
        />
        {/* More */}
        <circle r={small_r} cy={r * node.radius} fill="blue" />
      </g>
    </g>
  );
};
