package defs

type UserCredential struct {
	Username string `json:"user_name"`
	Password string `json:"pwd"`
}

type SignedUp struct {
	Success   bool   `json:"success"`
	SessionId string `json:"session_id"`
}

type VideoInfo struct {
	Id              string
	AuthorId        int
	Name            string
	DisplyCreatTime string
}

type Comment struct {
	Id      string
	Author  string
	VideoId string
	Content string
}

type SimpleSession struct {
	Username string
	TTL      int64
}
