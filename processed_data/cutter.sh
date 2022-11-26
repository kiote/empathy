function remove_columns() {
    # $1 - input file name
    cut --complement -f 2,3,4,5,6,7,8,9 -d, $1.csv > $1_processed.csv
    cut -f1-11 -d, $1_processed.csv > $1_cutted.csv
    rm $1_processed.csv
}

for folder in $(ls -d */); do
    cd ${folder}

    remove_columns "A"
    remove_columns "B"

    cd ..
done
