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

// TODO: test out that hello worlds 1 and 2 can happen

app.use(bodyParser.json());

app.get('/', function (req, res) {
  res.send('this is the chat server')
});

app.get('/broadcast', function (req, res) {
  wss.broadcast(req.body);
  res.send('OK');
});

wss.broadcast = (data) => {
  wss.clients.forEach((client) => {
    client.send(data);
  });
};

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
