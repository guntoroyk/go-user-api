# go-user-api

A RESTFul API build using Go to demonstrate basic CRUD, authentication, and authorization.
Code organized using Clean Architecture and Proxy Design Pattern.

GitHub Repository: https://github.com/guntoroyk/go-user-api

# API Documentation

This API was deployed to Google Kubernetes Engine (GKE) Cluster with public IP `34.142.137.17`.

Click button bellow to try the API through Postman:

[![Run in Postman](https://run.pstmn.io/button.svg)](<https://app.getpostman.com/run-collection/6929584-9e93e6a5-57ca-4250-931d-276052fe8dc0?action=collection%2Ffork&collection-url=entityId%3D6929584-9e93e6a5-57ca-4250-931d-276052fe8dc0%26entityType%3Dcollection%26workspaceId%3D376ac51e-7371-4129-b917-abb587ed642f#?env%5BGo%20User%20API%20(prod)%5D=W3sia2V5IjoidXJsIiwidmFsdWUiOiIzNC4xNDIuMTM3LjE3IiwiZW5hYmxlZCI6dHJ1ZSwidHlwZSI6ImRlZmF1bHQiLCJzZXNzaW9uVmFsdWUiOiIzNC4xNDIuMTM3LjE3Iiwic2Vzc2lvbkluZGV4IjowfSx7ImtleSI6InRva2VuIiwidmFsdWUiOiIiLCJlbmFibGVkIjp0cnVlLCJ0eXBlIjoiZGVmYXVsdCIsInNlc3Npb25WYWx1ZSI6IiIsInNlc3Npb25JbmRleCI6MX1d>)

Or click this URL to open Postman Public Collection https://elements.getpostman.com/redirect?entityId=6929584-9e93e6a5-57ca-4250-931d-276052fe8dc0&entityType=collection

Don't forget to change the Environment located on the upper right corner of Postman app.

If you having trouble accessing that URL, you can download the the postman-collection.json here http://34.142.137.17/docs, then import it to your Postman desktop application.

# Credential

- Role: Admin

```
Username: admin

Password: 1234
```

- Role: User

```
Username: user

Password: 1234
```

## How to run on Local

Clone the repository

```
git clone git@github.com:guntoroyk/go-user-api.git
```

Run on docker

```
docker-compose up --build -d
```

The API will available at http://localhost:8000

## How to run on Kubernetes cluster

```
kubectl apply -f k8s-deploy/app/app_deployment.yml

kubectl apply -f k8s-deploy/app/app_service.yml

kubectl get svc

kubectl get pods
```

## Software Design & Architecture Diagram

This API has 4 domains layer to follow the Clean Architecture:

- Entity layer
- Repository Layer
- Usecase Layer
- Handler Layer

With additional `Middleware` to intercept HTTP request & response, and `Proxy` between the Usecase and Repository to modify input & output data.

![Alt Software Design](./docs/software-design-diagram.png?raw=true "Software Design")

This API was deployed to Google Kubernetes Engine (GKE) cluster with following diagram:

![Alt Architecture Diagram](./docs/architecture-diagram.png?raw=true "Architecture Diagram")

The kubernetes deployment config is located on `./k8s-deploy` folder.

Run this command to deploy using `kubectl`:

```
kubectl apply -f k8s-deploy/app/app_deployment.yml

kubectl apply -f k8s-deploy/app/app_service.yml
```

Here is to check the deployment result:

![Alt Check kubernetes deployment](./docs/kubectl-get-svc-get-pods.png?raw=true "gCheck kubernetes deployment")

## Flowchart

1. Login

![Alt Login](./docs/flow-chart-login.png?raw=true "Login")

2. Create User

![Alt Create User](./docs/flow-chart-create-user.png?raw=true "Create User")

3. Update User

![Alt Update User](./docs/flow-chart-update-user.png?raw=true "Update User")

4. Get Users

![Alt Get Users](./docs/flow-chart-get-user.png?raw=true "Get Users")

5. Get User By ID

![Alt Get User By ID](./docs/flow-chart-get-user-by-id.png?raw=true "Get User By ID")

6. Delete User

![Alt Delete User](./docs/flow-chart-delete-user.png?raw=true "Delete User")
