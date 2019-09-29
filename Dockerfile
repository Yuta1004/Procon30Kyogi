FROM alpine:latest

COPY solver.py /tmp/

RUN     echo "[Info] Start Build..." && \
        set -x && \
        echo "[Info] Install Packages..." && \
        apk --update --no-cache add \
                python3 \
                python3-dev \
                gcc \
                g++ \
                make \
                musl \
                musl-dev \
                linux-headers \
                gfortran \
                openblas-dev && \
        echo "[Info] Insall Python libaries..." && \
        pip install --upgrade pip && \
        pip3 install numpy && \
        echo "[INFO] Setup solver.py" && \
        cd /tmp && \
        echo "#!/usr/bin/env python3" > test.py && \
        echo "import numpy" >> test.py && \
        chmod +x test.py && \
        chmod +x solver.py && \
        ./test.py
