
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

app.post('/eq', (req, res) => {
  const params = req.body;
  save_eq(params);
  res.redirect("eq");
})

app.listen(port, () => {
  console.log(`Example app listening at http://localhost:${port}`);
})

function save_demographic(params) {
  data = "Name,Age,Gender\r\n"+params.name+','+params.age+','+params.sex+'\r\n';

  fs.writeFile(file_name + 'dm.csv', data, function (err) {
    if (err) throw err;
    console.log('File is created successfully.');
  });
}

function save_eq(params) {
  data = 'q1,q2,q3,q4,q5,q6,q7,q8,q9,q10,q11,q12,q13,q14,q15,q16,q17,q18,q19,q20,';
  data += 'q21,q22,q23,q24,q25,q26,q27,q28,q29,q30,q31,q32,q33,q34,q35,q36,q37,q38,q39,q40\r\n';
  data += params.q1+','+params.q2+','+params.q3+','+params.q4+','+params.q5+','+params.q6 +','+ params.q7+','+params.q8+','+params.q9 + ',';
  data += params.q10 +','+ params.q11+','+params.q12+','+params.q13+','+params.q14+','+params.q15+','+params.q16 +','+ params.q17+',';
  data += params.q18 +','+params.q19 + ','+params.q20 +','+ params.q21+','+params.q22+','+params.q23+','+params.q24+','+params.q25+',';
  data += params.q26 +','+ params.q27+','+params.q28+','+params.q29 + ','+params.q30 + ','+params.q31+','+params.q32+','+params.q33+',';
  data += params.q34+','+params.q35+','+params.q36 +','+ params.q37+','+params.q38+','+params.q39 + ','+params.q40 +'\r\n';

  fs.writeFile(file_name + 'qe.csv', data, function (err) {
    if (err) throw err;
    console.log('File is created successfully.');
  });
}