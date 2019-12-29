package types

// RoomMap contains a map for english-greek room name mapping
type RoomMap struct {
	Rooms map[string]string `yaml:"rooms,omitempty"`
}
