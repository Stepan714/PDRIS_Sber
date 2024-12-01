FROM alpine:3.11

RUN apk add --no-cache openssh-client ansible git sshpass

# Скопируйте файлы Ansible, если они находятся на том же уровне
COPY playbook.yml /ansible/playbook.yml
COPY inventories /ansible/inventories
COPY templates /ansible/templates
COPY app /ansible/app

WORKDIR /ansible

CMD [""]

