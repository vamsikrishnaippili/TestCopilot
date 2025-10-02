# TeraBox Video Downloader

A ready-to-deploy application to download videos from TeraBox.  
**Stack:** React (frontend), Golang (backend).

## Features

- Paste your TeraBox video link and download instantly
- Select available video formats
- No registration/sign-in required
- Works on desktop, tablet, and mobile

---

## Local Development

### 1. Clone the repository

```bash
git clone https://github.com/vamsikrishnaippili/TestCopilot.git
cd TestCopilot
```

### 2. Start the Go backend

```bash
cd backend
go run main.go handler.go terabox.go
```
By default, it runs on [http://localhost:8080](http://localhost:8080)

### 3. Start the React frontend

```bash
cd frontend
npm install
npm start
```
The frontend runs on [http://localhost:3000](http://localhost:3000)

---

## Deployment

- Deploy the frontend to Vercel, Netlify, or similar.
- Deploy the backend (Go) to Render, Railway, or any cloud provider.
- Update the frontend API URLs in production if backend is hosted elsewhere.

---

## Note

This starter uses mock backend logic for TeraBox (for demonstration only).  
Production version requires real TeraBox link parsing and video streaming logic (to be implemented in `backend/terabox.go`).
