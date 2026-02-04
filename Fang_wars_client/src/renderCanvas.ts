import { sendMove } from "./serverInterface";
import ws from "./wsClient";

type Coordinates = {
  x: number;
  y: number;
};

function drawHero(
  coordinates: Coordinates,
  size: number,
  ctx: CanvasRenderingContext2D,
) {
  ctx.fillStyle = "rgb(0 0 200)";
  ctx.fillRect(coordinates.x, coordinates.y, size, size);
}

export default function renderCanvas() {
  const canvas = document.getElementById("canvas") as HTMLCanvasElement | null;
  if (!canvas) return;

  var ctx: CanvasRenderingContext2D | null;
  ctx = canvas.getContext("2d");
  if (!ctx) return;

  console.log("function called", ctx);
  var heroCoordinates = {
    x: 100,
    y: 100,
  };
  var size = 10;

  const step = 10;
  console.log(
    "🚀 ~ draw ~ heroX, heroY, size, size:",
    heroCoordinates,
    size,
    size,
  );

  drawHero(heroCoordinates, size, ctx);

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

  ws.onmessage = (e) => {
    if (!ctx) return;

    switch (e.data) {
      case "F":
        heroCoordinates.y -= step;
        break;
      case "B":
        heroCoordinates.y += step;
        break;
      case "L":
        heroCoordinates.x -= step;
        break;
      case "R":
        heroCoordinates.x += step;
        break;
    }

    ctx.clearRect(0, 0, 800, 600);
    drawHero(heroCoordinates, size, ctx);
    console.log("🚀 ~ renderCanvas ~ onmessage:", e.data, heroCoordinates);
  };
}
