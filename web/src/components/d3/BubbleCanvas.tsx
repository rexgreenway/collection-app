import { useRef, useLayoutEffect, useEffect, useState } from "react";

import * as d3 from "d3";

import styles from "./Bubble.module.css";
import type { BubbleNode } from "./types";
/**
 * Bubble defines a component that renders a D3.js powered Bubble Plot given
 * children elements that satisfy the Node interface.
 *
 */
const BubbleCanvas = ({ children }: { children: BubbleNode[] }) => {
  // divRef: references plot's container
  const divRef = useRef<HTMLDivElement>(null);
  // canvasRef: references the d3 canvas (necessary as React and D3 manipulate the DOM)
  const canvasRef = useRef<HTMLCanvasElement>(null);

  const [width, setWidth] = useState(300);
  const [height, setHeight] = useState(300);

  const handleResize = () => {
    setWidth(divRef.current ? divRef.current.offsetWidth : 300);
    setHeight(divRef.current ? divRef.current.offsetHeight : 300);
  };

  useEffect(() => {
    window.addEventListener("resize", handleResize);
    return () => window.removeEventListener("resize", handleResize);
  }, []);

  // const r = 20;

  useLayoutEffect(() => {
    handleResize();

    // Copy children nodes for manipulation by d3
    const nodes = children.map((c) => ({ ...c }));
    // Specify the color scale.
    const color = d3.scaleOrdinal(d3.schemeCategory10);

    // Get Canvas by ref
    const canvasNode = canvasRef.current;
    if (!canvasNode) return;

    // Get canvas context
    const context = canvasNode.getContext("2d");
    if (!context) return;

    // Set Canvas width and height
    d3.select(canvasNode).attr("width", width).attr("height", height);

    // Function that actually draws the shapes
    const drawRectangle = (d: BubbleNode, i: number) => {
      context.beginPath();
      context.arc(i * 50 + 50, 150, 3 * d.radius, 0, 2 * Math.PI);
      context.fillStyle = color(d.group);
      context.fill();
    };

    // Use a detached <div> as a virtual container
    const virtual = d3.create("div");

    // Standard data join on virtual elements
    // On a made up "custom element"
    virtual
      .selectAll("custom")
      .data(nodes)
      .join(
        (enter) => enter.append("custom").each((d, i) => drawRectangle(d, i)),
        // (update) => update.each((d, i) => drawRectangle(d, i)),
        (exit) => exit.remove(),
      );

    // GAP
  }, [children, height, width]);

  // Must set an explicit height for the top level div (no % values) otherwise
  // run into a `Maximum update depth exceeded.` Error.
  return (
    <div id="bubble" className={styles.Bubble} ref={divRef}>
      <canvas ref={canvasRef} />
    </div>
  );
};

export default BubbleCanvas;
