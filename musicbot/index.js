var express = require('express');
var app = express();

function choice(arr) {
  if (arr.length < 1) {
    throw "arr.length < 1";
  }
  return arr[Math.floor(Math.random() * arr.length)];
}

app.get('/', function(req, res) {
  res.send(choice([
    "Smells Like Teen Spirit is great because it's such a classic.",
    "Bangarang is cool because EDM is hot these days.",
  ]));
});

var server = app.listen(3000, function () {
  var host = server.address().address;
  var port = server.address().port;

  console.log('Example app listening at http://%s:%s', host, port);
});
