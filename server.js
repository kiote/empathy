
const express = require('express');
const bodyParser = require('body-parser');
const fs = require('fs');
const path = require('path');

const app = express();
const port = 8080;

const file_name = Date.now();

app.use(bodyParser.urlencoded({ extended: false }));
app.use(express.static('.'))

app.post('/demographic', (req, res) => {
  const params = req.body;
  save_demographic(params);
  res.redirect("eq");
})

app.get('/eq', (req, res) => {
  res.sendFile(path.join(__dirname, '/eq.html'));
})

app.listen(port, () => {
  console.log(`Example app listening at http://localhost:${port}`);
})

function save_demographic(params) {
  data = "Name,Age,Gender\r\n"+params.name+','+params.age+','+params.sex+'\r\n';

  fs.writeFile(file_name + '.csv', data, function (err) {
    if (err) throw err;
    console.log('File is created successfully.');
  });
}