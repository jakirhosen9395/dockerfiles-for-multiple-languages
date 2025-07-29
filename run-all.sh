ll.sh
#!/bin/bash

# Root directory of the repository
ROOT_DIR=/home/ubuntu/dockerfiles-for-multiple-languages

# Run .NET Calculator
echo "Starting .NET Calculator..."
cd $ROOT_DIR/dotnet-calculator/
docker build -t dotnet-calculator:v1 .
docker run -d --name dotnet-calculator -p 8081:8080 dotnet-calculator:v1
echo ".NET Calculator running on port 8081."

# Run Go Calculator
echo "Starting Go Calculator..."
cd $ROOT_DIR/go-calculator/
docker build -t go-calculator:v1 .
docker run -d --name go-calculator -p 8082:8080 go-calculator:v1
echo "Go Calculator running on port 8082."

# Run Java Calculator
echo "Starting Java Calculator..."
cd $ROOT_DIR/java-calculator/
docker build -t java-calculator:v1 .
docker run -d --name java-calculator -p 8083:8080 java-calculator:v1
echo "Java Calculator running on port 8083."

# Run Node.js Calculator
echo "Starting Node.js Calculator..."
cd $ROOT_DIR/node-calculator/
docker build -t node-calculator:v1 .
docker run -d --name node-calculator -p 8084:8080 node-calculator:v1
echo "Node.js Calculator running on port 8084."

# Run Python Calculator
echo "Starting Python Calculator..."
cd $ROOT_DIR/python-calculator/
docker build -t python-calculator:v1 .
docker run -d --name python-calculator -p 8085:8080 python-calculator:v1
echo "Python Calculator running on port 8085."

echo "All containers are now running!"
