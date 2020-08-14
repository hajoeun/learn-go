package syntax

import "errors"

// Dictionary type
type Dictionary map[string]string

var errNotFound = errors.New("Not Found")

// Search find value from Dictionary
func (d Dictionary) Search(word string) (string, error) {
	value, ok := d[word]

	if ok {
		return value, nil
	}

	return word, errNotFound
}
