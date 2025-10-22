# Shared Text Editor MVP Plan

## Architecture Overview

- **Backend**: Go WebSocket server using Gorilla WebSocket with room-based
  document management
- **Frontend**: React + Vite with Monaco Editor and WebSocket client
- **Sync Strategy**: Last-write-wins with broadcast to all room participants
- **Persistence**: In-memory only for MVP (no database)

## Backend Implementation

### 1. Go Dependencies & Setup

Update `backend/go.mod` to Go 1.23.1 and add Gorilla WebSocket:

```go
require github.com/gorilla/websocket v1.5.1
```

### 2. WebSocket Server Structure (`backend/main.go`)

**Key Components:**

- **Client struct**: Holds connection, room ID, user ID, and cursor position
- **Room struct**: Manages document content, connected clients, and broadcasts
- **Hub**: Manages all rooms (map[roomID]*Room)
- **Message types**: text-change, cursor-update, user-joined, user-left

**Core Functions:**

- `handleConnections()`: Upgrade HTTP to WebSocket, join/create room
- `handleMessages()`: Process incoming changes and broadcast to room
- `broadcastToRoom()`: Send updates to all clients in a room except sender
- Room-based isolation so edits only go to users in same room

### 3. API Endpoints

- `GET /ws?room=<roomID>&user=<userID>`: WebSocket connection endpoint
- `GET /health`: Health check

## Frontend Implementation

### 1. Dependencies (`frontend/package.json`)

Add:

- `@monaco-editor/react`: Monaco Editor wrapper
- No additional WebSocket library needed (native WebSocket API)

### 2. Component Structure

**Main Components:**

- `App.jsx`: Routing and room ID input page
- `Editor.jsx`: Monaco Editor with WebSocket integration
- `UserPresence.jsx`: Display connected users and their cursors
- `useWebSocket.js`: Custom hook for WebSocket connection management

### 3. Editor Component (`frontend/src/components/Editor.jsx`)

**Features:**

- Monaco Editor instance with language support (JavaScript/Python/etc.)
- WebSocket connection to `ws://localhost:8080/ws?room=<roomID>&user=<userID>`
- Listen for editor changes and send via WebSocket
- Receive remote changes and apply to editor (skip if local change)
- Track and display other users' cursor positions as decorations
- Handle user join/leave notifications

**Message Flow:**

```
User types → onChange event → Send to WS → Server broadcasts → Other clients update editor
User moves cursor → onCursorChange → Send to WS → Server broadcasts → Show remote cursors
```

### 4. Room Entry (`frontend/src/components/RoomEntry.jsx`)

- Input field for room ID
- Generate random user ID (UUID)
- Navigate to `/room/<roomID>` when joining

### 5. Styling

- Modern, clean UI with dark theme option
- Monaco Editor full-height layout
- Sidebar showing connected users with color-coded indicators
- Cursor overlays with username labels

## Message Protocol

### WebSocket Message Format (JSON)

```json
{
  "type": "text-change" | "cursor-update" | "user-joined" | "user-left",
  "userId": "uuid",
  "roomId": "string",
  "content": "full document text",
  "cursor": { "line": 1, "column": 5 }
}
```

## Implementation Order

1. Set up Go WebSocket server with room management
2. Implement broadcast logic for rooms
3. Install Monaco Editor in frontend
4. Create room entry UI
5. Build Editor component with Monaco integration
6. Connect WebSocket client to server
7. Implement text synchronization
8. Add cursor tracking and presence indicators
9. Test with multiple browser windows

## Testing Strategy

- Manual testing with 2-3 browser windows/tabs
- Join same room ID from multiple windows
- Verify text changes sync in real-time
- Verify cursors display correctly
- Test edge cases: disconnect/reconnect, rapid typing

### To-dos

- [ ] Set up Go backend with Gorilla WebSocket, room management, and broadcast
      logic
- [ ] Install Monaco Editor and set up frontend dependencies
- [ ] Create room entry UI component for joining rooms
- [ ] Build Editor component with Monaco Editor and WebSocket integration
- [ ] Implement bidirectional text synchronization between clients
- [ ] Add cursor tracking and user presence indicators
- [ ] Test multi-user collaboration with multiple browser windows
