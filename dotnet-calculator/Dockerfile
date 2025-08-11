#-----------------------------Multistage Dockerfile for .NET Calculator App-----------------------------#

# Build Stage
FROM mcr.microsoft.com/dotnet/sdk:9.0 AS build
WORKDIR /app

COPY *.csproj ./
RUN dotnet restore

COPY . ./
RUN dotnet publish -c Release -o out

# Runtime Stage
FROM mcr.microsoft.com/dotnet/aspnet:9.0 AS runtime
WORKDIR /app

COPY --from=build /app/out ./

# Port expose
EXPOSE 8080

# Run the app
ENTRYPOINT ["dotnet", "dotnet-calculator.dll"]





#-----------------------------Single-Stage Dockerfile for .NET Calculator App-----------------------------#
# # Use the .NET 9 SDK image to build and run the app
# FROM mcr.microsoft.com/dotnet/sdk:9.0

# # Set the working directory
# WORKDIR /dotnet-calculator

# # Copy the entire project into the container
# COPY . .

# # Restore the project dependencies
# RUN dotnet restore

# # Build the application in Release mode and publish to /dotnet-calculator/publish
# RUN dotnet publish -c Release -o /dotnet-calculator/publish

# # Set the entry point to run the app
# CMD ["dotnet", "/dotnet-calculator/publish/dotnet-calculator.dll"]

