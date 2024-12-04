# PSID_Sber

Выполните эти команды:
`git clone https://github.com/Stepan714/PDRIS_Sber.git`

`docker login`

`docker build -t your-dockerhub-username/flask-app .
docker push your-dockerhub-username/flask-app`

Изменить файл: k8s/flask-deployment.yaml, поставить свой your-dockerhub-username

`minikube start`

`kubectl apply -f k8s/redis-deployment.yaml`

`kubectl apply -f k8s/redis-service.yaml`

`kubectl apply -f k8s/flask-deployment.yaml`

`kubectl apply -f k8s/flask-service.yaml`
