package main

import (
	"bufio"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"

	"slices"

	"github.com/fatih/color"
)

const MAX_TRIES = 5
const WORD_LENGTH = 5
const WORD_LIST_PATH = "wordlist.txt"

type Letter uint8

const (
	INCORRECT_LETTER Letter = iota
	PRESENT_LETTER
	CORRECT_LETTER
)

func loadWordlist(wordListPath string) (words []string) {
	log.Println("Loading wordlist...")
	file, err := os.Open(wordListPath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		words = append(words, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return
}

func getTodaysWord(wordlist []string) string {
	log.Println("Picking a random word...")
	seed := time.Now().Round(24 * time.Hour).Unix()
	r := rand.New(rand.NewSource(seed))
	word := wordlist[r.Intn(len(wordlist))]
	return strings.ToLower(word)
}

func getUserWord() (word string) {
	log.Printf("Enter a %d letter word:\n", WORD_LENGTH)
	reader := bufio.NewReader(os.Stdin)
	// TODO: loop until input is valid
	var userInput string
	word, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	userInput = strings.ToLower(strings.TrimSpace(word))
	return userInput
}

func wordIsValid(word string, wordlist []string) (valid bool) {
	userWordLength := len(word)
	if userWordLength != WORD_LENGTH {
		log.Printf("Invalid word length: %d\n", userWordLength)
		return false
	}
	if !slices.Contains(wordlist, word) {
		log.Printf("Not a valid word: %q\n", word)
		return false
	}
	return true
}

func compareWordLetters(userWord, answer string) (result []Letter) {
	result = make([]Letter, len(answer))
	for pos, char := range userWord {
		i := strings.IndexRune(answer, char)
		if i == -1 {
			result[pos] = INCORRECT_LETTER
		} else if i == pos {
			result[pos] = CORRECT_LETTER
		} else {
			result[pos] = PRESENT_LETTER
		}
	}
	return
}

func printValidatedWord(userWord string, letters []Letter) {
	correct := color.New(color.FgHiWhite).Add(color.BgGreen).SprintFunc()
	present := color.New(color.FgHiWhite).Add(color.BgYellow).SprintFunc()
	incorrect := color.New(color.FgHiWhite).Add(color.BgHiBlack).SprintFunc()
	var result string
	for pos, val := range letters {
		currentLetter := string(userWord[pos])
		if val == CORRECT_LETTER {
			result += correct(currentLetter)
		} else if val == PRESENT_LETTER {
			result += present(currentLetter)
		} else {
			result += incorrect(currentLetter)
		}
	}
	log.Println(result)
}

func wordIsCorrect(result []Letter) bool {
	for _, val := range result {
		if val != CORRECT_LETTER {
			return false
		}
	}
	return true
}

func main() {
	log.Println("Welcome to GORDLE")
	triesRemaining := MAX_TRIES
	wordlist := loadWordlist(WORD_LIST_PATH)
	answer := getTodaysWord(wordlist)
	for triesRemaining > 0 {
		userInput := getUserWord()
		if wordIsValid(userInput, wordlist) {
			triesRemaining -= 1
			correctLetters := compareWordLetters(userInput, answer)
			printValidatedWord(userInput, correctLetters)
			if wordIsCorrect(correctLetters) {
				log.Println("Congratulations, you've guessed the correct word!")
				break
			} else if triesRemaining == 0 {
				log.Println("Sorry, no more tries.")
			} else {
				log.Printf("Sorry, try again. %d tries remaining", triesRemaining)
			}
		}
	}
}
