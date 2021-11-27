const http = require('http');
const fs = require('fs');

const server = http.createServer((req, res) => {

    req.on('error', err => {
        console.error(err);
        // Handle error...
        res.statusCode = 400;
        res.end('400: Bad Request');
        return;
    });

    res.on('error', err => {
        console.error(err);
        // Handle error...
    });

    fs.readFile('./public' + req.url, (err, data) => {
        if (err) {
            if (req.url === '/' && req.method === 'GET') {
                res.end('Welcome Home');
            } else if (req.url === '/tcs' && req.method === 'GET') {
                res.end('HI RCSer');
            } else if (req.url === '/demographic' && req.method === 'POST') {
              res.end('Saved');
            } else {
                res.statusCode = 404;
                res.end('404: File Not Found');
            }
        } else {
            // NOTE: The file name could be parsed to determine the
            // appropriate data type to return. This is just a quick
            // example.
            res.setHeader('Content-Type', 'application/octet-stream');
            res.end(data);
        }
    });

});

server.listen(8080, () => {
    console.log('Example app listening on port 8080!');
});