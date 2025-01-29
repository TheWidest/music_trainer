package main

import (
	"fmt"
	"slices"
	"bufio"
	"os"
	"log"
	"strings"
	"strconv"
	"math"
	"math/rand"
)
func printer(sheet [][]int) {
	fmt.Print("\n\n\n")
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
			fmt.Print("————")
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

func stringed_trainer(frets int, num_of_strings int, instrument string) {
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

	fmt.Println("How many attacks should there be in a set?")
	var number_of_attacks int
	for {
		str, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		str = strings.TrimSpace(str)

		if number_of_attacks_buf, err := strconv.Atoi(str); err == nil {
			number_of_attacks = number_of_attacks_buf
			break
		} else {
			fmt.Println("That's not a valid number, try again")
		}
	}
	sheet := make([][]int, number_of_attacks)
	for i := range sheet {
		sheet[i] = make([]int, num_of_strings)
	}

	fmt.Println("Press 'Enter' to get a new set")
	if chords {
		anchor_fret := int(rand.Int31n(int32(frets - 1)))
		for {
			for k := 0; k < number_of_attacks; k++ {
				chord_buffer := make([]int, num_of_strings)
				for i := 0; i < len(chord_buffer); i++{
					chord_buffer[i] = 99
				}
				var barre bool

				switch instrument { //generating initial notes (without barre or open strings)
				case "guitar":
					barre = rand_bool(1, 2)				
					fourth_finger := int(rand.Int31n(2))// 1 if the chord'll use extra finger,
														// otherwise 0

					for i := 0; i < 3 + fourth_finger; i++ { //'places fingers on strings'
						for {
							string_num := rand.Int31n(int32(num_of_strings))	//chooses the string for the finger
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
				case "bass":
					barre = rand_bool(1, 6)
					var third_finger int
					if rand_bool(1, 4) && !barre {
						third_finger = 1	// 1 if the chord'll use extra finger,
					} else {
						third_finger = 0	// otherwise 0
					}

					for i := 0; i < 2 + third_finger; i++ { //'places fingers on strings'
						for {
							string_num := rand.Int31n(int32(num_of_strings))	//chooses the string for the finger
							if chord_buffer[string_num] == 99 {	// checks if the string already 'has a finger on it'
								var fret_modifier int
								for {
									fret_modifier = int(rand.Int31n(4))
									if anchor_fret + fret_modifier <= frets { // checks if the fret is on the fretboard
										break
									}
								}
								chord_buffer[string_num] = anchor_fret + fret_modifier
								break
							}
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
				
				for i := 0; i < num_of_strings; i++ {
					if chord_buffer[i] == 99 {
						switch instrument {
						case "guitar":
							if rand_bool(1, 4) {
								chord_buffer[i] = 0
							} else {
								chord_buffer[i] = -1
							}
						case "bass":
							if rand_bool(1, 8) {
								chord_buffer[i] = 0
							} else {
								chord_buffer[i] = -1
							}
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
			fret := int(rand.Int31n(int32(frets + 1)))
			for k := 0; k < number_of_attacks; k++ {
				buff_attack := make([]int, num_of_strings)
				for i := 0; i < num_of_strings; i++ {
					buff_attack[i] = -1
				}
				fret = int(math.Abs(float64((fret + int(rand.NormFloat64()*9)) % frets)))
				buff_attack[int(rand.Int31n(int32(num_of_strings)))] = fret
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
	inst_options := []string{"f", "g", "b"} // f - flute, g - guitar, b - bass
	var str string

	fmt.Println("Choose the instrument")
	fmt.Println("f - flute, g - guitar, b - bass")

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
		stringed_trainer(22, 6, "guitar")
	case "b":
		stringed_trainer(20, 4, "bass")
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