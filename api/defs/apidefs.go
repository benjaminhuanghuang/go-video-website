package defs

type UserCredential struct {
	Username string `json: "user_name"`
	Password string `json: "pwd"`
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
