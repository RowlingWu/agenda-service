package entities

// curuser's(current user) most functions are similar to userlist
// used to do log in/out work
type CurUser struct {
  Username string `gorm:"primary_key;column:username"`
}

func (*CurUser) GroupName() string {
  return "current_user"
}

type CuruserServer struct{}
var CurServ = CuruserServer{}

func init() {
  addServ(&CurServ)
}

func (*CuruserServer) load() {
  c := &CurUser{}
  if !DB.HasTable(c) {
    DB.CreateTable(c)
  }
}

func (*CuruserServer) MyAdd(curuser *CurUser) {
  temp := DB.Begin()
  checkError(temp.Error)
  if err := temp.Create(curuser).Error; err!=nil {
    temp.Rollback()
    checkError(err)
  }
  temp.Commit()
}

func (*CuruserServer) MyDelete(duser string) {
  cur := &CurUser{Username: duser}
  temp := DB.Begin()
  if err := temp.Delete(cur).Error; err!=nil {
    temp.Rollback()
    checkError(err)
  }
  temp.Commit()
}
func (*CuruserServer) CurListAll() []CurUser {
  userlist := make([]CurUser,0,0)
  checkError(DB.Find(&userlist).Error)
  return userlist
}

func (*CuruserServer) CurQuery(username string) *CurUser {
  userlist := make([]CurUser,0,0)
  checkError(DB.Where([]string{username}).Find(&userlist).Error)
  if len(userlist) == 0 {
    //log.Fatal("The name: %s has been used.",u)
    return nil
  }
  return &userlist[0]
}
