package handlers

type GetStationsArgs struct {
	LocationID string `json:"loc_id"`
}

type StationReply struct {
	Name string
	URL  string
}

type GetStationsReply struct {
	Stations []StationReply `json:"stations,omitempty"`
}
