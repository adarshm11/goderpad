import { useEffect, useRef, useState } from 'react';
import './App.css';

type MessageType = 'system' | 'received' | 'sent' | 'error';

interface Message {
  type: MessageType;
  text: string;
}

type ConnectionStatus = 'connected' | 'disconnected';

function App() {
  const [messages, setMessages] = useState<Message[]>([]);
  const [connectionStatus, setConnectionStatus] = useState<ConnectionStatus>('disconnected');
  const [messageInput, setMessageInput] = useState('');
  const wsRef = useRef<WebSocket | null>(null);

  const connect = () => {
    if (wsRef.current?.readyState === WebSocket.OPEN) {
      return;
    }

    const ws = new WebSocket('ws://localhost:8080/ws');

    ws.onopen = () => {
      setConnectionStatus('connected');
      setMessages(prev => [...prev, { type: 'system', text: 'Connected to server' }]);
    };

    ws.onmessage = (event: MessageEvent) => {
      setMessages(prev => [...prev, { type: 'received', text: event.data }]);
    };

    ws.onerror = () => {
      setMessages(prev => [...prev, { type: 'error', text: 'WebSocket error occurred' }]);
    };

    ws.onclose = () => {
      setConnectionStatus('disconnected');
      setMessages(prev => [...prev, { type: 'system', text: 'Disconnected from server' }]);
    };

    wsRef.current = ws;
  };

  const disconnect = () => {
    if (wsRef.current) {
      wsRef.current.close();
      wsRef.current = null;
    }
  };

  const sendMessage = () => {
    if (wsRef.current?.readyState === WebSocket.OPEN && messageInput.trim()) {
      wsRef.current.send(messageInput);
      setMessages(prev => [...prev, { type: 'sent', text: messageInput }]);
      setMessageInput('');
    }
  };

  const clearMessages = () => {
    setMessages([]);
  };

  useEffect(() => {
    return () => {
      if (wsRef.current) {
        wsRef.current.close();
      }
    };
  }, []);

  const handleKeyDown = (event: React.KeyboardEvent<HTMLInputElement>) => {
    if (event.key === 'Enter') {
      sendMessage();
    }
  };

  return (
    <div style={{ padding: '20px', maxWidth: '800px', margin: '0 auto' }}>
      <h1>WebSocket Test Client</h1>

      <div style={{ marginBottom: '20px' }}>
        <div style={{ marginBottom: '10px' }}>
          Status: <strong style={{ color: connectionStatus === 'connected' ? 'green' : 'red' }}>
            {connectionStatus.toUpperCase()}
          </strong>
        </div>

        <button
          onClick={connect}
          disabled={connectionStatus === 'connected'}
          style={{ marginRight: '10px' }}
        >
          Connect
        </button>

        <button
          onClick={disconnect}
          disabled={connectionStatus === 'disconnected'}
          style={{ marginRight: '10px' }}
        >
          Disconnect
        </button>

        <button onClick={clearMessages}>
          Clear Messages
        </button>
      </div>

      <div style={{ marginBottom: '20px' }}>
        <input
          type="text"
          value={messageInput}
          onChange={event => setMessageInput(event.target.value)}
          onKeyDown={handleKeyDown}
          placeholder="Type a message..."
          disabled={connectionStatus === 'disconnected'}
          style={{ padding: '8px', width: '300px', marginRight: '10px' }}
        />
        <button
          onClick={sendMessage}
          disabled={connectionStatus === 'disconnected'}
        >
          Send Message
        </button>
      </div>

      <div style={{
        border: '1px solid #ccc',
        padding: '10px',
        height: '400px',
        overflowY: 'auto',
        backgroundColor: '#f5f5f5',
        fontFamily: 'monospace'
      }}>
        <div><strong>Messages:</strong></div>
        {messages.length === 0 ? (
          <div className="font-black" style={{ color: '#999', marginTop: '10px' }}>No messages yet...</div>
        ) : (
          messages.map((msg, index) => (
            <div
              key={index}
              style={{
                padding: '5px',
                marginTop: '5px',
                color: 'black',
                backgroundColor:
                  msg.type === 'system' ? '#e3f2fd' :
                  msg.type === 'sent' ? '#fff3e0' :
                  msg.type === 'error' ? '#ffebee' :
                  '#f1f8e9',
                borderLeft: `3px solid ${
                  msg.type === 'system' ? '#2196F3' :
                  msg.type === 'sent' ? '#FF9800' :
                  msg.type === 'error' ? '#f44336' :
                  '#8BC34A'
                }`
              }}
            >
              <span style={{ fontWeight: 'bold', marginRight: '10px' }}>
                [{msg.type}]
              </span>
              {msg.text}
            </div>
          ))
        )}
      </div>
    </div>
  );
}

export default App;
