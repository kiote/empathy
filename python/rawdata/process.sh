#!/bin/bash
# This script was tested under MacOS Big Sur and Arch Linux

function split_files() {
    # $1 - episode nr
    sed -n "1, 1p" gsr_pulse$1.csv > gsr$1.csv
    sed -n "2, 2p" gsr_pulse$1.csv > pulse$1.csv

    # change commas to new lines
    tr , '\n' < gsr$1.csv > gsrLN$1.csv
    tr , '\n' < pulse$1.csv > pulseLN$1.csv

    # remove extra files
    rm gsr$1.csv pulse$1.csv
}

for folder in $(ls -d */); do 
    cd ${folder}
    
    split_files "1"
    split_files "2"
done