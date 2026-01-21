import { useState, useEffect, useContext } from 'react';
import { useParams, useNavigate } from 'react-router-dom';
import { DarkModeContext } from '../../App';
import Popup from '../popup/Popup';
import { getInterviewContent } from '../../api/api';

function PastInterviewPage() {
  const { isDark } = useContext(DarkModeContext);
  const navigate = useNavigate();
  const { interviewId } = useParams<{ interviewId: string }>();
  const [showPopup, setShowPopup] = useState(false);
  const [interviewContent, setInterviewContent] = useState<string>('');
  const [roomName, setRoomName] = useState<string>('');

  useEffect(() => {
    if (!interviewId) {
      navigate('/');
    }
    const fetchInterviewContent = async () => {
      try {
        const response = await getInterviewContent(interviewId!);
        if (!response.ok) {
          setShowPopup(true);
        }
        setInterviewContent(response.data.document);
        setRoomName(response.data.roomName.substring(0, response.data.roomName.length - 3)); // remove .js extension
      } catch (error) {
        setShowPopup(true);
      }
    }
    fetchInterviewContent();
  }, [interviewId, navigate]);

  return (<>
    <Popup 
      message="sorry, an error occurred trying to fetch the interview content"
      buttonText="ok"
      isOpen={showPopup}
      onClickButton={() => {
        setShowPopup(false);
        navigate('/');
      }}
    />
    <div className={`min-h-screen p-8 ${isDark ? 'bg-slate-900 text-white' : 'bg-gray-100 text-gray-900'}`}>
      <h1 className='text-3xl font-bold text-center mb-6'>{roomName}</h1>
      <pre className={`whitespace-pre-wrap break-all bg-gray-200 p-6 rounded-lg shadow-md
        ${isDark ? 'bg-slate-800 text-white' : 'bg-white text-gray-900'}`}>
        {interviewContent}
      </pre>
    </div>

  </>);
}
export default PastInterviewPage;
