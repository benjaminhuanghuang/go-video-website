package dbops

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

/*
init(db login, truncate tables) -> run tests ->clear data
*/
var tempvid string

func clearTables() {
	dbConn.Exec("truncate users")
	dbConn.Exec("truncate video_info")
	dbConn.Exec("truncate comments")
	dbConn.Exec("truncate sessions")
}

func TestMain(m *testing.M) {
	clearTables()
	m.Run()
	clearTables()
}

func TestUserWorkFlow(t *testing.T) {
	t.Run("Add", testAddUser)
	t.Run("Get", testGetUser)
	t.Run("Delete", testDeleteUser)
	t.Run("Reget", testRegetUser)
}

func testAddUser(t *testing.T) {
	err := AddUserCredential("ben", "123")
	if err != nil {
		t.Error("Error of AddUser: v%", err)
	}
}

func testGetUser(t *testing.T) {
	pwd, err := GetUserCredential("ben")
	if err != nil {
		t.Error("Error of GetUser: v%", err)
	}
	if pwd != "123" {
		t.Error("Error of user password: v%", pwd)
	}
}

func testDeleteUser(t *testing.T) {
	err := DeleteUser("ben", "123")
	if err != nil {
		t.Error("Error of DeleteUser: v%", err)
	}
}

func testRegetUser(t *testing.T) {
	pwd, err := GetUserCredential("ben")
	if err != nil {
		t.Error("Error of RegetUser: v%", err)
	}
	if pwd != "" {
		t.Error("Deleteing user faileds")
	}
}

func TestVideoWorkFlow(t *testing.T) {
	clearTables()
	t.Run("PrepareUser", testAddUser)
	t.Run("AddVideo", testAddVideoInfo)
	t.Run("GetVideo", testGetVideoInfo)
	t.Run("DelVideo", testDeleteVideoInfo)
	t.Run("RegetVideo", testRegetVideoInfo)
}

func testAddVideoInfo(t *testing.T) {
	vi, err := AddNewVideo(1, "my-video")
	if err != nil {
		t.Error("Error of AddVideoInfo: v%", err)
	}
	tempvid = vi.Id
}

func testGetVideoInfo(t *testing.T) {
	_, err := GetVideoInfo("ben")
	if err != nil {
		t.Error("Error of GetUser: v%", err)
	}
}

func testDeleteVideoInfo(t *testing.T) {
	err := DeleteVideoInfo(tempvid)
	if err != nil {
		t.Error("Error of DeleteUser: v%", err)
	}
}

func testRegetVideoInfo(t *testing.T) {
	vid, err := GetVideoInfo(tempvid)
	if err != nil || vid != nil {
		t.Error("Error of RegetVideoInfo: v%", err)
	}
}

func TestComments(t *testing.T) {
	clearTables()
	t.Run("AddUser", testAddUser)
	t.Run("AddComments", testAddComments)
	t.Run("ListComments", testListComments)
}

func testAddComments(t *testing.T) {
	vid := "12345"
	aid := 1
	content := "I like this video"
	err := AddNewComments(vid, aid, content)
	if err != nil {
		t.Errorf("Error of AddCommentss: %v", err)
	}
}

func testListComments(t *testing.T) {
	vid := "12345"
	from := 1514764800
	to, _ := strconv.Atoi(strconv.FormatInt(time.Now().UnixNano()/1000000000, 10))

	res, err := ListComments(vid, from, to)
	if err != nil {
		t.Errorf("Error of ListComments: %v", err)
	}
	for i, ele := range res {
		fmt.Printf("comment: %d, %v \n", i, ele)
	}

}
