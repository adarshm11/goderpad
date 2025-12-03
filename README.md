# goderpad

Live coding sandbox tool, specifically designed for SJSU SCE.

**Note:** Vite is used for this but frontend will have to be migrated to vanilla React when integrated with Clark.

## Tech Stack

- **[Gorilla WebSocket](https://pkg.go.dev/github.com/gorilla/websocket#section-readme)** - WebSocket implementation for real-time communication
- **[Monaco Editor](https://www.npmjs.com/package/%40monaco-editor/react)** - React wrapper for the Monaco Editor (VSCode's editor)

## Clerk webhook requirements (development)

1. Create `backend/.env` with
   `CLERK_WEBHOOK_SIGNING_SECRET=YOUR_SIGNING_SECRET`. This file is git-ignored.
2. Start the backend: `cd backend && go run main.go`.
3. Expose port 8080 with ngrok: `ngrok http --url=<YOUR_STATIC_DOMAIN> 8080`.
4. In the Clerk dashboard, point the webhook endpoint to
   `<YOUR_STATIC_DOMAIN>/api/webhooks`.
5. **Note:** Currently, only the `user.created` event is supported.