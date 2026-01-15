const API_URL = 'http://localhost:8080';

export async function createRoom(userId: string, name: string, roomName: string) {
  try {
    const response = await fetch(`${API_URL}/createRoom`, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify({
        userId,
        name,
        roomName,
      }),
    });
    return await response.json()
  } catch (err) {
    return {
      ok: false,
      error: err instanceof Error ? err.message : 'Network error'
    };
  }
}

export async function joinRoom(userId: string, name: string, roomId: string) {
  try {
    const response = await fetch(`${API_URL}/joinRoom`, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify({
        userId,
        name,
        roomId
      }),
    });
    return await response.json();
  } catch (err) {
    return {
      ok: false,
      error: err instanceof Error ? err.message : 'Network error'
    };
  }
}
