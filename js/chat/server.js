// Using both express and ws to handle regular web stuff and websockets.
// Based on https://www.npmjs.com/package/ws

const http = require('http');
const WebSocket = require('ws');
const urllib = require('url');
const express = require('express');

const server = http.createServer();
const wss = new WebSocket.Server({ server: server });
const app = express();

const PORT = 2428;

// TODO: test out that hello worlds 1 and 2 can happen

app.use((req, res) => {
  res.send({ msg: 'hello world 1' });
});

wss.on('connection', (ws) => {
  ws.on('message', (message) => {
    console.log('received: %s', message);
  });

  ws.send('hello world 2');
});

server.on('request', app);
server.listen(PORT, () => {
  console.log('Listening on ' + server.address().port);
});
