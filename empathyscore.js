const fs = require('fs');
const path = require('path');
const csv = require('fast-csv');

const positive = ['q1','q3','q11','q13','q14','q15','q21','q22','q23','q24','q26','q27','q28','q29','q34','q35','q36','q37','q38','q39','q40'];
const negative = ['q2','q4','q5','q6','q7','q8','q9','q10','q12','q16','q17','q18','q19','q20','q25','q30','q31','q32','q33'];
let score = 0;
fs.createReadStream(path.resolve(__dirname, 'parse.csv'))
    .pipe(csv.parse({ headers: true }))
    .on('error', error => console.error(error))
    .on('data', row => {
        
        positive.forEach(q => {
            if (row[q] == 1) {
                score += 2;
            }
            if (row[q] == 2) {
                score += 1;
            }
        });

        negative.forEach(q => {
            if (row[q] == 4) {
                score += 2;
            }
            if (row[q] == 3) {
                score += 1;
            }
        });
    })
    .on('end', rowCount => {
        console.log(`Parsed ${rowCount} rows`)
        console.log("score: " + score);
    });

    