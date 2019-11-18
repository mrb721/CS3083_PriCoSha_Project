package tables

//Person ...
// Represents the person table in mysql
type Person struct {
	Username       string
	HashedPassword string
	Salt           string
	Fname          string
	Lname          string
	ColorPalette   string
}
