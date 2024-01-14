package wordle

import (
	"testing"
)

func TestNewWordFactory(t *testing.T) {
	// Valid file
	wf := NewWordFactory()
	if len(wf.wordList) == 0 {
		t.Errorf("Expected wordList to have some words, but it is empty")
	}
}

func TestMakeWord(t *testing.T) {
	wf := NewWordFactory()

	word1 := wf.MakeWord()
	word2 := wf.MakeWord()
	if word1 == word2 {
		t.Errorf("Expected MakeWord to return different words on each call, but got the same word")
	}

	if len(word1) != 5 {
		t.Errorf("Expected the returned word to have 5 characters, but got %d characters", len(word1))
	}
}

func TestNewGame(t *testing.T) {
	// New game
	game1 := NewGame()
	if game1.guesses != 6 {
		t.Errorf("Expected initial number of guesses to be 6, but got %d", game1.guesses)
	}

	if len(game1.result) != 5 {
		t.Errorf("Expected initial result to be a five-character string, but got %s", game1.result)
	}

	// Multiple games
	game2 := NewGame()
	if game1 == game2 {
		t.Errorf("Expected NewGame to create different game instances, but got the same instance")
	}
}

func TestMakeGuess(t *testing.T) {
	game := NewGame()

	// Correct guess
	correctGuess := game.word
	if result1, _ := game.MakeGuess(correctGuess); !result1 {
		t.Errorf("Expected correct guess to return true, but got false")
	}
	if game.result != correctGuess {
		t.Errorf("Expected game result to be %s, but got %s", correctGuess, game.result)
	}

	// Incorrect guess
	incorrectGuess := "ABCDE"
	if result2, _ := game.MakeGuess(incorrectGuess); result2 {
		t.Errorf("Expected incorrect guess to return false, but got true")
	}
	if game.result == incorrectGuess {
		t.Errorf("Expected game result not to be %s, but it is", incorrectGuess)
	}

	// Running out of guesses
	game = NewGame()
	for i := 0; i < 6; i++ {
		game.MakeGuess("XXXXX")
	}
	result3, message1 := game.MakeGuess("XXXXX")
	if result3 {
		t.Errorf("Expected out of guesses to return false, but got true")
	}
	if message1 != "Out of guesses" {
		t.Errorf("Expected out of guesses message to be 'Out of guesses', but got %s", message1)
	}
}

func TestCheckCharacters(t *testing.T) {
	game := NewGame()

	// Correct guess
	correctGuess := game.word
	message := game.checkCharacters(correctGuess)
	if message != "are in the correct place" {
		t.Errorf("Expected correct guess message to be 'are in the correct place', but got %s", message)
	}

	// Incorrect guess with some correct letters
	incorrectGuess := "CDEFG"
	message = game.checkCharacters(incorrectGuess)
	if message != "A B are in the correct place" {
		t.Errorf("Expected incorrect guess message to be 'A B are in the correct place', but got %s", message)
	}

	// Incorrect guess with all incorrect letters
	incorrectGuess = "XYZ"
	message = game.checkCharacters(incorrectGuess)
	if message != "X Y Z are the wrong characters" {
		t.Errorf("Expected incorrect guess message to be 'X Y Z are the wrong characters', but got %s", message)
	}
}

func TestRedIndex(t *testing.T) {
	game := NewGame()

	// Correct guess
	correctGuess := game.word
	redIndexes := game.redIndex(correctGuess)
	for i, isRed := range redIndexes {
		if isRed {
			t.Errorf("Expected all indexes to be not red for a correct guess, but index %d is red", i)
		}
	}

	// Incorrect guess
	incorrectGuess := "ABCDE"
	redIndexes = game.redIndex(incorrectGuess)
	for i, isRed := range redIndexes {
		if !isRed {
			t.Errorf("Expected all indexes to be red for an incorrect guess, but index %d is not red", i)
		}
	}
}

func TestGreenIndex(t *testing.T) {
	game := NewGame()

	// Correct guess
	correctGuess := game.word
	greenIndexes := game.greenIndex(correctGuess)
	for i, isGreen := range greenIndexes {
		if !isGreen {
			t.Errorf("Expected all indexes to be green for a correct guess, but index %d is not green", i)
		}
	}

	// Incorrect guess
	incorrectGuess := "ABCDE"
	greenIndexes = game.greenIndex(incorrectGuess)
	for i, isGreen := range greenIndexes {
		if isGreen {
			t.Errorf("Expected all indexes to be not green for an incorrect guess, but index %d is green", i)
		}
	}
}

func TestYellowIndex(t *testing.T) {
	game := NewGame()

	// Correct guess
	correctGuess := game.word
	yellowIndexes := game.yellowIndex(correctGuess)
	for i, isYellow := range yellowIndexes {
		if isYellow {
			t.Errorf("Expected all indexes to be not yellow for a correct guess, but index %d is yellow", i)
		}
	}

	// Incorrect guess with some correct letters but not in the correct place
	incorrectGuess := "CABDF"
	yellowIndexes = game.yellowIndex(incorrectGuess)
	for i, isYellow := range yellowIndexes {
		if !isYellow {
			t.Errorf("Expected all indexes to be yellow for an incorrect guess with some correct letters not in the correct place, but index %d is not yellow", i)
		}
	}

	// Incorrect guess with all incorrect letters
	incorrectGuess = "XYZ"
	yellowIndexes = game.yellowIndex(incorrectGuess)
	for i, isYellow := range yellowIndexes {
		if isYellow {
			t.Errorf("Expected all indexes to be not yellow for an incorrect guess with all incorrect letters, but index %d is yellow", i)
		}
	}
}

func TestInWord(t *testing.T) {
	game := NewGame()

	// Letter that is in the word
	letter := "A"
	if !game.inWord(letter) {
		t.Errorf("Expected letter '%s' to be in the word, but it's not", letter)
	}

	// Letter that is not in the word
	letter = "X"
	if game.inWord(letter) {
		t.Errorf("Expected letter '%s' not to be in the word, but it is", letter)
	}
}

func TestAtLeastOneTrue(t *testing.T) {
	// All false values
	boolList := []bool{false, false, false}
	if atLeastOneTrue(boolList) {
		t.Errorf("Expected atLeastOneTrue to return false for all false values, but it returned true")
	}

	// All true values
	boolList = []bool{true, true, true}
	if !atLeastOneTrue(boolList) {
		t.Errorf("Expected atLeastOneTrue to return true for all true values, but it returned false")
	}

	// Both true and false values
	boolList = []bool{false, true, false}
	if !atLeastOneTrue(boolList) {
		t.Errorf("Expected atLeastOneTrue to return true for a mix of true and false values, but it returned false")
	}
}
