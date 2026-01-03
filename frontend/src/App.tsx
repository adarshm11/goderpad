import { BrowserRouter, Routes, Route } from "react-router-dom";
import RoomPage from "./components/room/RoomPage";
import HomePage from "./components/home/HomePage";

function App() {
  return (
    <BrowserRouter>
      <Routes>
        <Route path="/" element={<HomePage />} />
        <Route path="/:roomId" element={<RoomPage />} />
      </Routes>
    </BrowserRouter>
  );
}

export default App
