# Use the .NET 9 SDK image to build and run the app
FROM mcr.microsoft.com/dotnet/sdk:9.0

# Set the working directory
WORKDIR /dotnet-calculator

# Copy the entire project into the container
COPY . .

# Restore the project dependencies
RUN dotnet restore

# Build the application in Release mode and publish to /dotnet-calculator/publish
RUN dotnet publish -c Release -o /dotnet-calculator/publish

# Set the entry point to run the app
CMD ["dotnet", "/dotnet-calculator/publish/dotnet-calculator.dll"]
