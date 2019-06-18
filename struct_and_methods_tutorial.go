package main


import (
    "fmt"
    "strconv"
)
type User struct{
    firstname, lastname string
    age int
}

func (u *User) fullName() string {
    return u.firstname + " " + u.lastname
}

func (u *User) getDetailString() string {
    return u.firstname + " " + u.lastname + " has " + strconv.Itoa(u.age) + " year of age"
}

func main() {

    u := User{"Jules", "Verne", 40}

    fmt.Println(u.fullName())
    fmt.Println(u.getDetailString())
    
}