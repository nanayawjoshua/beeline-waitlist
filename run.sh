#!/bin/bash
# Beeline Waitlist - Quick Start Script

echo "🚀 Starting Beeline Waitlist Server..."

# Install dependencies
go mod tidy

# Run server
go run server.go