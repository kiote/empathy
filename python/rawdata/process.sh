#!/bin/bash
# This script was tested under MacOS Big Sur and Arch Linux

function split_sensor_files() {
    # $1 - episode nr
    sed -n "1, 1p" gsr_pulse$1.csv > gsr$1.csv
    sed -n "2, 2p" gsr_pulse$1.csv > pulse$1.csv

    # change commas to new lines
    tr , '\n' < gsr$1.csv > gsrLN$1.csv
    tr , '\n' < pulse$1.csv > pulseLN$1.csv

    # remove extra files
    rm gsr$1.csv pulse$1.csv
}

function add_timestamp() {
    # $1 - input file name
    # $2 - ouput file name
    LINES=$(wc -l < $1)

    # length of one measurement in ms
    ONE_MEASUREMENT=$(( $EPISODE_LENGTH / $LINES ))

    # starting time
    CURRENT_TIME=0
    cat $1 | while read line 
    do
        echo "$line,$CURRENT_TIME"
        CURRENT_TIME=$(( $CURRENT_TIME + $ONE_MEASUREMENT ))
    done > $2
}

for folder in $(ls -d */); do
    cd ${folder}
    
    split_sensor_files "1"
    split_sensor_files "2"

    # 18 sec in ms
    EPISODE_LENGTH=18000

    add_timestamp "gsrLN1.csv" "gsrTS1.csv" 
    add_timestamp "gsrLN2.csv" "gsrTS2.csv" 

    add_timestamp "pulseLN1.csv" "pulseTS1.csv"
    add_timestamp "pulseLN2.csv" "pulseTS2.csv"

    rm "gsrLN1.csv" "gsrLN2.csv" "pulseLN1.csv" "pulseLN2.csv"

    cd ..
done