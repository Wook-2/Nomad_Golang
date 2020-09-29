package mydict

import "errors"

type Dictionary map[string]string

var errNotFound = errors.New("Not Found")
var errExists = errors.New("That word already exists")

//Search word from dictionary
func (d Dictionary) Search(word string) (string, error) {
	value, exist := d[word]
	if exist {
		return value, nil
	}
	return "", errNotFound
}

// Add word to dictionary
func (d Dictionary) Add(word, def string) error {
	_, err := d.Search(word)

	switch err {
	case errNotFound:
		d[word] = def
	case nil:
		return errExists
	}
	return nil
}

// Delete word from dictionary
func (d Dictionary) Delete(word string) {
	delete(d, word)
}
