FROM alpine:3.11

RUN apk add --no-cache openssh-client ansible git sshpass python3 py3-pip && \
    pip3 install docker

WORKDIR /ansible

COPY . .

CMD ["ansible-playbook", "DockerPlaybook.yml", "-i", "inventory.hosts"]
