import type { SimulationNodeDatum } from "d3";

/**
 * Node is a helper interface for rendering D3 Simulations with React & Typescript.
 *
 * This interface extends the D3.js type SimulationNodeDatum that is used by D3
 * simulations to dynamically update the positions of elements in the DOM. The
 * new required parameters allow for custom grouping of Nodes & defining how
 * the radii of Nodes are calculated.
 */
export interface BubbleNode extends SimulationNodeDatum {
  group: string;
  radius: number;
}
