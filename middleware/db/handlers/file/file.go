package file

import (
	"errors"
	"os"
)

//common errors between functions
var errOpen error = errors.New("file could not be opened")
var errCreate error = errors.New("file could not be created")
var errWrite error = errors.New("file could not be written to")
var errRead error = errors.New("file could not be read from")

//create ...
//creates the file to be used to dump the contents of the content byte array in
func create(string filepath) error {
	var _, err = os.Stat(filepath)

	if os.IsNotExist(err) {
		var file, err = os.Create(path)
		if err != nil {
			return errCreate
		}
		defer file.Close()
	}
	return err
}

//write ...
//writes the contents of the content byte array to the file
func write(filepath string, content []byte) error {
	file, err := os.OpenFile(filepath, os.O_RDWR, 0644)
	if err != nil {
		return errOpen
	}
	defer file.Close()

	file.Write(content)

	if err != nil {
		return errWrite
	}

	err = file.Sync()
	if err != nil {
		return errors.New("change could not be saved")
	}

	return err
}

//read ...
//allows the file to be read from
func read(filepath string) error {
	file, err := os.OpenFile(filepath, os.O_RDWR, 0644)
	if err != nil {
		return errOpen
	}
	defer file.Close()

	var content []byte

	for{
		_,err = file.Read(content)

		if err = io.EOF{break}

		if err != nil && err != io.EOF{
			if err != nil {
				return err
			}
			break
		}
	}
	return err
}

//delete ...
//in the event a file must be deleted
func delete(filepath string) error{
	err := os.Remove(filepath)

	if err != nil {
		return err
	}


}