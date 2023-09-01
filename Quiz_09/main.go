package main

const (
	ReadPermission = 1 << iota
	WritePermission
	ExecutePermission
)

func main() {
	var permissions byte = ReadPermission | WritePermission

	println(ReadPermission)
	println(WritePermission)
	println(ExecutePermission)

	println(permissions&ReadPermission == ReadPermission)
	println(permissions&WritePermission == WritePermission)
	println(permissions&ExecutePermission == ExecutePermission)
}
