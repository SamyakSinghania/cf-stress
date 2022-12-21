package main

const (
	// ErrNotFound means the definition could not be found for the given word
	ErrNotFound = DictionaryErr("could not find the word you were looking for")

	// ErrWordExists means you are trying to add a word that is already known
	ErrWordExists = DictionaryErr("cannot add word because it already exists")

	// ErrWordDoesNotExist occurs when trying to update a word not in the dictionary
	ErrWordDoesNotExist = DictionaryErr("cannot update word because it does not exist")
)

// DictionaryErr are errors that can happen when interacting with the dictionary.
// Notice that this is not a "struct" data type, but a "string" data type.
// This means that you can extend existing data types in Go, to add new functionality.
// We will see an application of this in our project while using
// http.Header, which is an extension of the map data type.
// Place your cursor on the "Header" comment to learn more about it.
//
// You need to figure out how to make this data type behave like an error as well,
// and return the 3 different errors as mentioned above, wherever applicable.
type DictionaryErr string

func (errorType DictionaryErr) Error() string {
	return string(errorType)
}

// Dictionary store definitions to words.
// We are using the same trick used by http.Header, that is, we create a custom
// data type, which gives us all the functionalities of maps PLUS any extra methods
// that we'd like to add onto this type.
type Dictionary map[string]string

// Search find a word in the dictionary.
// It should return the predefined errors, if any.
func (d Dictionary) Search(word string) (string, error) {
	str, isThere := d[word]
	err := error(nil)
	if !isThere {
		err = ErrNotFound
	}
	return str, err
}

// Add inserts a word and definition into the dictionary.
// Use the d.Search method instead of doing a raw query on the underlying map.
// This will help you to handle each error separately.
func (d Dictionary) Add(word, definition string) error {
	_, err := d.Search(word)
	if err == ErrNotFound {
		d[word] = definition
		return nil
	}
	err = ErrWordExists
	return err

}

// Update changes the definition of a given word.
// Use the d.Search method instead of doing a raw query on the underlying map.
// This will help you to handle each error separately.
func (d Dictionary) Update(word, definition string) error {
	_, err := d.Search(word)
	if err == ErrNotFound {
		err = ErrWordDoesNotExist
		return err
	}
	d[word] = definition
	return nil
}

// Delete removes a word from the dictionary.
func (d Dictionary) Delete(word string) {
	delete(d, word)
}

func main() {

}
