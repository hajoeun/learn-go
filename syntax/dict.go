package syntax

import "errors"

// Dictionary type
type Dictionary map[string]string

var errNotFound = errors.New("Not Found")
var errWordExist = errors.New("Word Exist")

// Search find value from Dictionary
func (d Dictionary) Search(word string) (string, error) {
	value, ok := d[word]

	if ok {
		return value, nil
	}

	return word, errNotFound
}

// Add definition to Dictionary
func (d Dictionary) Add(word, def string) error {
	_, err := d.Search(word)

	switch err {
	case errNotFound:
		d[word] = def
	case nil:
		return errWordExist
	}
	return nil
}
