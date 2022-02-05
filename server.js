import express from 'express';
import bodyParser from 'body-parser';
import fs from 'fs';
import path  from 'path';
import { checkConnection } from './sensors/gsr.js';

const app = express();
const port = 8080;

const file_name = Date.now();

// randomly returns 1 or 2 to deside which video to start with
function number() {
  return file_name % 2 == 0 ? 1 : 2;
}

app.use(bodyParser.urlencoded({ extended: false }));
app.use(express.static('.'))

/**
 * demographic
 */
app.post('/demographic', (req, res) => {
  const params = req.body;
  checkConnection();
  save_demographic(params);
  res.redirect("eq");
});

/**
 * eq (Empathy Quotient test)
 */
app.get('/eq', (req, res) => {
  res.sendFile(path.join(__dirname, '/eq.html'));
});

app.post('/eq', (req, res) => {
  const params = req.body;
  save_eq(params);
  res.redirect("video");
});

/**
 * video
 */
app.get('/video', (req, res) => {
  res.sendFile(path.join(__dirname, '/video' + number() + '.html'));
});

/**
 * se (situational empathy test)
 */
app.get('/se1', (req, res) => {
  res.sendFile(path.join(__dirname, '/se1.html'));
});

app.get('/se2', (req, res) => {
  res.sendFile(path.join(__dirname, '/se2.html'));
});

app.post('/se1', (req, res) => {
  const params = req.body;
  save_se(1, params);
  if (se_file_exists(2)) {
    res.redirect("done");
  } else {
    res.sendFile(path.join(__dirname, '/video2.html'));
  }
});

app.post('/se2', (req, res) => {
  const params = req.body;
  save_se(2, params);
  if (se_file_exists(1)) {
    res.redirect("done");
  } else {
    res.sendFile(path.join(__dirname, '/video1.html'));
  }
});

/**
 * done
 */
app.get('/done', (req, res) => {
  res.sendFile(path.join(__dirname, '/done.html'));
});

app.listen(port, () => {
  console.log(`App listening at http://localhost:${port}`);
});

function save_demographic(params) {
  if (params.nodrugs == undefined) {
    params.nodrugs = "0";
  }
  console.log(params);
  data = "Race,Age,Gender,Drugs\r\n"+params.race+','+params.age+','+params.sex+','+params.nodrugs+'\r\n';

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

function save_se(num, params) {
  data = 'q1,q2,q3,q4,q5\r\n';
  data += params.q1+','+params.q2+','+params.q3+','+params.q4+','+params.q5 +'\r\n';

  fs.writeFile(file_name + 'se' + num + '.csv', data, function (err) {
    if (err) throw err;
    console.log('File is created successfully.');
  });
}

function se_file_exists(num) {
  let path = './' + file_name + 'se' + num + '.csv';
  console.log(path);
  return fs.existsSync('./' + file_name + 'se' + num + '.csv')
}