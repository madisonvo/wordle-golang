package wordle

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

// This creates the WordFactory structure
type WordFactory struct {
	wordList []string
}

// This Function reads 'Words.txt' to get all 5 letter words.
func NewWordFactory() *WordFactory {
	wordList, err := readLines("words.txt")
	if err != nil {
		fmt.Println("Error reading words.txt:", err)
	}

	return &WordFactory{wordList: wordList}
}

// This function creates a list of all 5 letter words
func readLines(filename string) ([]string, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	lines := strings.Split(string(data), "\n")

	var result []string

	for i := 0; i < len(lines); i++ {
		line := lines[i]
		word := strings.TrimSpace(line)
		if word != "" {
			result = append(result, word)
		}
	}

	return result, nil
}

// This gets a random word each time the method is called.
func (wf *WordFactory) MakeWord() string {
	rand.NewSource(time.Now().UnixNano())
	randomIndex := rand.Intn(len(wf.wordList))
	word := wf.wordList[randomIndex]
	fmt.Println("This is the word you're trying to guess: " + word)
	return word
}

// Wordle is responsible for the operation of the Wordle game.
type Wordle struct {
	wordList  *WordFactory
	word      string
	guesses   int
	result    string
	numLetter map[string]int
}

// This creates a new instance of the wordle game.
func NewGame() *Wordle {
	wl := NewWordFactory()
	word := wl.MakeWord()

	numLetter := make(map[string]int)
	for i := 0; i < len(word); i++ {
		letter := string(strings.ToLower(string(word[i])))
		numLetter[letter]++
	}

	return &Wordle{
		wordList:  wl,
		word:      word,
		guesses:   5,
		result:    "-----",
		numLetter: numLetter,
	}
}

// This function is responsible for making the guesses
func (w *Wordle) MakeGuess(input string) (bool, string) {
	fmt.Println("This was your guess: " + input)
	if w.guesses > 0 {
		w.guesses -= 1
		fmt.Println("This is how many guesses you have left: " + strconv.Itoa(w.guesses))
		if strings.ToLower(input) == w.word {
			w.result = input
			fmt.Println("You guessed " + w.result + "!")
			guesses := 5 - w.guesses
			message := ""
			if guesses == 1 {
				message = "Victory! You guessed in " + strconv.Itoa(guesses) + " turn."
			} else {
				message = "Victory! You guessed in " + strconv.Itoa(guesses) + " turns."
			}

			return true, message
		} else {
			message := w.checkCharacters(input)
			return false, message
		}
	} else if w.guesses <= 0 {
		fmt.Println("Out of guesses :(")
		message := "Out of guesses: :("
		return false, message
	}

	return false, ""
}

func (w *Wordle) checkCharacters(input string) string {
	wrong := w.redIndex(input)
	semiCorrect := w.yellowIndex(input)
	correct := w.greenIndex(input)
	message := ""

	if atLeastOneTrue(wrong) {
		for i, boolean := range wrong {
			if boolean {
				message += string(input[i]) + " "
			}
		}

		message += "are the wrong characters\n"
	}

	if atLeastOneTrue(semiCorrect) {
		for i, boolean := range semiCorrect {
			if boolean {
				message += string(input[i]) + " "
			}
		}

		message += "are not in the correct place\n"
	}

	if atLeastOneTrue(correct) {
		for j, boolean := range correct {
			if boolean {
				message += string(input[j]) + " "
			}
		}

		message += "are in the correct place"
	}

	return message
}

func (w *Wordle) redIndex(input string) []bool {
	var notInWord []bool
	for i := 0; i < 5; i++ {
		if !w.inWord(string(input[i])) || w.numLetter[string(input[i])] <= 0 {
			notInWord = append(notInWord, true)
		} else {
			notInWord = append(notInWord, false)
		}
	}
	return notInWord

}

func (w *Wordle) greenIndex(input string) []bool {
	var correctPlace []bool
	for i := 0; i < 5; i++ {
		if strings.ToLower(string(input[i])) == string(w.word[i]) {
			correctPlace = append(correctPlace, true)
		} else {
			correctPlace = append(correctPlace, false)
		}
	}
	return correctPlace
}

func (w *Wordle) yellowIndex(input string) []bool {
	var yellowSpaces []bool
	for i := 0; i < 5; i++ {
		if w.inWord(string(input[i])) && strings.ToLower(string(input[i])) != string(w.word[i]) {
			yellowSpaces = append(yellowSpaces, true)
		} else {
			yellowSpaces = append(yellowSpaces, false)
		}
	}
	return yellowSpaces
}

func (w *Wordle) inWord(letter string) bool {
	for i := 0; i < 5; i++ {
		if letter == string(w.word[i]) {
			return true
		}
	}
	return false
}

func atLeastOneTrue(boolList []bool) bool {
	for _, boolean := range boolList {
		if boolean {
			return true
		}
	}

	return false
}
