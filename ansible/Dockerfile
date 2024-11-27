FROM python:3.9-alpine

WORKDIR /app

COPY . /app

RUN apk add --no-cache openssh-client git sshpass && \
    pip install ansible docker

CMD ["sleep", "infinity"]
