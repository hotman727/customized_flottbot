package utils

import (
	"strings"

	"github.com/target/flottbot/models"
)

// GetRoomIDs helps find a room by name, if we have 'cached' it
func GetRoomIDs(wantRooms []string, bot *models.Bot) []string {
	rooms := []string{}

	for _, room := range wantRooms {
		roomMatch := bot.Rooms[strings.ToLower(room)]
		if len(roomMatch) > 0 {
			rooms = append(rooms, roomMatch)
		} else {
			bot.Log.Debugf("Room '%s' does not exist", room)
		}
	}

	return rooms
}
