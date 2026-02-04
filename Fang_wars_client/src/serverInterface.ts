import ws from "./wsClient.ts";

function sendMove(direction: string) {
  console.log("🚀 ~ sendMove ~ direction:", direction);

  ws.send(JSON.stringify({ type: "move", direction }));
}

export { sendMove };
