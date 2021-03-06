// Using both express and ws to handle regular web stuff and websockets.
// Based on https://www.npmjs.com/package/ws

const http = require('http');
const WebSocket = require('ws');
const urllib = require('url');
const express = require('express');
const bodyParser = require('body-parser');
const cors = require('cors');

const server = http.createServer();
const wss = new WebSocket.Server({ server: server });
const app = express();

const PORT = 2428;

const MESSAGES = [{
  id: '1234',
  content: 'Hello world',
  timestamp: 0,
}];

app.use(bodyParser.json());
app.use(cors());

app.get('/', function (req, res) {
  res.send('this is the chat server')
});

app.get('/messages', (req, res) => {
  console.log('handling GET /messages');
  res.send(JSON.stringify(MESSAGES));
});

function makeID() {
  let answer = '';
  for (let i = 0; i < 8; i++) {
    let index = Math.floor(Math.random() * 32);
    answer += 'ABCDEFGHJKMNPQRSTVWXYZ0123456789'[index];
  }
  return answer;
}


app.post('/messages', (req, res) => {
  console.log('handling POST /messages');
  if (!req.body.id) {
    req.body.id = makeID();
  }
  if (!req.body.timestamp) {
    req.body.timestamp = (new Date()).getTime();
  }
  MESSAGES.push(req.body);
  wss.broadcast(req.body);
  res.send('OK');
})

// Just for testing
app.post('/broadcast', (req, res) => {
  console.log('handling POST /broadcast');
  wss.broadcast(req.body);
  res.send('OK');
});

// Takes something JSON-encodable
wss.broadcast = (json) => {
  let payload = JSON.stringify(json);
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
  console.log('Running chat server on port ' + server.address().port);
});
