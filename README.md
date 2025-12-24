# Beeline Waitlist

A beautiful coming soon page for Beeline with a countdown timer to New Year's Eve 2025.

## Features

- 🎯 **Countdown Timer** to December 31, 2025
- 📧 **Email Waitlist** collection
- 🎨 **Modern Design** with gradient backgrounds
- 📱 **Responsive** for all devices
- 🔧 **Simple Backend** in Go with Fiber
- 💾 **Data Persistence** with JSON storage

## Quick Start

```bash
# Clone or navigate to the directory
cd beeline-waitlist

# Run the server
./run.sh

# Or manually:
go mod tidy
go run server.go
```

The server will start on `http://localhost:8080`

## API Endpoints

- `GET /` - Main waitlist page
- `POST /api/waitlist` - Add email to waitlist
- `GET /api/waitlist/count` - Get current waitlist count
- `GET /health` - Health check

## Deployment

### Local Development
```bash
go run server.go
```

### Production (with PM2)
```bash
pm2 start server.go --name beeline-waitlist
```

### Docker
```dockerfile
FROM golang:1.21-alpine
WORKDIR /app
COPY . .
RUN go mod tidy
RUN go build -o server server.go
EXPOSE 8080
CMD ["./server"]
```

## Customization

- **Launch Date**: Edit the date in `index.html` JavaScript
- **Styling**: Modify the CSS in `index.html`
- **Features**: Update the features section in HTML
- **Backend**: Extend `server.go` for additional functionality

## Waitlist Data

Emails are stored in `waitlist.json` with timestamps. In production, consider:
- Database integration (PostgreSQL, MongoDB)
- Email verification
- Admin dashboard
- Email notifications

## Technologies

- **Frontend**: HTML5, CSS3, JavaScript
- **Backend**: Go with Fiber framework
- **Styling**: Modern CSS with gradients and animations