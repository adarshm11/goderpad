const API_URL = 'http://localhost:8080';

export async function createRoom(name: string, customRoomId?: string) {
  try {
    const response = await fetch(`${API_URL}/createRoom`, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify({
        name,
        customRoomId: customRoomId || undefined,
      }),
    });
    return await response.json()
  } catch (error) {
    return { success: false, error: 'Network error' };
  }
}

export async function joinRoom(name: string, roomId: string) {
  try {
    const response = await fetch(`${API_URL}/joinRoom`, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify({ name, roomId }),
    });
    return await response.json();
  } catch (error) {
    return { success: false, error: 'Network error' };
  }
}
