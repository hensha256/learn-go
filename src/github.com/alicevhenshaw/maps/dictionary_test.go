package dictionary

import "testing"

func TestSearch(t *testing.T) {
	// map[keyType]valueType
    dictionary := Dictionary {
		"test": "this is just a test",
	}

	t.Run("Word exists in dict", func(t *testing.T) {
		definition, err := dictionary.Search("test")
		want := "this is just a test"

		if err != nil {
			t.Errorf("Got an error when none was expected: %s", err)
		}
		assertStrings(t, definition, want)
	})

	t.Run("Word doesn't exist in dict", func(t *testing.T) {
		definition, err := dictionary.Search("unknown")
		want := ""

		if err == nil {
			t.Errorf("Got no error when one was expected")
		}
		// .Error() converts the error to a string
		assertStrings(t, err.Error(), ErrorWordNotInDict.Error())
		assertStrings(t, definition, want)
	})
}

func TestAdd(t *testing.T) {
	// map[keyType]valueType
    dictionary := Dictionary {
		"test": "this is just a test",
	}

	addErr := dictionary.Add("cat", "is not a dog")
	if addErr != nil {
		t.Errorf("Error: %s", addErr)
	}

	definition, _ := dictionary.Search("cat")
	assertStrings(t, definition, "is not a dog")
}

func TestUpdate(t *testing.T) {
	// map[keyType]valueType
    dictionary := Dictionary {
		"test": "this is just a test",
	}

	dictionary.Add("cat", "is not a dog")
	dictionary.Update("cat", "is not a dog or bat")

	definition, _ := dictionary.Search("cat")
	assertStrings(t, "is not a dog or bat", definition)
}

func assertStrings(t *testing.T, got, want string) {
	t.Helper()

	if got != want {
        t.Errorf("got '%s' want '%s' given", got, want)
    }
}