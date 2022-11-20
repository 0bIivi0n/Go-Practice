package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println()
	fmt.Println("**************************************************************")
	fmt.Println("*                                                            *")
	fmt.Println("*       *** Bienvenue dans le jeu du plus ou moins ***       *")
	fmt.Println("*   retrouvez le chiffre mystère à l'aide des indications.   *")
	fmt.Println("*                                                            *")
	fmt.Println("**************************************************************")
	fmt.Println("")

	PlusOuMoins()

	fmt.Println("Au revoir!")

}

func PlusOuMoins() {

	rand.Seed(time.Now().UnixNano()) // Initialise la source de rand selon l'heure (meilleur RNG)
	rand := rand.Intn(100)           // génère un int aléatoire entre 1 et 100
	var userInput int
	playAgain := ""

	for userInput != rand {
		fmt.Println("Indiquez un chiffre entre 1 et 100: ")
		fmt.Scanln(&userInput) // Scanln récupère l'entrée console et la stoque dans userInput

		if userInput < rand {
			fmt.Println("c'est plus!")
			fmt.Println()
		} else {
			fmt.Println("c'est moins!")
			fmt.Println()
		}
	}

	// Message réussite
	fmt.Printf("Bravo! Vous avez trouvé le chiffre mystère: %v \n", rand)
	fmt.Println()

	// Nouvelle partie
	for playAgain != "y" && playAgain != "n" {
		fmt.Println("rejouer ? (y/n): ")
		fmt.Scanln(&playAgain)
		if playAgain == "y" {
			fmt.Printf("\x1bc")
			PlusOuMoins()
		} else if playAgain != "n" {
			fmt.Println("KesTuDi ?")
			fmt.Println()
		}
	}
	return
}
