import { useNavigate } from "react-router";

import type { CollectionResponse } from "../api/types";

import type { BubbleNode } from "../components/d3/types";

import styles from "./Bubble.module.css";

export const SimpleBubble = (node: CollectionResponse & BubbleNode) => {
  const navigate = useNavigate();

  return (
    <g className={styles.Bubble} onClick={() => navigate(`./${node.id}`)}>
      {/* BUBBLE ITSELF */}
      <circle className={styles.Circle} r={node.radius} fill="blue" />

      {/* BUBBLE NAME */}
      <text
        textAnchor="middle"
        dominantBaseline="central"
        className={styles.Text}
      >
        {node.name}
      </text>

      {/* NO. OF ITEMS IN COLLECTION */}
      {node.itemCount && (
        <text
          textAnchor="middle"
          dominantBaseline="central"
          className={styles.Size}
        >
          {node.itemCount}
        </text>
      )}
    </g>
  );
};

// export const ComplexBubble = ({ node }: BubbleComponent) => {
//   const r = 50;
//   const small_r = 30;

//   return (
//     <g key={node.group} transform={`translate(${node.x}, ${node.y})`}>
//       {/* Collection -> Group with the circles inside, parent group controls the positions */}
//       <g>
//         <circle r={r * node.radius} />
//         {/* Add */}
//         <circle r={small_r} cx={r * node.radius} fill="red" />
//         {/* Inspect */}
//         <circle
//           r={small_r}
//           cx={(r * node.radius) / Math.sqrt(2)}
//           cy={(r * node.radius) / Math.sqrt(2)}
//           fill="green"
//         />
//         {/* More */}
//         <circle r={small_r} cy={r * node.radius} fill="blue" />
//       </g>
//     </g>
//   );
// };
