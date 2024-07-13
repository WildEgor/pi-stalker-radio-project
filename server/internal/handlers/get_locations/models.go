package handlers

type GetLocationArgs struct {
}

type LocationCoordinate struct {
	Lat  string `json:"lat"`
	Long string `json:"long"`
}

type Location struct {
	Name        string
	Coordinates LocationCoordinate
}

type GetLocationReply struct {
	Locations []Location `json:"locations,omitempty"`
}
