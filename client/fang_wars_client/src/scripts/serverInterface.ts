import ws from "./wsClient.ts";

function sendMove(direction: string) {
  ws.send(JSON.stringify({ type: "move", direction }));
}

export { sendMove };
