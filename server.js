
const express = require('express');
const app = express();
const port = 8080;

const bodyParser = require('body-parser');

app.use(bodyParser.json());

app.post('/demographic', (req, res) => {
  console.log(req.body);
  res.send('saved');
})

app.listen(port, () => {
  console.log(`Example app listening at http://localhost:${port}`);
})