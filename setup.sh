#!/bin/bash

# Movie Tracker Setup Script
echo "🎬 Movie Tracker Setup"
echo "====================="

# Check if Go is installed
if ! command -v go &> /dev/null; then
    echo "❌ Go is not installed. Please install Go 1.21 or later."
    exit 1
fi

# Check if PostgreSQL is running
if ! command -v psql &> /dev/null; then
    echo "❌ PostgreSQL is not installed. Please install PostgreSQL."
    exit 1
fi

echo "✅ Go and PostgreSQL found"

# Install dependencies
echo "📦 Installing dependencies..."
go mod tidy

# Check if .env exists
if [ ! -f ".env" ]; then
    echo "⚠️  .env file not found. Please create one with the following variables:"
    echo "DATABASE_URL=postgres://username:password@localhost:5432/movietracker?sslmode=disable"
    echo "TMDB_API_KEY=your_tmdb_api_key_here"
    echo "SESSION_SECRET=your_very_secure_session_secret_here"
    echo "PORT=8080"
    echo "ENVIRONMENT=development"
    echo ""
    echo "Get your TMDB API key from: https://www.themoviedb.org/settings/api"
    exit 1
fi

echo "✅ Environment variables configured"

# Ask user if they want to create the database
read -p "Do you want to create the PostgreSQL database? (y/n): " -n 1 -r
echo
if [[ $REPLY =~ ^[Yy]$ ]]; then
    echo "📊 Creating database..."
    # Extract database info from .env
    DB_URL=$(grep DATABASE_URL .env | cut -d '=' -f2)
    if [[ $DB_URL == *"movietracker"* ]]; then
        createdb movietracker 2>/dev/null || echo "Database might already exist"
        echo "✅ Database setup complete"
    else
        echo "⚠️  Please ensure your DATABASE_URL includes the database name"
    fi
fi

echo ""
echo "🚀 Setup complete! You can now run:"
echo "   go run main.go"
echo ""
echo "Then visit: http://localhost:8080"
echo ""
echo "Happy movie tracking! 🎬✨"