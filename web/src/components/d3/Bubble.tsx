import { useRef, useLayoutEffect, useEffect, useState } from "react";

import * as d3 from "d3";

import styles from "./Bubble.module.css";
import type { BubbleNode } from "./types";

/**
 * Bubble defines a component that renders a D3.js powered Bubble Plot given
 * children elements that satisfy the Node interface.
 *
 */
const Bubble = ({ children }: { children: BubbleNode[] }) => {
  // divRef: references plot's container
  const divRef = useRef(null);
  // svgRef: references the d3 svg (necessary as React and D3 manipulate the DOM)
  const svgRef = useRef(null);

  // Responsive plot sizing
  const [width, SetWidth] = useState(300);
  const [height, SetHeight] = useState(300);
  const handleResize = () => {
    SetWidth(divRef.current ? divRef.current["offsetWidth"] : 300);
    SetHeight(divRef.current ? divRef.current["offsetHeight"] : 300);
  };

  // Hook watching for window resizing
  useEffect(() => {
    window.addEventListener("resize", handleResize);
    return () => window.removeEventListener("resize", handleResize);
  });

  // Radius Scaling
  const r = 20;

  // Render d3 simulation
  useLayoutEffect(() => {
    handleResize();

    // Copy children nodes for manipulation by d3
    const nodes = children.map((c) => ({ ...c }));

    // Specify the color scale.
    const color = d3.scaleOrdinal(d3.schemeCategory10);

    // Create and start simulation
    const simulation = d3
      .forceSimulation(nodes)
      // Centering nodes toward (0, 0) per node
      .force("x", d3.forceX().strength(0.08))
      .force("y", d3.forceY().strength(0.08))
      // Force for 'border' of nodes creates collisions
      .force(
        "collide",
        d3.forceCollide<BubbleNode>((d) => r * d.radius + 2).iterations(12),
      )
      // Inter-node gravity
      .force(
        "charge",
        d3.forceManyBody<BubbleNode>().strength((d) => r * d.radius),
      );

    // Establish SVG sizing and ViewBox
    const svgElement = d3
      .select(svgRef.current)
      .attr("width", width)
      .attr("height", height)
      .attr("viewBox", [-width / 2, -height / 2, width, height]);

    // Join Node Data to simulation as circles
    const node = svgElement
      .selectAll<SVGCircleElement, SVGCircleElement>("circle")
      .data<BubbleNode>(nodes)
      .join("circle")
      .attr("r", (d) => r * d.radius)
      .attr("fill", (d) => color(d.group));

    // Reheat the simulation when drag starts, and fix the subject position.
    function dragStart(
      event: d3.D3DragEvent<SVGCircleElement, SVGCircleElement, BubbleNode>,
    ) {
      if (!event.active) simulation.alphaTarget(0.4).restart();
      event.subject.fx = event.subject.x;
      event.subject.fy = event.subject.y;
    }

    // Update the subject (dragged node) position during drag.
    function dragged(
      event: d3.D3DragEvent<SVGCircleElement, SVGCircleElement, BubbleNode>,
    ) {
      event.subject.fx = event.x;
      event.subject.fy = event.y;
    }

    // Restore the target alpha so the simulation cools after dragging ends.
    // Unfix the subject position now that it’s no longer being dragged.
    function dragEnd(
      event: d3.D3DragEvent<SVGCircleElement, SVGCircleElement, BubbleNode>,
    ) {
      if (!event.active) simulation.alphaTarget(0);
      event.subject.fx = null;
      event.subject.fy = null;
    }
    // Add drag behaviour
    node.call(
      d3
        .drag<SVGCircleElement, BubbleNode>()
        .on("start", dragStart)
        .on("drag", dragged)
        .on("end", dragEnd),
    );

    // Change node on tick
    function ticked() {
      node.attr("cx", (d) => d.x!).attr("cy", (d) => d.y!);
    }

    // Turn on Simulation
    simulation.on("tick", ticked);
  }, [children, height, width]);

  // Must set an explicit height for the top level div (no % values) otherwise
  // run into a `Maximum update depth exceeded.` Error.
  return (
    <div className={styles.Bubble} ref={divRef}>
      <svg className="m-auto" ref={svgRef} />
    </div>
  );
};

export default Bubble;
