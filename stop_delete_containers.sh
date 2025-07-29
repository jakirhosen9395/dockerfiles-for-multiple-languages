#!/bin/bash

# Root directory of the repository
ROOT_DIR=/home/ubuntu/dockerfiles-for-multiple-languages

# Stop and remove .NET Calculator container
echo "Stopping and removing .NET Calculator container..."
docker stop dotnet-calculator
docker rm dotnet-calculator
echo ".NET Calculator container removed."

# Stop and remove Go Calculator container
echo "Stopping and removing Go Calculator container..."
docker stop go-calculator
docker rm go-calculator
echo "Go Calculator container removed."

# Stop and remove Java Calculator container
echo "Stopping and removing Java Calculator container..."
docker stop java-calculator
docker rm java-calculator
echo "Java Calculator container removed."

# Stop and remove Node.js Calculator container
echo "Stopping and removing Node.js Calculator container..."
docker stop node-calculator
docker rm node-calculator
echo "Node.js Calculator container removed."

# Stop and remove Python Calculator container
echo "Stopping and removing Python Calculator container..."
docker stop python-calculator
docker rm python-calculator
echo "Python Calculator container removed."

# Optional: Remove the images if you no longer need them
echo "Removing images..."
docker rmi dotnet-calculator:v1
docker rmi go-calculator:v1
docker rmi java-calculator:v1
docker rmi node-calculator:v1
docker rmi python-calculator:v1
echo "Images removed."

echo "All containers and images (optional) are now stopped and removed."
