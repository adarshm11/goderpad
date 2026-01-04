import { useState, useContext } from 'react';
import CreateRoom from './CreateRoom';
import JoinRoom from './JoinRoom';
import { DarkModeContext } from '../../App';

function HomePage() {
  const [showCreateRoom, setShowCreateRoom] = useState(false);
  const { isDark } = useContext(DarkModeContext);

  return (
    <div className={`min-h-screen ${isDark ? 'bg-slate-900 text-white' : 'bg-gray-100 text-gray-900'}`}>
      <h1 className='text-4xl font-bold text-center pt-20'>welcome to goderpad</h1>
      <h3 className='text-xl text-center pt-4'>sce's interview platform</h3>
      <div className='flex flex-row gap-30 justify-center mt-20'>
        {showCreateRoom ? (
          <CreateRoom onSwitchToJoin={() => setShowCreateRoom(false)} isDark={isDark} />
        ) : (
          <JoinRoom onSwitchToCreate={() => setShowCreateRoom(true)} isDark={isDark} />
        )}
      </div>
    </div>
  );
}

export default HomePage;
