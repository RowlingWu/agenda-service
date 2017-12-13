package entity

// User is used as a format in sqlite's table.
type User struct {
	// ID should BE a primary key, auto-increment, not null
	Id int
	// NAME is not null
	Name string
	// PASSWD is not null
	Passwd string
	Email  string
	Tel    string
}

// check whether user has logged in
/*
var (
	// CurUser is the path where curUser.txt is stored
	CurUser string
)

func init() {
	u, err := user.Current()
	checkError(err)
	var homeDir = u.HomeDir
	CurUser = homeDir + "/curUser.txt"

	if _, err = os.Stat(CurUser); err != nil {
		f, err := os.Create(CurUser)
		defer f.Close()
		checkError(err)
	}
}
*/

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
