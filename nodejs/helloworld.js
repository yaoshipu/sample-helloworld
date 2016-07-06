var http = require('http');

var server = http.createServer(function (request, response) {
  console.log("Handle " + request.method + " " + request.url);
  response.writeHead(200, {"Content-Type": "text/plain"});
  response.end("Hello world from my Node program!\n");
});

server.listen(8000);

console.log("Server running at http://127.0.0.1:8000/");