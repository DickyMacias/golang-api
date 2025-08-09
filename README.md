# 🎬 Movie Tracker App

A full-featured web application built with Go, Gin Framework, HTMX, and PostgreSQL for managing your personal movie collection with TMDB API integration.

## ✨ Features

- **User Authentication**: Secure registration and login with session-based authentication
- **Movie Search**: Search movies using The Movie Database (TMDB) API
- **Personal Collection**: Add movies to your favorites with different statuses
- **Status Management**: Track movies as "To Watch", "Watched", or "Recommended"
- **Rating System**: Rate movies from 1-10 stars
- **Personal Notes**: Add notes and track who recommended each movie
- **Interactive UI**: Real-time updates using HTMX without page reloads
- **Responsive Design**: Mobile-friendly interface using Tailwind CSS

## 🛠 Tech Stack

- **Backend**: Go + Gin Framework
- **Frontend**: HTMX + Go Templates + Tailwind CSS
- **Database**: PostgreSQL + GORM
- **Authentication**: Session-based with secure cookies
- **External API**: The Movie Database (TMDB) API

## 📋 Prerequisites

Before you begin, ensure you have the following installed:

- [Go](https://golang.org/download/) (1.21 or later)
- [PostgreSQL](https://www.postgresql.org/download/) (12 or later)
- A [TMDB API Key](https://www.themoviedb.org/settings/api)

## 🚀 Quick Start

### 1. Clone the Repository

```bash
git clone <your-repo-url>
cd movie-tracker
```

### 2. Set Up Environment Variables

Create a `.env` file in the project root:

```env
DATABASE_URL=postgres://username:password@localhost:5432/movietracker?sslmode=disable
TMDB_API_KEY=your_tmdb_api_key_here
SESSION_SECRET=your_very_secure_session_secret_here
PORT=8080
ENVIRONMENT=development
```

**Important**: 
- Replace `username`, `password` with your PostgreSQL credentials
- Get your TMDB API key from [TMDB Settings](https://www.themoviedb.org/settings/api)
- Use a strong, random session secret (at least 32 characters)

### 3. Set Up Database

Create the PostgreSQL database:

```sql
CREATE DATABASE movietracker;
```

The application will automatically create the required tables on startup.

### 4. Install Dependencies

```bash
go mod tidy
```

### 5. Run the Application

```bash
go run main.go
```

The application will be available at `http://localhost:8080`

## 📁 Project Structure

```
movie-tracker/
├── main.go                 # Application entry point
├── config/
│   └── config.go          # Configuration management
├── models/
│   ├── user.go           # User model
│   ├── movie.go          # TMDB movie models
│   └── favorite.go       # Favorite movie model
├── handlers/
│   ├── auth_handler.go   # Authentication handlers
│   ├── tmdb_handler.go   # TMDB API handlers
│   └── favorites_handler.go # Favorites CRUD handlers
├── services/
│   ├── auth_service.go   # Authentication business logic
│   ├── tmdb_service.go   # TMDB API integration
│   └── favorites_service.go # Favorites business logic
├── middleware/
│   ├── auth_middleware.go   # Authentication middleware
│   └── session_middleware.go # Session management
├── database/
│   └── connection.go     # Database connection setup
├── routes/
│   └── routes.go         # Route definitions
├── templates/           # HTML templates
├── static/             # Static assets (CSS, JS)
└── .env               # Environment variables
```

## 🎯 Usage

### Registration & Login
1. Visit `http://localhost:8080`
2. Register a new account with username, email, and password
3. Login with your credentials

### Adding Movies
1. Navigate to the **Search** page
2. Search for movies using the search bar
3. Click "Add to Favorites" on any movie
4. Set the status (To Watch, Watched, Recommended)
5. Optionally add rating, notes, and who recommended it

### Managing Your Collection
1. Visit the **Favorites** page to see all your movies
2. Use the filter tabs to view movies by status
3. Update movie status using the dropdown
4. Rate movies by clicking on stars (1-10)
5. Remove movies by clicking the delete button

### Dashboard
- View statistics of your movie collection
- Quick access to main features
- See popular movies for discovery

## 🔧 API Endpoints

### Public Routes
- `GET /` - Redirect to login
- `GET /login` - Login page
- `POST /login` - Process login
- `GET /register` - Registration page
- `POST /register` - Process registration

### Protected Routes
- `GET /dashboard` - User dashboard
- `GET /search` - Movie search page
- `GET /favorites` - Favorites list
- `POST /logout` - Logout

### API Routes (HTMX/JSON)
- `GET /api/movies/search?q=query` - Search movies
- `GET /api/movies/popular` - Popular movies
- `GET /api/movies/trending` - Trending movies
- `POST /api/favorites` - Add to favorites
- `PATCH /api/favorites/:id/status` - Update status
- `PATCH /api/favorites/:id/rating` - Update rating
- `DELETE /api/favorites/:id` - Remove from favorites
- `GET /api/stats` - User statistics

## 🏗 Development

### Running in Development Mode

```bash
# Watch for changes (requires air)
go install github.com/cosmtrek/air@latest
air

# Or run normally
go run main.go
```

### Database Migrations

The application uses GORM's AutoMigrate feature. Tables are created automatically on startup.

To manually create the database schema:

```sql
-- Users table
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(50) UNIQUE NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    password_hash VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Favorite movies table
CREATE TABLE favorite_movies (
    id SERIAL PRIMARY KEY,
    user_id INTEGER REFERENCES users(id) ON DELETE CASCADE,
    tmdb_id INTEGER NOT NULL,
    title VARCHAR(255) NOT NULL,
    overview TEXT,
    release_date DATE,
    poster_path VARCHAR(255),
    genre_ids JSONB,
    status VARCHAR(20) DEFAULT 'por_ver',
    rating INTEGER CHECK (rating >= 1 AND rating <= 10),
    notes TEXT,
    recommended_by VARCHAR(100),
    added_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    watched_at TIMESTAMP,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    UNIQUE(user_id, tmdb_id)
);
```

## 🔐 Security Features

- Password hashing using bcrypt
- Session-based authentication with secure cookies
- CSRF protection through SameSite cookie attribute
- Input validation and sanitization
- SQL injection prevention through GORM

## 🚀 Production Deployment

### Environment Setup

Set `ENVIRONMENT=production` in your production environment.

### Database
- Use a managed PostgreSQL service (AWS RDS, Google Cloud SQL, etc.)
- Enable SSL connections by removing `sslmode=disable` from DATABASE_URL
- Regular backups recommended

### Security
- Use strong, unique SESSION_SECRET
- Enable HTTPS and set secure cookie flags
- Consider adding rate limiting
- Regular security updates

### Performance
- Enable Gin's release mode
- Use a reverse proxy (nginx, Cloudflare)
- Implement caching for TMDB API responses
- Monitor with APM tools

## 🤝 Contributing

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## 📝 License

This project is licensed under the MIT License - see the LICENSE file for details.

## 🙏 Acknowledgments

- [The Movie Database (TMDB)](https://www.themoviedb.org/) for the excellent movie API
- [HTMX](https://htmx.org/) for making frontend interactions seamless
- [Gin Framework](https://gin-gonic.com/) for the fast and flexible web framework
- [Tailwind CSS](https://tailwindcss.com/) for the beautiful styling

## 📞 Support

If you have any questions or run into issues:

1. Check the [Issues](../../issues) page
2. Create a new issue with detailed information
3. Include your Go version, PostgreSQL version, and error logs

---

**Happy Movie Tracking! 🎬✨**