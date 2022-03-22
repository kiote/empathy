#!/bin/bash
function add_timestanp() {
    # $1 - input file name
    # $2 - ouput file name
    LINES=$(wc -l < $1)

    # length of one measurement in ms
    ONE_MEASUREMENT=$(( $EPISODE_LENGTH / $LINES ))

    # starting time
    CURRENT_TIME=0
    cat $1 | while read line 
    do
        echo "$line;$CURRENT_TIME"
        CURRENT_TIME=$(( $CURRENT_TIME + $ONE_MEASUREMENT ))
    done > $2
}

for folder in $(ls -d */); do 
    cd ${folder}
    # # first gsr_pulse split
    # sed -n "1, 1p" gsr_pulse2.csv > gsr2.csv
    # sed -n "2, 2p" gsr_pulse2.csv > pulse2.csv

    # # second gsr_pulse split
    # sed -n "1, 1p" gsr_pulse1.csv > gsr1.csv
    # sed -n "2, 2p" gsr_pulse1.csv > pulse1.csv

    # # change commas to new lines
    # tr , '\n' < gsr1.csv > gsrLN1.csv
    # tr , '\n' < gsr2.csv > gsrLN2.csv

    # tr , '\n' < pulse1.csv > pulseLN1.csv
    # tr , '\n' < pulse2.csv > pulseLN2.csv

    # # remove extra files
    # rm gsr_pulse1.csv gsr_pulse2.csv gsr1.csv gsr2.csv pulse1.csv pulse2.csv

    # 18 sec in ms
    EPISODE_LENGTH=18000

    add_timestanp "gsrLN1.csv" "gsrTS1.csv" 
    add_timestanp "gsrLN2.csv" "gsrTS2.csv" 
done