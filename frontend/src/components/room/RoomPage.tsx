import { useParams, useNavigate } from "react-router-dom";
import { useState, useEffect, useContext } from "react";
import { joinRoom, getRoomDetails } from "../../api/api";
import EnterName from "./EnterName";
import CodeEditor from "./CodeEditor";
import { DarkModeContext } from "../../App";

function RoomPage() {
  const { roomId } = useParams<{ roomId: string }>();
  const navigate = useNavigate();
  const { isDark } = useContext(DarkModeContext);
  const [name, setName] = useState('');
  const [isJoined, setIsJoined] = useState(false);
  const [isLoading, setIsLoading] = useState(false);
  const [roomName, setRoomName] = useState('sce interview');

  useEffect(() => {
    const storedData = localStorage.getItem(`room-${roomId}-name`);
    if (storedData) {
      try {
        const { name: storedName, expiry } = JSON.parse(storedData);
        const now = new Date().getTime();
        if (now < expiry) {
          setName(storedName);
          setIsJoined(true);
        } else {
          localStorage.removeItem(`room-${roomId}-name`);
        }
      } catch (e) {
        localStorage.removeItem(`room-${roomId}-name`);
      }
    }
  }, [roomId]);

  useEffect(() => {
    const fetchRoomDetails = async () => {
      if (roomId) {
        const response = await getRoomDetails(roomId);
        if (response.success && response.data.roomName) {
          setRoomName(response.data.roomName);
        }
      }
    };
    fetchRoomDetails();
  }, [roomId]);

  const handleJoinRoom = async () => {
    if (!name.trim() || !roomId) return;
    
    setIsLoading(true);
    const response = await joinRoom(name, roomId);
    setIsLoading(false);

    if (response.success) {
      const now = new Date().getTime();
      const expiry = now + (24 * 60 * 60 * 1000); // 24 hours
      const data = JSON.stringify({ name, expiry });
      localStorage.setItem(`room-${roomId}-name`, data);
      setIsJoined(true);
    } else {
      alert(response.error || 'Failed to join room');
      navigate('/');
    }
  };

  if (!isJoined) {
    return (
      <EnterName
        roomId={roomId || ''}
        name={name}
        setName={setName}
        isLoading={isLoading}
        onJoinRoom={handleJoinRoom}
      />
    );
  }

  return (
    <div className={`min-h-screen ${isDark ? 'bg-slate-900 text-white' : 'bg-gray-100 text-gray-900'}`}>
      <h1 className="absolute top-6 left-0 right-0 text-center text-2xl font-bold text-white z-10">
        {roomName}
      </h1>
      <CodeEditor />
    </div>
  );
}

export default RoomPage;