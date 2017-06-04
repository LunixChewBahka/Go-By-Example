package main

type User struct {
	Name string
}

func main() {
	u := &User{Name: "Leto"}
	println(u.Name)
	modify(u)
	println(u.Name)
}

func modify(u *User) {
	*u = User{Name: "Jarelim"}
}
