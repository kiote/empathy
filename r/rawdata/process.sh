#!/bin/bash
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

    # remove extra files
    rm gsr_pulse1.csv gsr_pulse2.csv gsr1.csv gsr2.csv pulse1.csv pulse2.csv

    # 18 sec
    cat gsrLN1.csv | while read line 
    do
        echo $line;
    done
done