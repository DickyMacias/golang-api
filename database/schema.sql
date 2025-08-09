-- Movie Tracker Database Schema
-- PostgreSQL

-- Users table
CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(50) UNIQUE NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    password_hash VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP
);

-- Favorite movies table
CREATE TABLE IF NOT EXISTS favorite_movies (
    id SERIAL PRIMARY KEY,
    user_id INTEGER REFERENCES users(id) ON DELETE CASCADE,
    tmdb_id INTEGER NOT NULL,
    title VARCHAR(255) NOT NULL,
    overview TEXT,
    release_date DATE,
    poster_path VARCHAR(255),
    genre_ids JSONB,
    status VARCHAR(20) DEFAULT 'por_ver' CHECK (status IN ('por_ver', 'vista', 'recomendada')),
    rating INTEGER CHECK (rating >= 1 AND rating <= 10),
    notes TEXT,
    recommended_by VARCHAR(100),
    added_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    watched_at TIMESTAMP,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP,
    UNIQUE(user_id, tmdb_id)
);

-- Indexes for better performance
CREATE INDEX IF NOT EXISTS idx_favorite_movies_user_id ON favorite_movies(user_id);
CREATE INDEX IF NOT EXISTS idx_favorite_movies_status ON favorite_movies(user_id, status);
CREATE INDEX IF NOT EXISTS idx_favorite_movies_added_at ON favorite_movies(user_id, added_at DESC);
CREATE INDEX IF NOT EXISTS idx_users_username ON users(username);
CREATE INDEX IF NOT EXISTS idx_users_email ON users(email);

-- Comments for documentation
COMMENT ON TABLE users IS 'Application users with authentication credentials';
COMMENT ON TABLE favorite_movies IS 'User favorite movies with personal metadata';
COMMENT ON COLUMN favorite_movies.status IS 'Movie status: por_ver (to watch), vista (watched), recomendada (recommended)';
COMMENT ON COLUMN favorite_movies.rating IS 'Personal rating from 1-10 stars';
COMMENT ON COLUMN favorite_movies.tmdb_id IS 'The Movie Database ID for external API reference';