const WebSocket = require('ws');

const ws = new WebSocket('ws://localhost:2428');

ws.on('open', () => {
  ws.send('something');
});

ws.on('message', (data, flags) => {
  console.log('client received:', data);
});
