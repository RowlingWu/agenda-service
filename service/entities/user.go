package entities

type User struct {
  // Id should BE a primary key, auto-increment, not null
	Id int `json:"Id" gorm:"column:Id"`
	// NAME is not null
	Name string `json:"Name" gorm:"primary_key;column:Name"`
	// PASSWD is not null
	Passwd string `json:"Passwd" gorm:"column:Passwd"`
	Email  string `json:"Email" gorm:"column:Email"`
	Tel    string `json:"Tel" gorm:"column:Tel"`
}

func (*User) GroupName() string {
  return "user"
}

type UserServer struct{}
var UserServ = UserServer{}

func init() {
  addServ(&UserServ)
}

func (*UserServer) load() {
  u := &User{}
  if !DB.HasTable(u) {
    DB.CreateTable(u)
  }
}

func (*UserServer) MyAdd(newuser *User) {
  temp := DB.Begin()
  checkError(temp.Error)
  if err := temp.Create(newuser).Error; err!=nil {
    temp.Rollback()
    checkError(err)
  }
  temp.Commit()
}

func (*UserServer) MyDelete(newuser *User) {
  temp := DB.Begin()
  checkError(temp.Error)
  if err := temp.Delete(newuser).Error; err!=nil {
    temp.Rollback()
    checkError(err)
  }
  temp.Commit()
}

func (*UserServer) ListAll() []User {
  userlist := make([]User,0,0)
  checkError(DB.Find(&userlist).Error)
  return userlist
}

func (*UserServer) MyQuery(username string) *User {
  userlist := make([]User,0,0)
  checkError(DB.Where([]string{username}).Find(&userlist).Error)
  if len(userlist) == 0 {
    //log.Fatal("The name: %s has been used.",u)
    return nil
  }
  return &userlist[0]
}
