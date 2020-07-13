import React from "react";

export default function TileMap({ height, width }) {
  const rows = [];
  for (let y = 0; y < height; y++) {
    let row = [];
    for (let x = 0; x < width; x++) {
      row.push(<div className="tile" id={`x${x}-y${y}`}></div>);
    }
    rows.push(row);
  }

  return (
    <div className="tile-map">
      {rows.map((row) => {
        return <div className="tile-row">{row.map((tile) => tile)}</div>;
      })}
    </div>
  );
}
