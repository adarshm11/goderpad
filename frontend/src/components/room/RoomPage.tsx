import { useParams, useNavigate } from 'react-router-dom';
import { useState, useEffect, useContext } from 'react';
import { joinRoom } from '../../api/api';
import EnterName from './EnterName';
import CodeEditor from './CodeEditor';
import { DarkModeContext, UserContext } from '../../App';
import { DEFAULT_CODE } from '../../util/constants';

function RoomPage() {
  const { roomId } = useParams<{ roomId: string }>();
  const navigate = useNavigate();
  const { isDark } = useContext(DarkModeContext);
  const { userId } = useContext(UserContext);
  const [userName, setUserName] = useState('');
  const [isJoined, setIsJoined] = useState(false);
  const [isLoading, setIsLoading] = useState(false);
  const [roomName, setRoomName] = useState('sce interview');
  const [code, setCode] = useState(DEFAULT_CODE);
  const [ws, setWs] = useState<WebSocket | null>(null);
  const [users, setUsers] = useState<Array<{
    userId: string;
    userName: string;
    cursorPosition: {
      lineNumber: number;
      column: number 
    } | null
  }>>([]);

  const handleJoinRoom = async () => {
    if (!userName.trim() || !roomId) return;
    
    setIsLoading(true);
    const response = await joinRoom(userId, userName, roomId);
    setIsLoading(false);

    if (response.ok) {
      setRoomName(response.data.roomName || 'sce interview');
      setCode(response.data.document || DEFAULT_CODE);
      setUsers(response.data.users || []);
      const now = new Date().getTime();
      const expiry = now + (24 * 60 * 60 * 1000);
      const data = JSON.stringify({ userName, expiry });
      localStorage.setItem(`goderpad-cookie-${roomId}`, data);
      setIsJoined(true);
    } else {
      alert(response.error || 'Failed to join room');
      navigate('/');
    }
  };

  const sendWsMessage = (message: any) => {
    if (ws && ws.readyState === WebSocket.OPEN) {
      ws.send(JSON.stringify(message));
    }
  };

  useEffect(() => {
    if (!roomId) {
      navigate('/');
      return;
    }

    const storedData = localStorage.getItem(`goderpad-cookie-${roomId}`);
    if (!storedData) return;

    const joinWithStoredData = async () => {
      try {
        const { userName: storedUserName, expiry } = JSON.parse(storedData);
        const now = new Date().getTime();
        
        if (now >= expiry) {
          localStorage.removeItem(`goderpad-cookie-${roomId}`);
          return;
        }

        const response = await joinRoom(userId, storedUserName, roomId);
        
        if (response.ok) {
          setRoomName(response.data.roomName || 'sce interview');
          setCode(response.data.document || DEFAULT_CODE);
          setUsers(response.data.users || []);
          setUserName(storedUserName);
          setIsJoined(true);
          
          // Update expiry
          const updatedExpiry = now + (24 * 60 * 60 * 1000);
          localStorage.setItem(`goderpad-cookie-${roomId}`, JSON.stringify({ userName: storedUserName, expiry: updatedExpiry }));
        } else {
          localStorage.removeItem(`goderpad-cookie-${roomId}`);
        }
      } catch (e) {
        localStorage.removeItem(`goderpad-cookie-${roomId}`);
      }
    };

    joinWithStoredData();
  }, [roomId, userId, navigate]);

  // Setup WebSocket connection and handlers when the user successfully joins the room
  useEffect(() => {
    if (!isJoined || !roomId) return;
    const websocket = new WebSocket(`ws://localhost:8080/ws/${roomId}`);
    setWs(websocket);

    websocket.onopen = async () => {
      sendWsMessage({
        userId,
        type: 'user_joined',
        payload: {
          roomId,
          userName
        }
      });
    }

    websocket.onclose = () => {
      sendWsMessage({
        userId,
        type: 'user_left',
        payload: {
          roomId,
        }
      })
    }

    websocket.onmessage = (event) => {
      const message = JSON.parse(event.data);

      switch (message.type) {
        case 'cursor_update':
          const user = users.find(u => u.userId === message.payload.userId);
          if (user) {
            user.cursorPosition = {
              lineNumber: message.payload.lineNumber,
              column: message.payload.column
            };
            setUsers([...users]);
          }
          break;

        case 'user_joined':
          setUsers(prevUsers => [
            ...prevUsers,
            {
              userId: message.payload.userId,
              userName: message.payload.userName,
              cursorPosition: {
                lineNumber: 1,
                column: 1
              }
            }
          ])
          break;

        case 'user_left':
          setUsers(prevUsers => prevUsers.filter(u => u.userId !== message.payload.userId));
          break;

        case 'code_update':
          setCode(message.payload.code);
          break;

        default:
          break;
      }
    }

    return () => {
      websocket.close();
      setWs(null);
    };
  }, [isJoined, roomId]);

  if (!isJoined) {
    return (
      <EnterName
        roomName={roomName}
        userName={userName}
        setUserName={setUserName}
        isLoading={isLoading}
        onJoinRoom={handleJoinRoom}
      />
    );
  }

  return (
    <div className={`min-h-screen ${isDark ? 'bg-slate-900 text-white' : 'bg-gray-100 text-gray-900'}`}>
      <div className='relative'>
        <h1 className={`absolute top-6 left-0 right-0 text-center text-2xl font-bold z-10 ${isDark ? 'text-white' : 'text-gray-900'}`}>
          {roomName}
        </h1>
      </div>
      <CodeEditor
        code={code}
        setCode={setCode}
        sendWsMessage={sendWsMessage}
        users={users}
      />
    </div>
  );
}

export default RoomPage;