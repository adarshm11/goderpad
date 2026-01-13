import { useState, useContext } from 'react';
import { createRoom } from '../../api/api';
import { useNavigate } from 'react-router-dom';
import { DarkModeContext } from '../../App';
import { UserContext } from '../../App';

function HomePage() {
  const { isDark } = useContext(DarkModeContext);
  const [userName, setUserName] = useState('');
  const [roomName, setRoomName] = useState('');
  const navigate = useNavigate();
  const { userId } = useContext(UserContext);

  const handleCreateRoom = async () => {
    let response;
    if (roomName === '') {
      response = await createRoom(userId, userName, `${userName}'s sce interview`);
    } else {
      response = await createRoom(userId, userName, roomName);
    }
    if (response.success) {
      const roomId = response.data.roomId;
      const expiry = new Date().getTime() + (24 * 60 * 60 * 1000); // 24 hours
      const data = JSON.stringify({ userName, expiry });
      localStorage.setItem(`goderpad-cookie-${roomId}`, data);
      navigate(`/${roomId}`);
    } else {
      alert(response.error || 'Failed to create room');
    }
  };

  return (
    <div className={`min-h-screen ${isDark ? 'bg-slate-900 text-white' : 'bg-gray-100 text-gray-900'}`}>
      <h1 className='text-4xl font-bold text-center pt-20'>welcome to goderpad</h1>
      <h3 className='text-xl text-center pt-4'>sce's interview platform</h3>
      <div className='flex flex-row gap-30 justify-center mt-20'>
        <div className='flex flex-col gap-6 w-full max-w-md mx-auto'>
          <h2 className='text-3xl font-semibold mb-2 text-center'>create an interview room</h2>

          <div className='flex flex-col gap-3'>
            <label htmlFor='name' className='text-lg font-medium'>your name</label>
            <input
              id='name'
              type='text'
              value={userName}
              onChange={(e) => { setUserName(e.target.value); }}
              onKeyDown={(e) => {
                if (e.key === 'Enter' && userName.trim()) {
                  handleCreateRoom();
                }
              }}
              placeholder='enter your name'
              className={`px-5 py-4 text-lg rounded-lg focus:outline-none focus:border-blue-500 ${isDark
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
                if (e.key === 'Enter' && userName.trim()) {
                  handleCreateRoom();
                }
              }}
              placeholder='name your room'
              className={`px-5 py-4 text-lg rounded-lg focus:outline-none focus:border-blue-500 ${isDark
                  ? 'bg-slate-800 border border-slate-700 text-white'
                  : 'bg-white border border-gray-300 text-gray-900'
                }`}
            />
          </div>

          <button
            onClick={handleCreateRoom}
            className='mt-4 px-6 py-4 text-lg bg-green-600 text-white rounded-lg hover:bg-green-700 cursor-pointer transition-colors disabled:opacity-50'
            disabled={!userName.trim()}
          >
            create room
          </button>
        </div>
      </div>
    </div>
  );
}

export default HomePage;
