package main

import "fmt"

type UserInterface interface {
	doStuff() string
}

type User struct {
	name string
}

func (u *User) doStuff() string {
	return "Hallo, saya adalah seorang " + u.name
}

type Dosen struct {
	*User
}

func NewDosen() UserInterface {
	return &Dosen{
		User: &User{
			name: "Dosen",
		},
	}
}

type Mahasiswa struct {
	*User
}

func NewMahasiswa() UserInterface {
	return &Mahasiswa{
		User: &User{
			name: "Mahasiswa",
		},
	}
}

func CreateUser(name string) UserInterface {
	if name == "Dosen" {
		return NewDosen()
	} else if name == "Mahasiswa" {
		return NewMahasiswa()
	}
	return nil
}

func main() {
	Dosen := CreateUser("Dosen")
	fmt.Println(Dosen.doStuff())
	Mahasiswa := CreateUser("Mahasiswa")
	fmt.Println(Mahasiswa.doStuff())
}
