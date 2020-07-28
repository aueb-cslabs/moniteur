package types

type RoomMap struct {
	Rooms        map[string]string `yaml:"rooms,omitempty"`
	RoomsInverse map[string]string `yaml:"-"`
}
