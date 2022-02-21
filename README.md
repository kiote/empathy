# Experiment runner for master thesis

This repo is made to run experiment for my thesis. It controls NeuLog sensors (GSR, Pulse) and iMotions program.

## To start the experiment

1. Download zip of this repo.
2. Run go/m.exe (only Windows supported)
3. Open http://localhost:8090 in incognito window

## To calculate empathy score

Copy eq.csv file from corresponding folder to project root, rename it to `parse.csv` and run 

```
npm i
npm start
```

## What is happening here

```mermaid
sequenceDiagram
    Go script ->> Web page: Start experiment page
    Web page ->> Go script: EQ test results
    loop 2 times
        Go script ->> Web page: Show randomly chosen video1
        Go script ->> Sensors: Start experiment
        Go script ->> iMotion: Set marker start
        Web page ->> Web page: regirect to short test
        Go script ->> iMotion: Set marker stop
        Go script ->> Sensors: Stop experiment
        Sensors ->> Go script: Sensors results
        Web page ->> Go script: Short test results    
    end 
```
