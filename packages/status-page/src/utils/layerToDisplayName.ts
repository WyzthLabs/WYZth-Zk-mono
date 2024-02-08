import { Layer } from "../domain/layer";

export const layerToDisplayName = (layer: Layer) =>
  layer === Layer.Two ? "WYzth_zkevm L2" : "WYzth_zkevm L3";
