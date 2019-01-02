package main

import "testing"

func TestHello(t *testing.T) {

	//define a function within a function by assigning it to a variable
	//you can then call this variable later by giving it args
	assertCorrectMessage := func(t *testing.T, got, want string) {
        t.Helper()
        if got != want {
            t.Errorf("got '%s' want '%s'", got, want)
        }
	}
	
	t.Run("saying hello world", func(t *testing.T) {
    	got := Hello()
		want := "Hello, World"
	
		assertCorrectMessage(t, got, want)
	})

	t.Run("saying hello to people", func(t *testing.T) {
    	got := HelloYou("Alice")
		want := "Hello, Alice"
	
		assertCorrectMessage(t, got, want)
	})

	t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
        got := HelloYou("")
		want := "Hello, World"

		assertCorrectMessage(t, got, want)
    })
}