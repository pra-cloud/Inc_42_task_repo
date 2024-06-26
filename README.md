# DevOps Task for Inc42


This repository contains the code and configurations for a DevOps task aimed at deploying a Go application using Docker. Below are the instructions to get started:

 Docker Hub User Account
The Docker Hub user account associated with this project is pm55.

## Go Application
The Go application is containerized and available as a Docker image on Docker Hub.

 Docker Image
Image Name: pm55/go-app
Tag: latest
Running the Docker Container
To run the Docker container with the Go application, execute the following command:
```bash
docker run -d -p 8080:8080 pm55/go-app:latest
```
This command will start the container in detached mode (-d) and map port 8080 on the host to port 8080 in the container. Replace pm55/go-app:latest with the appropriate image tag if you're using a specific version.

 Accessing the Application
You can access the Go application in your browser using the following URL:

Localhost: http://localhost:8080

## Next Js Application
The Next Js application is containerized and available as a Docker image on Docker Hub.

 Docker Image
Image Name: pm55/next-js-app
Tag: latest
Running the Docker Container
To run the Docker container with the Next.js application, execute the following command:
```bash
docker run -d -p 3000:3000 pm55/next-js-app:latest
```

This command will start the container in detached mode (-d) and map port 3000 on the host to port 3000 in the container. Replace pm55/next-js-app:latest with the appropriate image tag if you're using a specific version.

 Accessing the Application
You can access the Next.js application in your browser using the following URL:

Localhost: http://localhost:3000


## WordPress Application

### NOTE :- WordPress won't run until you provide the MySQL database environment variable to it at the runtime of this Docker container. You can refer to my docker-compose.yml file, which is inside the wordpress-app directory.

The WordPress application is containerized and available as a Docker image on Docker Hub.

Docker Image
Image Name: pm55/wordpress-app
Tag: latest
Running the Docker Container
To run the Docker container with the WordPress application, execute the following command:
```bash
docker run -d -p 8000:80 pm55/wordpress-app:latest
```

This command will start the container in detached mode (-d) and map port 8000 on the host to port 80 in the container. Replace pm55/wordpress-app:latest with the appropriate image tag if you're using a specific version.

Accessing the Application
You can access the WordPress application in your browser using the following URL:

Localhost: http://localhost:8000


### AWS CI/CD Pipeline
For achieving continuous integration and continuous deployment (CI/CD), we've implemented an AWS CI/CD pipeline. AWS CodeBuild handles linting and code optimization for specific technologies using the aws-buildspec-cicd.yml file, while AWS CodeDeploy deploys the application inside an EC2 server with zero downtime using appspec.yml.

**AWS CodeBuild Setup**

AWS CodeBuild is configured to handle the following tasks:
Linting and Code Optimization: AWS CodeBuild utilizes the aws-buildspec-cicd.yml file to perform linting and code optimization for the respective applications.

**AWS CodeDeploy Setup**
AWS CodeDeploy manages the deployment process onto an EC2 server with zero downtime. It uses the appspec.yml file to orchestrate the deployment process, ensuring smooth and seamless updates.

# Docker Compose
To build all the applications simultaneously, you can utilize the docker-compose.yaml file located in the docker-compose directory.

**Building Applications**
Execute the following command to build the applications using the docker-compose.yaml file:
```bash
docker compose up -d
```

This command will build and start all the services defined in the docker-compose.yaml file in detached mode (-d), allowing you to continue working in your terminal.

**Stopping and Removing Builds**
To stop or remove all the builds created using the docker-compose.yaml file, execute the following command:
```bash
docker compose down -d
```

This command will stop and remove all the containers, networks, and volumes created by the docker-compose.yaml file in detached mode (-d), ensuring clean teardown of the environment.
