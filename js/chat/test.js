const request = require('request');
const WebSocket = require('ws');


const ws = new WebSocket('ws://localhost:2428');

ws.on('message', (data, flags) => {
  console.log('client received:', data);
});

ws.on('open', () => {
  ws.send('something');
});

function post(json) {
  const options = {
    uri: 'http://localhost:2428/broadcast',
    method: 'POST',
    json: json,
  };

  request(options, (error, response, body) => {
    if (error) {
      console.log('error:', error);
    } else if (response.statusCode != 200) {
      console.log('status code:', response.statusCode);
    } else {
      console.log('success! posted:', json);
    }
  });
}

post({ hello: 'world' });
