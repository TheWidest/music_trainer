package main

import (
	"fmt"
	"slices"
	//"strconv"
	"bufio"
	"os"
	"log"
	"strings"
	"math/rand"
)

const frets = 22

func printer(sheet [][]int) {
	fmt.Print("\n")
	for i := 0; i < len(sheet[0]); i++ {
		fmt.Print("—")
		for j := 0; j < len(sheet); j++ {
			switch {
			case sheet[j][i] == -2:
				fmt.Print("———")
			case sheet[j][i] == -1:
				fmt.Print("— 0")
			case sheet[j][i] < 10:
				fmt.Print("——", sheet[j][i])
			default:
				fmt.Print(" ", sheet[j][i])
			}
			if j < len(sheet) - 1 {		w
				fmt.Print("————")
			}
		}
		fmt.Print("\n")
	}
}

func flute_trainer() {
	fmt.Println("Flute Trainer")
}

func guitar_trainer() {
	number_of_attacks := 10
	sheet := make([][]int, number_of_attacks)
	for i := range sheet {
		sheet[i] = make([]int, 6)
	}
	chords_set := []string{"y", "yes"}
	no_chords_set := []string{"n", "no"}
	var chords bool

	fmt.Println("Do you want to use chords?")
	for {
		str, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		str = strings.TrimSpace(str)
		str = strings.ToLower(str)
		if slices.Contains(chords_set, str) {
			chords = true
			break
		} else if slices.Contains(no_chords_set, str) {
			chords = false
			break
		} else {
			fmt.Println("That's not a valid answer. Try 'y', 'yes', 'n' or 'no'")
		}
	}

	fmt.Println("Press 'Enter' to get a new set")
	if chords {
		anchor_fret := int(rand.Int31n(21))
		for {
			for k := 0; k < number_of_attacks; k++ {
				barre := rand_bool(1, 2)
				third_finger := int(rand.Int31n(2))	// 1 if the chord'll uses extra finger,
				fourth_finger := int(rand.Int31n(2))// otherwise 0
				chord_buffer := make([]int, 6)
				for i := 0; i < len(chord_buffer); i++{
					chord_buffer[i] = 99
				}
				//99 - untouched cell
				//-2 - skip
				//-1 - 0th fret 

				for i := 0; i < 3 + third_finger + fourth_finger; i++ { //'places fingers on strings'
					for {
						string_num := rand.Int31n(6)	//chooses the string for the finger
						if chord_buffer[string_num] == 99 {	// checks if the string already 'has a finger on it'
							var fret_modifier int
							for {
								fret_modifier = int(rand.Int31n(5))
								if anchor_fret + fret_modifier <= frets { // checks if the fret is on the fretboard
									break
								}
							}
							chord_buffer[string_num] = anchor_fret + fret_modifier
							break
						}
					}
				}

				if barre {
					for i := 0; i < slices.Index(chord_buffer, slices.Min(chord_buffer)); i++ {
						if chord_buffer[i] == 99 {
							chord_buffer[i] = slices.Min(chord_buffer)
						}
					}
				}
				for i := 0; i < 6; i++ {
					if chord_buffer[i] == 99 {
						if rand_bool(1, 4) {
							chord_buffer[i] = -1
						} else {
							chord_buffer[i] = -2
						}
					}
				}
				fmt.Println("4_________")
				sheet[k] = chord_buffer
				// for i := 0; i < 6; i++ {
				// 	switch {
				// 	case chord_buffer[i] == -2:
				// 		sheet[k][i] = -2
				// 	case chord_buffer[i] == -1:
				// 		sheet[k][i] = -1
				// 	case chord_buffer[i] < 5:
				// 		sheet[k][i] = anchor_fret + chord_buffer[i]
				// 	default:
				// 		log.Fatal()
				// 	}
				// }
			}
			printer(sheet)
			_, err := bufio.NewReader(os.Stdin).ReadString('\n')
			if err != nil {
				log.Fatal(err)
			}
		}
	} else {
		fmt.Println("single note")
	}

}

func initiator(inst_letter *string) {
	inst_options := []string{"f", "g"} // f - flute, g - guitar

	fmt.Println("Choose the instrument")
	fmt.Println("f - flute, g - guitar")
	flag := false
	for flag == false {
		str, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		str = strings.TrimSpace(str)
		str = strings.ToLower(str)
		if slices.Contains(inst_options, str) {
			*inst_letter = str
			flag = true
		} else {
			fmt.Println("that's not a valid option")
		}
	}
	switch *inst_letter {
	case "f":
		flute_trainer()
	case "g":
		guitar_trainer()
	}
}

func main() {
	var letter string

	initiator(&letter)

	//fmt.Println(letter)
}

func rand_bool(numerator int, denominator int) bool {
	if int(rand.Int31n(int32(denominator))) < numerator {
		return true
	} else {
		return false
	}
}