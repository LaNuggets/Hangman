package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

var linesRead int // Variable pour suivre le nombre de lignes lues

// readWordsFromFile read words in the words package
func readWordsFromFile(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	var words []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		word := strings.TrimSpace(scanner.Text())
		if word != "" {
			words = append(words, word)
		}
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return words, nil
}

func main() {
	// Reading words from the "words.txt" file
	words, err := readWordsFromFile("words.txt")
	if err != nil {
		fmt.Println("Erreur lors du chargement du fichier")
		return
	}
	// Initialize random number
	rand.Seed(time.Now().UnixNano())
	// Choose a random word
	randomIndex := rand.Intn(len(words))
	randomWord := words[randomIndex]
	// Create a hidden word
	hiddenWord := createHiddenWord(randomWord)
	// Combine the random word with the hidden word for display
	displayWord := mergeWords(randomWord, hiddenWord)
	// Displays the chosen random word with the revealed letters
	fmt.Printf("Mot aléatoire choisi : %s\n", displayWord)
	// Count the number of characters
	fmt.Printf("Nombre de caractères dans le mot d'origine : %d\n", len(randomWord))
	// Play hangman with 10 attempts
	playHangman(randomWord, hiddenWord, displayWord, 10)
}

// createHiddenWord generates a hidden word by showing a random part of the original word's letters.
func createHiddenWord(word string) string {
	revealedCount := len(word)/2 - 1
	hiddenWord := make([]rune, len(word))
	// Créez un tableau d'indices pour les caractères révélés
	revealedIndices := make([]int, revealedCount)
	for i := 0; i < revealedCount; i++ {
		index := rand.Intn(len(word))
		revealedIndices[i] = index
		hiddenWord[index] = rune(word[index])
	}
	// Remplissez le reste avec des caractères "_"
	for i := range hiddenWord {
		if hiddenWord[i] == 0 {
			hiddenWord[i] = '_'
		}
	}
	return string(hiddenWord)
}

// Merge words character by character
func mergeWords(randomWord string, hiddenWord string) string {
	mergedWord := ""
	for i := 0; i < len(randomWord); i++ {
		if hiddenWord[i] == '_' {
			mergedWord += "_"
		} else {
			mergedWord += string(randomWord[i])
		}
	}
	return mergedWord
}

func playHangman(word string, hiddenWord string, displayWord string, maxAttempts int) {
	guesses := make([]rune, len(word))
	for i := range guesses {
		guesses[i] = '_'
	}
	attempts := maxAttempts
	usedLetters := make(map[rune]bool)
	for attempts > 0 {
		fmt.Printf("Mot caché : %s\n", displayWord)
		fmt.Printf("Tentatives restantes : %d\n", attempts)
		fmt.Print("Lettres déjà utilisées : ")
		printUsedLetters(usedLetters)
		var guess string
		fmt.Print("Devinez une lettre (ou tapez '/' pour quitter) : ")
		fmt.Scan(&guess)
		if guess == "/" {
			fmt.Println("Vous avez quitté le jeu.")
			return
		}
		if len(guess) != 1 || !isLetter([]rune(guess)[0]) {
			fmt.Println("Veuillez entrer une seule lettre valide.")
			continue
		}
		runeGuess := []rune(guess)[0]
		// checks if the letter is already used
		if usedLetters[runeGuess] {
			fmt.Printf("La lettre %c a déjà été utilisée.\n", runeGuess)
			continue
		}
		usedLetters[runeGuess] = true
		found := false
		for i, letter := range word {
			if displayWord[i] == '_' && letter == runeGuess {
				displayWord = displayWord[:i] + string(runeGuess) + displayWord[i+1:]
				found = true
			}
		}
		if !found {
			attempts--
			fmt.Println("Lettre incorrecte. Essayez à nouveau.")
			// design()
		}
		if displayWord == word {
			fmt.Printf("Bravo ! Vous avez deviné le mot : %s\n", word)
			break
		}
	}
	if displayWord != word {
		fmt.Printf("Game Over, le mot était : %s\n", word)
	}
}
func isLetter(r rune) bool {
	return (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z')
}

func printUsedLetters(usedLetters map[rune]bool) {
	for letter, used := range usedLetters {
		if used {
			fmt.Printf("%c ", letter)
		}
	}
	fmt.Println()
}

/* func design() {

	// Ouvrir le fichier
	file, error := os.Open("hangman.txt")
	if error != nil {
		fmt.Println("Erreur lors de l'ouverture du fichier :", error)
		return
	}
	defer file.Close()

	// Créer un lecteur de fichiers
	fileReader := bufio.NewReader(file)

	// Sauter les lignes déjà lues
	for i := 0; i < linesRead; i++ {
		_, error := fileReader.ReadString('\n')
		if error != nil {
			break // Arrêter la boucle si nous atteignons la fin du fichier
		}
	}

	// Lire et afficher les lignes
	for i := 0; i < 8; i++ {
		line, error := fileReader.ReadString('\n')
		if error != nil {
			break // Arrêter la boucle si nous atteignons la fin du fichier
		}
		fmt.Print(line)
	}

	// Mettre à jour le nombre de lignes lues
	linesRead += 8
}
*/
