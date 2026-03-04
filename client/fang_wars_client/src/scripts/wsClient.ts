var ws = new WebSocket("ws://localhost:8080/ws");
ws.onopen = () => console.log("Connected");
ws.onerror = (e) => console.error(e);

export default ws;
