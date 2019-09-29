package dbops

import "testing"

/*
init(db login, truncate tables) -> run tests ->clear data
*/

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
