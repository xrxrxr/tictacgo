package main

import (
    "bufio" // https://golang.org/pkg/bufio/
    "fmt"  // https://golang.org/pkg/fmt/
    "os"  // https://golang.org/pkg/os/
    "strconv" // https://golang.org/pkg/strconv/
)

// constantes globales
const (
    tailleDamier   = 9
    symboleJoueur1 = "X"
    symboleJoueur2 = "O"
)

// Variables globales
var (
    tableauMorpion = [tailleDamier]string{ // création d'un damier sans aucune cases remplies
        "1", "2", "3",
        "4", "5", "6",
        "7", "8", "9"}
    joueur1 = true // c'est le joueur 1 qui commence en 1er
)

func main() {
    jouer() // Lancement du jeu
}

/**
* description de la fonction : Permet de lancer le jeu
*
* @return rien
*/
func jouer() {

    var numeroCase int
    for true {
        affichage()
        numeroCase = gestionEntreeUtilisateur() // Récupération de l'entrée utilisateur
        remplirCase(numeroCase)
        if gagner() { // Vérifier si le joueur a gagné
            affichage()
            fmt.Println(nomJoueur(), "vous avez gagné !")
            os.Exit(0) // on quitte la partie
        } else if partieNulle() { // Vérifier si match nul
            affichage()
            fmt.Println("Partie nulle !")
            os.Exit(0) // on quitte la partie
        }
        joueur1 = !joueur1 // on change de joueur
    }
}

/**
* description de la fonction : Permet d'afficher le damier en prenant en compte les cases remplies
*
* @return rien
*/
func affichage() {
    for i := 0; i < len(tableauMorpion); i++ {
        fmt.Print(" ", tableauMorpion[i], " ")
        if (i+1)%3 == 0 { // retour à la ligne après avoir affiché 3 éléments
            fmt.Println()
        }
    }
}

/**
* description de la fonction : Retourne le nom des joueurs
*
* @return string
*/
func nomJoueur() string {
    if joueur1 {
        return "Joueur 1 "
    } else {
        return "Joueur2 "
    }
}

/**
* description de la fonction : Permet de vérifier l'entrée utilisateur
*
* @return int : retourne l'entrée utilisateur
*/
func gestionEntreeUtilisateur() int {

    var (
        bonneEntree = false // variable qui permet de vérifier si l'utilisateur a rentré la valeur qu'on attend de lui
        numeroCase  = 0
        err         error
        scanner     = bufio.NewScanner(os.Stdin)
    )

    for bonneEntree == false {
        fmt.Print(nomJoueur(), "entrez un nombre compris entre 1 à ", tailleDamier, " : ")
        scanner.Scan()
        numeroCase, err = strconv.Atoi(scanner.Text())
        if err != nil { //vérifier si l'utilisateur a rentré un nombre
            fmt.Println("Entrez un nombre et non autre chose !")
        } else if numeroCase < 1 || numeroCase > tailleDamier {
            fmt.Println("Votre nombre doit être compris entre 0 à", tailleDamier, "!")
        } else if tableauMorpion[numeroCase-1] == symboleJoueur1 || tableauMorpion[numeroCase-1] == symboleJoueur2 { // vérifier si la case est libre
            fmt.Println("Cette case est déjà prise !")
        } else {
            bonneEntree = true
        }
    }
    return numeroCase - 1 // n'oubliez pas que la taille d'un tableau commence toujours par 0 ;)
}

/**
* description de la fonction : Permet de remplir la case choisie par le joueur
*
* @param numeroCase
* @return rien
*/
func remplirCase(numeroCase int) {
    if joueur1 {
        tableauMorpion[numeroCase] = symboleJoueur1
    } else {
        tableauMorpion[numeroCase] = symboleJoueur2
    }
}

/**
* description de la fonction : Permet de savoir si le joueur a gagné
*
* @return bool
*/
func gagner() bool {

    /*
        tableauxdeGain est un tableau à double dimensions où j'ai rajouté
        les différents cas d'utilisation où il est possible de gagner.
    */
    tableauxdeGain := [][tailleDamier]bool{
    {
        true, true, true,
        false, false, false,
        false, false, false},

    {
        false, false, true,
        false, false, true,
        false, false, true},
    {
        false, false, false,
        false, false, false,
        true, true, true},
    {
        true, false, false,
        true, false, false,
        true, false, false},
    {
        true, false, false,
        false, true, false,
        false, false, true},
    {
        false, false, true,
        false, true, false,
        true, false, false},
    {
        false, true, false,
        false, true, false,
        false, true, false}}

    // création d'un damier temporaire
    var tableauMorpionBool [tailleDamier]bool

    for index, valeur := range tableauMorpion {
        if joueur1 && valeur == symboleJoueur1 { // si c'est le tour du joueur 1 et que la case possède le bon symbole du joueur 1
            tableauMorpionBool[index] = true
        } else if !joueur1 && valeur == symboleJoueur2 {
            tableauMorpionBool[index] = true
        }
    }

    ressemblance := 0 // Nombre de true qui sont sur les mêmes cases dans le tableau tableauMorpionBool et dans le tableau tableauxdeGain
    for _, tableauGain := range tableauxdeGain {
        for i := 0; i < len(tableauMorpionBool); i++ {
            if tableauMorpionBool[i] == true && tableauMorpionBool[i] == tableauGain[i] { // si c'est à true dans le même index de tableauMorpionBool et tableauGain alors on incrémente la ressemblance
                ressemblance++
                if ressemblance == 3 { // si les cases du tableau tableauMorpionBool sont 3 fois les mêmes que sur l'un des tableaux tableauxdeGain alors ça veut dire que le joueur a gagné
                    return true
                }
            }
        }
        ressemblance = 0 // On remet le compteur à 0 pour vérifier un autre tableau du tableauxdeGain
    }
    return false
}

/*
* description de la fonction : Permet de vérifier si la partie est nulle
*
* @return bool
*/
func partieNulle() bool {
    occurence := 0

    for _, valeur := range tableauMorpion {
        if valeur == symboleJoueur1 || valeur == symboleJoueur2 {
            occurence++ // incrémenter de 1 si une case est remplite par un symbole
        }
    }

    /*
        si toutes les cases sont remplies de symboles
        et que le joueur n'a pas encore gagné alors la partie est nulle
    */
    return (occurence == len(tableauMorpion))
}
