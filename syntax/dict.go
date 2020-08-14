package syntax

import "errors"

// Dictionary type
type Dictionary map[string]string

var errNotFound = errors.New("Not found")
var errWordExist = errors.New("Word exist")
var errCannotUpdate = errors.New("Word not exist")

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

// Update word of definition
func (d Dictionary) Update(word, def string) error {
	_, err := d.Search(word)

	switch err {
	case nil:
		d[word] = def
	case errNotFound:
		return errCannotUpdate
	}

	return nil
}

// Delete word of definition
func (d Dictionary) Delete(word string) {
	delete(d, word)
}
