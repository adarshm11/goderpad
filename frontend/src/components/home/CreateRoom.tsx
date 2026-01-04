import { useState } from 'react';
import { useNavigate } from 'react-router-dom';
import { createRoom } from '../../api/api';

interface CreateRoomProps {
  onSwitchToJoin: () => void;
  isDark: boolean;
}

function CreateRoom({ onSwitchToJoin, isDark }: CreateRoomProps) {
  const [name, setName] = useState('');
  const [roomName, setRoomName] = useState('');
  const navigate = useNavigate();

  const handleCreateRoom = async () => {
    let response;
    if (roomName === '') {
      response = await createRoom(name, `${name}'s sce interview`);
    } else {
      response = await createRoom(name, roomName);
    }
    if (response.success) {
      const expiry = new Date().getTime() + (24 * 60 * 60 * 1000); // 24 hours
      const data = JSON.stringify({ name, expiry });
      localStorage.setItem(`room-${response.data.roomId}-name`, data);
      navigate(`/${response.data.roomId}`);
    } else {
      alert(response.error || 'Failed to create room');
    }
  };

  return (
    <div className='flex flex-col gap-6 w-full max-w-md mx-auto'>
      <h2 className='text-3xl font-semibold mb-2 text-center'>create an interview room</h2>
      
      <div className='flex flex-col gap-3'>
        <label htmlFor='name' className='text-lg font-medium'>your name</label>
        <input
          id='name'
          type='text'
          value={name}
          onChange={(e) => { setName(e.target.value); }}
          onKeyDown={(e) => {
            if (e.key === 'Enter' && name.trim()) {
              handleCreateRoom();
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

      <div className='flex flex-col gap-3'>
        <label htmlFor='roomName' className='text-lg font-medium'>
          room name <span className={`text-sm ${isDark ? 'text-gray-400' : 'text-gray-600'}`}>(optional)</span>
        </label>
        <input
          id='roomName'
          type='text'
          value={roomName}
          onChange={(e) => { setRoomName(e.target.value); }}
          onKeyDown={(e) => {
            if (e.key === 'Enter' && name.trim()) {
              handleCreateRoom();
            }
          }}
          placeholder='name your room'
          className={`px-5 py-4 text-lg rounded-lg focus:outline-none focus:border-blue-500 ${
            isDark
              ? 'bg-slate-800 border border-slate-700 text-white'
              : 'bg-white border border-gray-300 text-gray-900'
          }`}
        />
      </div>

      <button
        onClick={handleCreateRoom}
        className='mt-4 px-6 py-4 text-lg bg-green-600 text-white rounded-lg hover:bg-green-700 cursor-pointer transition-colors disabled:opacity-50'
        disabled={!name.trim()}
      >
        create room
      </button>

      <p className={`text-center text-sm mt-4 ${isDark ? 'text-gray-400' : 'text-gray-600'}`}>
        or{' '}
        <button
          onClick={onSwitchToJoin}
          className={`cursor-pointer hover:underline ${isDark ? 'text-blue-400' : 'text-blue-600'}`}
        >
          join an existing room
        </button>
      </p>
    </div>
  );
}

export default CreateRoom;