package main

import (
	"fmt"
	"slices"
	"bufio"
	"os"
	"log"
	"strings"
	"math"
	"math/rand"
)

const frets = 22

func printer(sheet [][]int) {
	fmt.Print("\n")
	for i := 0; i < len(sheet[0]); i++ {
		fmt.Print("—")
		for j := 0; j < len(sheet); j++ {
			switch {
			case sheet[j][i] == -1:
				fmt.Print("———")
			case sheet[j][i] < 10:
				fmt.Print("——", sheet[j][i])
			default:
				fmt.Print("—", sheet[j][i])
			}
			if j < len(sheet) - 1 {
				fmt.Print("————")
			}
		}
		fmt.Print("\n")
	}
}

func flute_trainer() {
	fmt.Println("flute dud")
	// number_of_attacks := 10
	// sheet := make([][]int, number_of_attacks)
	// for i := range sheet {
	// 	sheet[i] = make([]int, 6)
	// }
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
				//-1 - skip 

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
							chord_buffer[i] = 0
						} else {
							chord_buffer[i] = -1
						}
					}
				}
				sheet[k] = chord_buffer
				anchor_fret = int(math.Abs(float64((anchor_fret + int(rand.NormFloat64()*5)) % frets)))
			}
			printer(sheet)
			_, err := bufio.NewReader(os.Stdin).ReadString('\n')
			if err != nil {
				log.Fatal(err)
			}
		}
	} else {
		for {
			fret := int(rand.Int31n(23))
			for k := 0; k < number_of_attacks; k++ {
				buff_attack := make([]int, 6)
				for i := 0; i < 6; i++ {
					buff_attack[i] = -1
				}
				fret = int(math.Abs(float64((fret + int(rand.NormFloat64()*9)) % frets)))
				buff_attack[int(rand.Int31n(6))] = fret
				sheet[k] = buff_attack
			}

			printer(sheet)
			_, err := bufio.NewReader(os.Stdin).ReadString('\n')
			if err != nil {
				log.Fatal(err)
			}
		}
	}
}

func initiator() {
	inst_options := []string{"f", "g"} // f - flute, g - guitar
	var str string

	fmt.Println("Choose the instrument")
	fmt.Println("f - flute, g - guitar")

	for {
		str_buf, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		str_buf = strings.TrimSpace(str_buf)
		str_buf = strings.ToLower(str_buf)
		if slices.Contains(inst_options, str_buf) {
			str = str_buf
			break
		} else {
			fmt.Println("that's not a valid option")
		}
	}
	switch str {
	case "f":
		flute_trainer()
	case "g":
		guitar_trainer()
	}
}

func main() {
	initiator()
}

func rand_bool(numerator int, denominator int) bool {
	if int(rand.Int31n(int32(denominator))) < numerator {
		return true
	} else {
		return false
	}
}