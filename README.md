
# Dockerfiles for Multiple Languages

This repository contains Dockerfiles for building and running basic calculator applications in multiple programming languages. The goal of this project is to demonstrate how to containerize applications developed in different languages using Docker.

The project includes the following languages:
- .NET (C#)
- Go
- Java
- Node.js
- Python

## Output
<p align="center">
  <img src="https://raw.githubusercontent.com/jakirhosen9395/dockerfiles-for-multiple-languages/main/validate/dotnet.png" width="500" />
  <img src="https://raw.githubusercontent.com/jakirhosen9395/dockerfiles-for-multiple-languages/main/validate/golang.png" width="500" />
  <img src="https://raw.githubusercontent.com/jakirhosen9395/dockerfiles-for-multiple-languages/main/validate/java.png" width="500" />
  <img src="https://raw.githubusercontent.com/jakirhosen9395/dockerfiles-for-multiple-languages/main/validate/nodejs.png" width="500" />
  <img src="https://raw.githubusercontent.com/jakirhosen9395/dockerfiles-for-multiple-languages/main/validate/python.png" width="500" />
</p>

## Prerequisites

Before running the Docker containers for these applications, you must install Docker on your system. Follow the steps below to install Docker on Ubuntu.

### Step 1: Install Docker on Ubuntu

1. Update your system's package list:
   ```bash
   sudo apt update
   ```

2. Upgrade your installed packages:
   ```bash
   sudo apt-get upgrade -y
   ```

3. Download and install Docker:
   ```bash
   curl -fsSL https://test.docker.com -o test-docker.sh
   sudo sh test-docker.sh
   ```

4. Add the `docker` group and add your user to this group to avoid using `sudo` with Docker commands:
   ```bash
   sudo groupadd docker
   sudo usermod -aG docker ubuntu
   newgrp docker
   ```

5. Enable Docker services to start on boot:
   ```bash
   sudo systemctl enable docker.service
   sudo systemctl enable containerd.service
   ```

After performing the above steps, Docker will be installed and ready for use.

## Bash Script to Run All Containers Automatically

In addition to the Dockerfiles, a bash script (`docker_run_all.sh`) has been included in the repository to automatically build and run all the Docker containers for the calculator applications. This script will run all the containers on different ports without requiring manual intervention.

### Steps to Run the Script

1. First, ensure the script is executable:
   ```bash
   chmod +x run-all.sh
   ```

2. Run the script:
   ```bash
   ./run-all.sh
   ```

The script will automatically:
- Build Docker images for each language (C#, Go, Java, Node.js, Python).
- Run each container on different ports to avoid conflicts:
  - .NET Calculator on port 8081
  - Go Calculator on port 8082
  - Java Calculator on port 8083
  - Node.js Calculator on port 8084
  - Python Calculator on port 8085

This approach simplifies the process of running all containers at once.

## Bash Script to Stop and Remove All Containers

A bash script (`stop_delete_containers.sh`) has been included in the repository to stop and remove all the containers created by `run-all.sh`. This script can also remove the Docker images associated with the containers, freeing up space.

### Steps to Run the Stop/Delete Script

1. Ensure the script is executable:
   ```bash
   chmod +x stop_delete_containers.sh
   ```

2. Run the script to stop and remove all containers:
   ```bash
   ./stop_delete_containers.sh
   ```

This script will:
- Stop each running container.
- Remove the stopped containers.
- Optionally, remove the images created for the containers.

## Language-Specific Dockerfiles

Each directory contains a Dockerfile for building a containerized calculator application in the respective language.

### Dotnet Calculator

1. Navigate to the `dotnet-calculator` directory:
   ```bash
   cd dockerfiles-for-multiple-languages/dotnet-calculator/
   ```

2. Build the Docker image:
   ```bash
   docker build -t dotnet-calculator:v1 .
   ```

3. Run the container:
   ```bash
   docker run -d --name dotnet-calculator -p 8081:80 dotnet-calculator:v1
   ```

4. Check if the container is running:
   ```bash
   docker ps
   ```

5. To stop the container:
   ```bash
   docker stop <container_id>
   ```

6. To remove the container:
   ```bash
   docker rm -f <container_id>
   ```

### Golang Calculator

1. Navigate to the `go-calculator` directory:
   ```bash
   cd ../go-calculator/
   ```

2. Build the Docker image:
   ```bash
   docker build -t go-calculator .
   ```

3. Run the container:
   ```bash
   docker run -d --name go-calculator -p 8082:8080 go-calculator:v1
   ```

4. To stop and remove the container:
   ```bash
   docker stop <container_id>
   docker rm -f <container_id>
   ```

### Java Calculator

1. Navigate to the `java-calculator` directory:
   ```bash
   cd ../java-calculator/
   ```

2. Build the Docker image:
   ```bash
   docker build -t java-calculator:v1 .
   ```

3. Run the container:
   ```bash
   docker run -d --name java-calculator -p 8083:8080 java-calculator:v1
   ```

4. To stop and remove the container:
   ```bash
   docker stop java-calculator
   docker rm -f java-calculator
   ```

### Node.js Calculator

1. Navigate to the `node-calculator` directory:
   ```bash
   cd ../node-calculator/
   ```

2. Build the Docker image:
   ```bash
   docker build -t node-calculator:v1 .
   ```

3. Run the container:
   ```bash
   docker run -d --name node-calculator -p 8084:8080 node-calculator:v1
   ```

4. To stop and remove the container:
   ```bash
   docker stop node-calculator
   docker rm -f <container_id>
   ```

### Python Calculator

1. Navigate to the `python-calculator` directory:
   ```bash
   cd ../python-calculator/
   ```

2. Build the Docker image:
   ```bash
   docker build -t python-calculator:v1 .
   ```

3. Run the container:
   ```bash
   docker run -d --name python-calculator -p 8085:8080 python-calculator:v1
   ```

4. To stop and remove the container:
   ```bash
   docker stop <container_id>
   docker rm -f <container_id>
   ```

## Conclusion

This repository demonstrates how to containerize simple calculator applications written in different programming languages. You can extend this project by adding more features or adapting the code for other languages.

Feel free to fork or contribute to the repository as you like. Happy coding!