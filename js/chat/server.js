// Using both express and ws to handle regular web stuff and websockets.
// Based on https://www.npmjs.com/package/ws

const http = require('http');
const WebSocket = require('ws');
const urllib = require('url');
const express = require('express');
const bodyParser = require('body-parser');

const server = http.createServer();
const wss = new WebSocket.Server({ server: server });
const app = express();

const PORT = 2428;

app.use(bodyParser.json());

app.get('/', function (req, res) {
  res.send('this is the chat server')
});

app.post('/broadcast', (req, res) => {
  console.log('handling /broadcast');
  wss.broadcast(req.body);
  res.send('OK');
});

// Takes JSON
wss.broadcast = (data) => {
  let payload = JSON.stringify(data);
  wss.clients.forEach((client) => {
    client.send(payload);
  });
};

wss.on('connection', (ws) => {
  ws.on('message', (message) => {
    console.log('received: %s', message);
  });

  ws.send('you have connected to the websocket server');
});

server.on('request', app);
server.listen(PORT, () => {
  console.log('Listening on ' + server.address().port);
});
