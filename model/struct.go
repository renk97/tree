package model

type Tree struct {
	Id     int    `json:"id"`
	Root   string `json:"root"`
	Leaves []byte `json:"leaves"`
}

type IOTree struct {
	Id       int      `json:"id"`
	Root     string   `json:"root"`
	Leaves   []string `json:"leaves"`
	SigBytes int      `json:"sigBytes"`
}
