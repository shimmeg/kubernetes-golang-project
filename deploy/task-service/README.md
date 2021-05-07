1. How to execute MongoDB database:
   helm repo add bitnami https://charts.bitnami.com/bitnami
   helm install mongo bitnami/mongodb -f mongodb/mongo_values.yaml

2. Kubernetes Job to init mongo database:
   kubectl apply -f initdb.yaml

3. Executing the application itself with helm chart:
   helm install task-app ./task-service-chart 