import { useParams } from "react-router-dom";

function RoomPage() {
  const { roomId } = useParams<{ roomId: string }>();
  return <div>Welcome to Room: {roomId}</div>;
}

export default RoomPage;