import { useContext, useEffect, useState, useRef } from 'react';
import { DarkModeContext } from '../../App';

interface PopupProps {
  message: string;
  buttonText: string;
  onClickButton: () => void;
  isOpen: boolean;
}

function Popup({ message, buttonText, onClickButton, isOpen }: PopupProps) {
  const { isDark } = useContext(DarkModeContext);
  const [isClosing, setIsClosing] = useState(false);
  const [shouldRender, setShouldRender] = useState(isOpen);
  const isClosingRef = useRef(false);

  useEffect(() => {
    if (isOpen) {
      setShouldRender(true);
      setIsClosing(false);
      isClosingRef.current = false;
    } else if (shouldRender && !isClosingRef.current) {
      isClosingRef.current = true;
      setIsClosing(true);
      const timer = setTimeout(() => {
        setShouldRender(false);
        setIsClosing(false);
      }, 200); // Match animation duration
      return () => clearTimeout(timer);
    }
  }, [isOpen, shouldRender]);

  useEffect(() => {
    if (!isOpen) return;

    const handleKeyDown = (e: KeyboardEvent) => {
      if (e.key === 'Enter' && !isClosingRef.current) {
        e.preventDefault();
        e.stopPropagation();
        onClickButton();
      }
    };

    window.addEventListener('keydown', handleKeyDown, true);
    return () => window.removeEventListener('keydown', handleKeyDown, true);
  }, [isOpen, onClickButton]);

  if (!shouldRender) return null;

  return (
    <div className={`fixed inset-0 z-50 flex items-center justify-center backdrop-blur-sm pointer-events-none ${isClosing ? 'animate-fade-out' : 'animate-fade-in'}`}>
      <div className={`rounded-xl shadow-2xl p-6 max-w-md w-full mx-4 pointer-events-auto ${isClosing ? 'animate-zoom-out' : 'animate-zoom-in'} ${isDark ? 'bg-slate-800 text-white' : 'bg-white text-gray-900'}`}>
        <p className={`text-base mb-6 ${isDark ? 'text-gray-300' : 'text-gray-700'}`}>
          {message}
        </p>
        <div className="flex justify-end">
          <button 
            onClick={() => onClickButton()} 
            className="px-6 py-2.5 bg-green-600 text-white rounded-lg hover:bg-green-700 transition-colors font-semibold"
          >
            {buttonText}
          </button>
        </div>
      </div>
    </div>
  );
}
export default Popup;
