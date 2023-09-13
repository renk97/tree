package model

// DB I/O
type Tree struct {
	Id     int    `json:"id"`
	Root   string `json:"root"`
	Leaves []byte `json:"leaves"`
}

// API I/O
type IOTree struct {
	Id     int      `json:"id"`
	Root   string   `json:"root"`
	Leaves []string `json:"leaves"`
}

// pending
type Leaf struct {
	Words    []string `json:"words"`
	SigBytes int      `json:"sigBytes"`
}

// pending
type HashIOTree struct {
	Id     int    `json:"id"`
	Root   string `json:"root"`
	Leaves []Leaf `json:"leaves"`
}
