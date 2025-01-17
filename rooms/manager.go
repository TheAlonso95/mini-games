package rooms

import (
	"net/http"
	"sync"

	"github.com/google/uuid"
)

var (
	rooms   = make(map[string]*Room)
	roomsMu sync.Mutex
)

type Room struct {
	ID    string
	Game  interface{}
	Mutex sync.Mutex
}

func CreateRoomHandler(w http.ResponseWriter, r *http.Request) {
	roomsMu.Lock()
	defer roomsMu.Unlock()

	roomID := generateRoomID()
	rooms[roomID] = &Room{
		ID:   roomID,
		Game: nil, // Game instance will be assigned when a game starts
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(roomID))
}

func JoinRoomHandler(w http.ResponseWriter, r *http.Request) {
	//TODO: Implement JoinRoomHandler logic
}

func GetRoom(roomID string) *Room {
	roomsMu.Lock()
	defer roomsMu.Unlock()
	return rooms[roomID]
}

func generateRoomID() string {
	return uuid.New().String()
}
