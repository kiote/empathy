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

    # rename
    cp *Anonymous* emotions.csv

    # get stimuli 1
    STIMULI1_START=$(awk '/Stimuli\+started/ {print FNR}' emotions.csv | head -n 1)
    STIMULI1_START=$(( $STIMULI1_START + 1 ))
    STIMULI1_FINISH=$(awk '/Stimuli\+finished/ {print FNR}' emotions.csv | head -n 1)
    STIMULI1_FINISH=$(( $STIMULI1_FINISH - 1 ))

    sed -n "${STIMULI1_START},${STIMULI1_FINISH}p" emotions.csv > emotions1.csv 

    # awk -i inplace '{$47=$48=$49=$50=$51=$52=$53=$54=$55=$56=$57=$58=$59=$60=$61=$62=$63=$64=$65=$66=$67=$68=$69=$70=$71=$72=$73=$74=$75=$76=$77=$78=$79=$80=$81=$82=$83=$84=$85=$86=$87=$88=$89=$90=$91=$92=$93=$94=$95=$96=$97=$98=$99=$100=$101=$102=$103=$104=$105=$106=$107=$108=$109=$110=$111=$112=$113=$114=$115=$116=$117=$118=$119=$120=$121=$122=$123=$124=$125=$126=$127=$128=$129=$130=$131=$132=$133=$134=$135=$136=$137=$138=$139=$140=$141=$142=$143=$144=$145=$146=$147=$148=$149=$150=$151=$152=$153=$154=$155=$156=$157=$158=$159=""; print $0}' emotions1.csv

    # get stimuli 2
    STIMULI2_START=$(awk '/Stimuli\+started/ {print FNR}' emotions.csv | tail -n 1)
    STIMULI2_START=$(( $STIMULI2_START + 1 ))
    STIMULI2_FINISH=$(awk '/Stimuli\+finished/ {print FNR}' emotions.csv | tail -n 1)
    STIMULI2_FINISH=$(( $STIMULI2_FINISH - 1 ))

    sed -n "${STIMULI2_START},${STIMULI2_FINISH}p" emotions.csv > emotions2.csv

    cd ..
done