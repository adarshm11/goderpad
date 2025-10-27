Broadcast Architecture

  To share messages across all connected frontends, you'd need:

  1. Connection Hub/Manager

  A central coordinator that maintains a registry of all active WebSocket connections. This would typically be a struct with:
  - A map/slice of active connections
  - Channels for registering new connections
  - Channels for unregistering disconnected connections
  - A broadcast channel for messages to distribute

  2. Hub Goroutine

  A single long-running goroutine that:
  - Listens for new connection registrations
  - Listens for disconnection events
  - Receives messages from any client
  - Broadcasts those messages to all registered connections
  - Handles cleanup when connections close

  3. Per-Connection Goroutines

  Each WebSocket connection would have two goroutines:
  - Reader: Reads messages from the client and sends them to the hub's broadcast channel
  - Writer: Listens for messages from the hub and writes them to the client

  4. Message Flow

  1. Client A sends a message → Reader goroutine receives it
  2. Reader sends message to hub's broadcast channel
  3. Hub receives message and iterates through all registered connections
  4. Hub sends message to each connection's outbound channel
  5. Each Writer goroutine receives and sends to its client
  6. All clients (including Client A) see the message

  5. Lifecycle Management

  - When /ws handler is called, create the connection goroutines and register with hub
  - If a client disconnects or errors, unregister from hub and clean up goroutines
  - The hub needs to be initialized once when your server starts

  This pattern is commonly called a pub-sub hub or broadcast hub - it's the standard way to build chat rooms, collaborative
  editors, or any real-time multi-user application.