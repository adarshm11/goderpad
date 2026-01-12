import { useNavigate } from "react-router-dom";

interface EnterNameProps {
  roomName: string;
  name: string;
  setName: (name: string) => void;
  isLoading: boolean;
  onJoinRoom: () => void;
}

function EnterName({ roomName, name, setName, isLoading, onJoinRoom }: EnterNameProps) {
  const navigate = useNavigate();

  return (
    <div className="min-h-screen flex items-center justify-center bg-gray-50 dark:bg-slate-900">
      <div className="flex flex-col gap-6 w-full max-w-md mx-auto p-6">
        <h2 className="text-3xl font-semibold mb-2 text-center dark:text-white">
          join {roomName}
        </h2>
        
        <div className="flex flex-col gap-3">
          <label htmlFor="name" className="text-lg font-medium dark:text-white">
            your name
          </label>
          <input
            id="name"
            type="text"
            value={name}
            onChange={(e) => setName(e.target.value)}
            onKeyDown={(e) => {
              if (e.key === 'Enter' && name.trim()) {
                onJoinRoom();
              }
            }}
            placeholder="enter your name"
            className="px-5 py-4 text-lg rounded-lg focus:outline-none focus:border-blue-500 bg-white border border-gray-300 text-gray-900 dark:bg-slate-800 dark:border-slate-700 dark:text-white"
            autoFocus
          />
        </div>

        <button
          onClick={onJoinRoom}
          disabled={!name.trim() || isLoading}
          className="mt-4 px-6 py-4 text-lg bg-blue-600 text-white rounded-lg hover:bg-blue-700 transition-colors disabled:opacity-50 disabled:cursor-not-allowed"
        >
          {isLoading ? 'joining...' : 'join room'}
        </button>

        <button
          onClick={() => navigate('/')}
          className="text-sm text-gray-600 dark:text-gray-400 hover:underline cursor-pointer"
        >
          ← back to home
        </button>
      </div>
    </div>
  );
}

export default EnterName;