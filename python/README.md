## Run with Anaconda

1. You'll need Docker to be installed in your system
3. `docker pull continuumio/anaconda3`
4. Run container:
(remember to put your path instead of `/home/ekaterina/Sandbox/`)
```
docker run -i -t -p 8888:8888 -v /home/ekaterina/Sandbox/empathy/python/:/opt/notebooks/ continuumio/anaconda3 /bin/bash -c "\
    conda install jupyter -y --quiet && \
    mkdir -p /opt/notebooks && \
    jupyter notebook \
    --notebook-dir=/opt/notebooks --ip='*' --port=8888 \
    --no-browser --allow-root"
```
4. Connect with provided adderes (http://127.0.0.1:8888/?token=\<should be provided in output\>)

## Raw data

Copy folders with data to `rawdata` folder (there is already one folder for example)

## Results

Results are presented in file https://docs.google.com/spreadsheets/d/1dNKvPI8Yzw7o3Z_3Ux3id1OIN9FHpvlIxgIzwnYc6_Y/edit#gid=2015187546

Stimuli were presented to participants in randomized order. Even NR of participant: video1, video2. Odd NR of participant: video2, video1

## Useful commands

Find line number with particular text
```
awk '/Stimuli\+started/ {print FNR}' emotions.csv
```

Write to file starting from line N of another
```
tail -n +15116 emotions.csv > emotions1.csv
```
