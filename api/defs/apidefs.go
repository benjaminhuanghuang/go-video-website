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
