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
		fmt.Print("_")
		for j := 0; j < len(sheet); j++ {
			switch {
			case sheet[j][i] == -2:
				fmt.Print("__")
			case sheet[j][i] == -1:
				fmt.Print("_0")
			case sheet[j][i] < 10:
				fmt.Print("_", sheet[j][i])
			default:
				fmt.Print(sheet[j][i])
			}
			if j < len(sheet) - 1 {
				fmt.Print("__")
			}
		}
		fmt.Print("\n")
	}
}

func flute_trainer() {
	fmt.Println("Flute Trainer")
}

func guitar_trainer() {
	sheet := make([][]int, 10)
	for i := range sheet {
		sheet[i] = make([]int, 6)
	}
	chords_set := []string{"y", "yes"}
	no_chords_set := []string{"n", "no"}
	var chords bool

	anchor_fret := 0

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
		for {
			for k := 0; k < 10; k++ {
				barre := rand_bool(1, 2)
				//third_finger := int(rand.Int31n(2))
				fourth_finger := int(rand.Int31n(2))// 1 if the chord'll use 4 fingers, otherwise 0
				chord_buffer := make([]int, 6)
				for i := 0; i < len(chord_buffer); i++{
					chord_buffer[i] = 99
				}
				//99 - untouched cell
				//-2 - skip
				//-1 - 0th fret 

				for i := 0; i < 3 + fourth_finger /*+ third_finger*/; i++ {
					fret_modifier := int(rand.Int31n(5))
					for {
						string_num := rand.Int31n(6)
						if chord_buffer[string_num] == 99 {			// checks if the string already
							chord_buffer[string_num] = fret_modifier	//'has a finger on it'
							break
						}
					}
				}

				if barre {
					for i := 0; i < slices.Min(chord_buffer); i++ {
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
				
				for i := 0; i < 6; i++ {
					switch {
					case chord_buffer[i] == -2:
						sheet[k][i] = -2
					case chord_buffer[i] == -1:
						sheet[k][i] = -1
					case chord_buffer[i] < 5:
						sheet[k][i] = anchor_fret + chord_buffer[i]
					default:
						log.Fatal()
					}
				}
				
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