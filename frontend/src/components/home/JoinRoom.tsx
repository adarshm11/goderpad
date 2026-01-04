import { useState } from 'react';
import { useNavigate } from 'react-router-dom';
import { joinRoom } from '../../api/api';

interface JoinRoomProps {
  onSwitchToCreate: () => void;
  isDark: boolean;
}

function JoinRoom({ onSwitchToCreate, isDark }: JoinRoomProps) {
  const [roomId, setRoomId] = useState('');
  const [name, setName] = useState('');
  const navigate = useNavigate();

  const handleJoinRoom = async () => {
    const response = await joinRoom(name, roomId);
    if (response.success) {
      navigate(`/${roomId}`);
    } else {
      alert(response.error || 'Failed to join room');
    }
  }

  return (
    <div className='flex flex-col gap-6 w-full max-w-md mx-auto'>
      <h2 className='text-3xl font-semibold mb-2 text-center'>join an interview room</h2>
      
      <div className='flex flex-col gap-3'>
        <label htmlFor='roomId' className='text-lg font-medium'>room ID</label>
        <input
          id='roomId'
          type='text'
          value={roomId}
          onChange={(e) => setRoomId(e.target.value)}
          placeholder='enter room ID'
          className={`px-5 py-4 text-lg rounded-lg focus:outline-none focus:border-blue-500 ${
            isDark
              ? 'bg-slate-800 border border-slate-700 text-white'
              : 'bg-white border border-gray-300 text-gray-900'
          }`}
        />
      </div>

      <div className='flex flex-col gap-3'>
        <label htmlFor='name' className='text-lg font-medium'>your name</label>
        <input
          id='name'
          type='text'
          value={name}
          onChange={(e) => setName(e.target.value)}
          onKeyDown={(e) => {
            if (e.key === 'Enter' && name.trim() && roomId.trim()) {
              handleJoinRoom();
            }
          }}
          placeholder='enter your name'
          className={`px-5 py-4 text-lg rounded-lg focus:outline-none focus:border-blue-500 ${
            isDark
              ? 'bg-slate-800 border border-slate-700 text-white'
              : 'bg-white border border-gray-300 text-gray-900'
          }`}
        />
      </div>

      <button
        className='mt-4 px-6 py-4 text-lg bg-blue-600 text-white rounded-lg hover:bg-blue-700 cursor-pointer transition-colors disabled:opacity-50 disabled:cursor-not-allowed'
        onClick={handleJoinRoom}
        disabled={!roomId.trim() || !name.trim()}
      >
        join room
      </button>

      <p className={`text-center text-sm mt-4 ${isDark ? 'text-gray-400' : 'text-gray-600'}`}>
        or{' '}
        <button
          onClick={onSwitchToCreate}
          className={`cursor-pointer hover:underline ${isDark ? 'text-blue-400' : 'text-blue-600'}`}
        >
          create your own room
        </button>
      </p>
    </div>
  )
}

export default JoinRoom;