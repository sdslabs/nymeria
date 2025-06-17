#!/bin/bash

# Development environment runner script
echo "🚀 Starting development environment..."

# Build and run with docker-compose dev configuration
docker-compose -f docker-compose.dev.yaml up --build --force-recreate

# Cleanup function
cleanup() {
    echo "🧹 Cleaning up..."
    docker-compose -f docker-compose.dev.yaml down
}

# Set trap to cleanup on script exit
trap cleanup EXIT 