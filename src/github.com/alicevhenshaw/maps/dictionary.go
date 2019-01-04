package dictionary

import "errors"

type Dictionary map[string]string

var ErrorWordNotInDict = errors.New("Searched word not in the dictionary")
var ErrorWordAlreadyDefined = errors.New("Trying to redefine a word")
var ErrorWordNotDefined = errors.New("Word not yet defined")

func (d Dictionary) Search(word string) (string, error) {
	// a map CAN return a second argument, a boolean, that's true when the entry has been created
	// although this makes little difference in this case with a dictionary, the difference between
	// no entry, and an entry of "" may be crucial in some settings
	definition, isDefined := d[word]

	if !isDefined {
		return "", ErrorWordNotInDict
	}
	return definition, nil
}

func (d Dictionary) checkDefined(word string) bool {
	_, isDefined := d[word]
	return isDefined
}

func (d Dictionary) Add(word, definition string) error {
	if d.checkDefined(word) {
		return ErrorWordAlreadyDefined
	}
	d[word] = definition
	return nil
}

func (d Dictionary) Update(word, definition string) error {
	if !d.checkDefined(word) {
		return ErrorWordNotDefined
	}
	d[word] = definition
	return nil
}

func (d Dictionary) Delete(word string) error {
	if !d.checkDefined(word) {
		return ErrorWordNotDefined
	}
	// Delete the element properly - this also reset the boolean to false
	delete(d, word)
	return nil
}