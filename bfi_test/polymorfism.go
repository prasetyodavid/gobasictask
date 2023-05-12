package main

import (
  "fmt"
)

type user struct {
  name string
  Age int
}

// Contoh interface untuk kelas user
type UserService interface {
  GetUserAge() string
}

func (u user) getName() string {
  return u.name
}

func (u user) GetUserAge() string {
  return fmt.Sprintf("%s is %d years old", u.getName(), u.Age)
}
// Implementasikan interface UserService
func NewUser(name string, age int) UserService {
  return user{name: name, Age: age}
}

func main() {
  user := NewUser("John Lenon", 29)
  fmt.Println(user.GetUserAge())
}

/*

Pada kode di atas, kita membuat interface UserService, yang mana nantinya akan diimplementasikan untuk kelas user. Jadi jika kelas user tersebut tidak memiliki method GetUserAge, maka akan mengembalikan panic error seperti berikut.

./main.go:26:14: cannot use user literal (type user) as type UserService in return argument:
    user does not implement UserService (missing GetUserAge method)
    
*/
