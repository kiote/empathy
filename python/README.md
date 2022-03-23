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