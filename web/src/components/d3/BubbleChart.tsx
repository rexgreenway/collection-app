import {
  useRef,
  useEffect,
  useState,
  useLayoutEffect,
  type ComponentType,
} from "react";

import * as d3 from "d3";

import type { BubbleNode } from "./types";

import styles from "./Bubble.module.css";

const BubbleChart = <T extends BubbleNode>({
  data,
  element: Element,
}: {
  data: T[];
  element: ComponentType<T>;
}) => {
  // divRef: references plot's container
  const divRef = useRef<HTMLDivElement>(null);

  // Responsive plot sizing
  const [width, SetWidth] = useState(300);
  const [height, SetHeight] = useState(300);
  const handleResize = () => {
    SetWidth(divRef.current ? divRef.current.offsetWidth : 300);
    SetHeight(divRef.current ? divRef.current.offsetHeight : 300);
  };

  // Hook watching for window resizing
  useEffect(() => {
    handleResize();
    window.addEventListener("resize", handleResize);
    return () => window.removeEventListener("resize", handleResize);
  }, []);

  // Radius Scaling
  const r = 50;
  const small_r = 30;

  // Color scale persists across renders
  // const color = useMemo(() => d3.scaleOrdinal(d3.schemeCategory10), []);

  // Simulated node positions
  const [simulatedNodes, setSimulatedNodes] = useState<T[]>([]);

  // Run d3 simulation (maths only, no DOM manipulation)
  useLayoutEffect(() => {
    const nodes: T[] = data.map((c) => ({ ...c }));

    const simulation = d3
      .forceSimulation(nodes)
      // Centering nodes toward (0, 0) per node
      .force("x", d3.forceX().strength(0.08))
      .force("y", d3.forceY().strength(0.08))
      // Force for 'border' of nodes creates collisions
      .force(
        "collide",
        d3
          .forceCollide<BubbleNode>((d) => r * d.radius + (r - small_r))
          .iterations(12),
      )
      // Inter-node gravity
      .force(
        "charge",
        d3.forceManyBody<BubbleNode>().strength((d) => r * d.radius),
      )
      .on("tick", () => {
        // Triggers rerender of whole react component
        // okay for up to ~100 elements
        setSimulatedNodes([...nodes]);
      });

    return () => {
      simulation.stop();
    };
  }, [data]);

  // Must set an explicit height for the top level div (no % values) otherwise
  // run into a `Maximum update depth exceeded.` Error.
  return (
    <div className={styles.Bubble} ref={divRef}>
      <svg
        className="m-auto"
        width={width}
        height={height}
        viewBox={`${-width / 2} ${-height / 2} ${width} ${height}`}
      >
        {simulatedNodes.map((d) => (
          // Wrapped in g that handles translation
          <g key={d.group} transform={`translate(${d.x}, ${d.y})`}>
            <Element {...d} />
          </g>
        ))}
      </svg>
    </div>
  );
};

export default BubbleChart;
