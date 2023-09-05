package main

import (
	"reflect"
	"testing"
)

func Test_getTodaysWord(t *testing.T) {
	type args struct {
		wordlist []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"produces_a_word", args{[]string{"drums"}}, "drums"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getTodaysWord(tt.args.wordlist); got != tt.want {
				t.Errorf("getTodaysWord() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_loadWordlist(t *testing.T) {
	type args struct {
		wordListPath string
	}
	tests := []struct {
		name          string
		args          args
		wantWordCount int
	}{
		{"fetches_the_correct_number_of_words_from_wordlist", args{WORD_LIST_PATH}, 14855},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotWords := loadWordlist(tt.args.wordListPath); len(gotWords) != tt.wantWordCount {
				t.Errorf("loadWordlist() = count=%v, want count=%v", len(gotWords), tt.wantWordCount)
			}
		})
	}
}

func Test_wordIsValid(t *testing.T) {
	wordlist := loadWordlist(WORD_LIST_PATH)
	type args struct {
		word     string
		wordlist []string
	}
	tests := []struct {
		name      string
		args      args
		wantValid bool
	}{
		{"word_valid", args{"drums", wordlist}, true},
		{"word_too_long", args{"invalid", wordlist}, false},
		{"word_not_in_dictionary", args{"aaaaa", wordlist}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotValid := wordIsValid(tt.args.word, tt.args.wordlist); gotValid != tt.wantValid {
				t.Errorf("wordIsValid() = %v, want %v", gotValid, tt.wantValid)
			}
		})
	}
}

func Test_compareWordLetters(t *testing.T) {
	type args struct {
		userWord string
		answer   string
	}
	tests := []struct {
		name       string
		args       args
		wantResult []Letter
	}{
		{"word_correct_all", args{"drums", "drums"}, []Letter{CORRECT_LETTER, CORRECT_LETTER, CORRECT_LETTER, CORRECT_LETTER, CORRECT_LETTER}},
		{"word_incorrect_some", args{"drool", "drums"}, []Letter{CORRECT_LETTER, CORRECT_LETTER, INCORRECT_LETTER, INCORRECT_LETTER, INCORRECT_LETTER}},
		{"word_present_letters", args{"darts", "drums"}, []Letter{CORRECT_LETTER, INCORRECT_LETTER, PRESENT_LETTER, INCORRECT_LETTER, CORRECT_LETTER}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotResult := compareWordLetters(tt.args.userWord, tt.args.answer); !reflect.DeepEqual(gotResult, tt.wantResult) {
				t.Errorf("validateWordLetters() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}

func Test_wordIsCorrect(t *testing.T) {
	type args struct {
		result []Letter
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"word_correct", args{[]Letter{CORRECT_LETTER, CORRECT_LETTER, CORRECT_LETTER, CORRECT_LETTER, CORRECT_LETTER}}, true},
		{"word_incorrect", args{[]Letter{INCORRECT_LETTER, CORRECT_LETTER, CORRECT_LETTER, CORRECT_LETTER, CORRECT_LETTER}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := wordIsCorrect(tt.args.result); got != tt.want {
				t.Errorf("wordIsCorrect() = %v, want %v", got, tt.want)
			}
		})
	}
}
