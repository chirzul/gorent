# gorent

Go Rent is an backend API for vehicle rent system built on Go Language and Cloudinary for Media API

## 🛠️ Installation Steps

1. Clone the repository

```bash
git clone https://github.com/chirzul/gorent.git
```

2. Install dependencies

```bash
go get -u ./...
# or
go mod tidy
```

3. Database Migration and Rollback

```bash
go run . migrate --up //for database migration
# or
go run . migrate --down //for rollback
```

4. Add Env

```sh
  JWT_KEYS = Your JWT Key
  APP_PORT = Your Port
  # Database
  DB_USER = Your DB User
  DB_HOST = Your DB Host
  DB_NAME = Your DB Name
  DB_PASSWORD = Your DB Password
  # Cloudinary
  CLOUDINARY_NAME = Your Cloudinary cloud name
  CLOUDINARY_API_KEY = Your Cloudinary API key
  CLOUDINARY_API_SECRET = Your Cloudinary API secret
```

5. Run the app

```bash
go run . serve
```
