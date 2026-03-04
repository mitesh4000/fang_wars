import { sendMove } from "./serverInterface";
import { PIXEL_SIZE } from "../config";
import ws from "./wsClient";

type PixelCoordinates = {
  row: number;
  column: number;
};

var headDirection = "U";
var bodyArray = [
  { column: 1, row: 1 },
  { column: 1, row: 2 },
  { column: 1, row: 3 },
];

function renderPixel(
  ctx: CanvasRenderingContext2D,
  pixelCordinatex: number,
  pixelCordinatey: number,
  size: number,
  color: string,
) {
  ctx.fillStyle = color;
  ctx.fillRect(pixelCordinatex, pixelCordinatey, size, size);
}

type SnakeState = {
  headDirection: string;
  bodyArray: Array<PixelCoordinates>;
};

function renderSnake(
  state: SnakeState, // Array<PixelCoordinates>,
  ctx: CanvasRenderingContext2D,
) {
  console.log("🚀 ~ state:", state.bodyArray);
  for (let i = 0; i < state.bodyArray.length; i++) {
    renderDotOnGrid(state.bodyArray[i], ctx);
  }
}

function renderDotOnGrid(
  cords: { row: number; column: number },
  ctx: CanvasRenderingContext2D,
) {
  renderPixel(
    ctx,
    (PIXEL_SIZE + 1) * cords.row,
    (PIXEL_SIZE + 1) * cords.column,
    PIXEL_SIZE,
    "rgb(98, 191, 121)",
  );
  // ctx.fillRect((size + 1) * row, (size + 1) * column, size, size);
}

function renderGrid(size: number, ctx: CanvasRenderingContext2D) {
  for (let i = 0; i <= 10; i++) {
    for (let j = 0; j <= 10; j++) {
      renderPixel(
        ctx,
        (size + 1) * j,
        (size + 1) * i,
        size,
        "rgb(216, 232, 220)",
      );
      // ctx.fillRect((size + 1) * j, (size + 1) * i, size, size);
    }
  }
}

export default function game() {
  const canvas = document.getElementById("canvas") as HTMLCanvasElement | null;
  if (!canvas) return;

  var context: CanvasRenderingContext2D | null;
  context = canvas.getContext("2d");
  if (!context) return;

  var size = 50;

  ws.onmessage = (event) => {
    if (!context) return;
    const data = JSON.parse(event.data);

    renderGrid(size, context);
    renderSnake(data, context);
  };

  document.addEventListener("keydown", (e) => {
    switch (e.key) {
      case "ArrowUp":
      case "w":
      case "W":
        sendMove("F");
        break;
      case "ArrowDown":
      case "s":
      case "S":
        sendMove("B");
        break;
      case "ArrowLeft":
      case "a":
      case "A":
        sendMove("L");
        break;
      case "ArrowRight":
      case "d":
      case "D":
        sendMove("R");
        break;
    }

    e.preventDefault();
  });
}
