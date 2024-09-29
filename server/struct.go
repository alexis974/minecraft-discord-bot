package server

type MCServerInfo struct {
	Online  bool        `json:"online"`
	Players PlayersInfo `json:"players"`
}

type PlayersInfo struct {
	Online int      `json:"online"`
	Max    int      `json:"max"`
	List   []Player `json:"list"`
}

type Player struct {
	Name string `json:"name"`
	Uuid string `json:"uuid"`
}
