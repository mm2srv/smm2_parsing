package smm2_parsing

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"os"
	"strings"
)

type ReplayEntryType uint8

const (
	Time ReplayEntryType = iota
	Joysticks
	Inputs
)

type ReplayInputType uint8

const (
	X ReplayInputType = iota
	Y
	A
	B
	R
	L
	ZR
	ZL
	Up
	Down
	Left
	Right
	Plus
	Minus
	JoyUp
	JoyDown
	JoyLeft
	JoyRight
)

type Entry struct {
	entry_type ReplayEntryType
	// Time entry
	frames uint8
	// Joysticks entry
	joystick_x int16
	joystick_y int16
	// Inputs entry
	inputs []ReplayInputType
}

type Replay struct {
	entries     []Entry
	totalFrames int
}

func (s *Replay) InputToName(input ReplayInputType) string {
	return []string{
		"KEY_X",
		"KEY_Y",
		"KEY_A",
		"KEY_B",
		"KEY_R",
		"KEY_L",
		"KEY_ZR",
		"KEY_ZL",
		"KEY_DUP",
		"KEY_DDOWN",
		"KEY_DLEFT",
		"KEY_DRIGHT",
		"KEY_PLUS",
		"KEY_MINUS",
		"KEY_JUP",
		"KEY_JDOWN",
		"KEY_JLEFT",
		"KEY_JRIGHT",
	}[input]
}

func (s *Replay) HandleInput(reader *bytes.Reader) bool {
	input := make([]byte, 0x4)
	binary.Read(reader, binary.LittleEndian, input)

	var entry Entry
	entry.entry_type = Inputs

	if (input[3] & 0b00001000) == 0b00001000 {
		entry.inputs = append(entry.inputs, X)
	}

	if (input[3] & 0b00010000) == 0b00010000 {
		entry.inputs = append(entry.inputs, Y)
	}

	if (input[3] & 0b00000001) == 0b00000001 {
		entry.inputs = append(entry.inputs, A)
	}

	if (input[3] & 0b00000010) == 0b00000010 {
		entry.inputs = append(entry.inputs, B)
	}

	if (input[2] & 0b01000000) == 0b01000000 {
		entry.inputs = append(entry.inputs, R)
	}

	if (input[2] & 0b00100000) == 0b00100000 {
		entry.inputs = append(entry.inputs, L)
	}

	if (input[3] & 0b00100000) == 0b00100000 {
		entry.inputs = append(entry.inputs, ZR)
	}

	if (input[3] & 0b00000100) == 0b00000100 {
		entry.inputs = append(entry.inputs, ZL)
	}

	if (input[1] & 0b00000001) == 0b00000001 {
		entry.inputs = append(entry.inputs, Up)
	}

	if (input[1] & 0b00000010) == 0b00000010 {
		entry.inputs = append(entry.inputs, Down)
	}

	if (input[1] & 0b00000100) == 0b00000100 {
		entry.inputs = append(entry.inputs, Left)
	}

	if (input[1] & 0b00001000) == 0b00001000 {
		entry.inputs = append(entry.inputs, Right)
	}

	if (input[2] & 0b00001100) == 0b00001100 {
		// Multiple bits
		entry.inputs = append(entry.inputs, Plus)
	}

	if (input[2] & 0b00010010) == 0b00010010 {
		// Multiple bits
		entry.inputs = append(entry.inputs, Minus)
	}

	// The "joystick" inputs are not always consistent with the joysticks
	if (input[1] & 0b00010000) == 0b00010000 {
		entry.inputs = append(entry.inputs, JoyUp)
	}

	if (input[1] & 0b00100000) == 0b00100000 {
		entry.inputs = append(entry.inputs, JoyDown)
	}

	if (input[1] & 0b01000000) == 0b01000000 {
		entry.inputs = append(entry.inputs, JoyLeft)
	}

	if (input[1] & 0b10000000) == 0b10000000 {
		entry.inputs = append(entry.inputs, JoyRight)
	}

	/*
		var inputString string
		if len(inputDebug) == 0 {
			inputString = "NONE"
		} else {
			inputString = strings.Join(inputDebug, ", ")
		}
		fmt.Printf("%v %s\n", input, inputString)
	*/

	s.entries = append(s.entries, entry)

	// Return if joystick data is included
	joystickIncluded := (input[1] & 0b10000000) == 0b10000000
	return joystickIncluded
}

func (s *Replay) HandleJoysticks(reader *bytes.Reader) {
	joysticks := make([]int16, 0x2)
	binary.Read(reader, binary.BigEndian, joysticks)

	var entry Entry
	entry.entry_type = Joysticks
	entry.joystick_x = joysticks[0]
	entry.joystick_y = joysticks[1]

	s.entries = append(s.entries, entry)

	// Ranges from -2^14 (-16384) to 2^14 (16384) in both X and Y
	//fmt.Printf("Joysticks %d %d\n", joysticks[0], joysticks[1])
}

func (s *Replay) HandleTimePassed(units int) {
	//units -= 1

	if units > 0 {
		var entry Entry
		entry.entry_type = Time
		entry.frames = uint8(units)
		s.entries = append(s.entries, entry)

		// Not yet in sync, TODO
		fmt.Printf("Frames passed %d\n", units)
		s.totalFrames += units
	}
}

func (s *Replay) HandleEndCap(reader *bytes.Reader) {
	// Bytes generally with nonzero first bytes and nonzero last bytes, 0 everywhere else
	endCap := make([]byte, 0x7)
	binary.Read(reader, binary.LittleEndian, endCap)

	// TODO right number can sometimes be 2 bytes
	//fmt.Printf("EndCap %d %d\n", endCap[0], endCap[6])
}

func (s *Replay) Load(buf []byte) error {
	reader := bytes.NewReader(buf)

	// Always the same
	magic := make([]byte, 4)
	binary.Read(reader, binary.LittleEndian, magic)

	// Always the same
	unk1 := make([]byte, 8)
	binary.Read(reader, binary.LittleEndian, unk1)

	// Changes
	unk2 := make([]byte, 0x10)
	binary.Read(reader, binary.LittleEndian, unk2)

	// Changes
	unk3 := make([]byte, 0x3C)
	binary.Read(reader, binary.LittleEndian, unk3)

	// The 0x42 checks
	unk4 := make([]byte, 0x9)
	binary.Read(reader, binary.LittleEndian, unk4)

	// The 0x40 or 0x41
	firstKey, _ := reader.ReadByte()
	if firstKey == 0x40 {
		s.HandleEndCap(reader)
	} else if firstKey == 0x41 {
		// Handle one input
		if s.HandleInput(reader) {
			// Handle joysticks
			s.HandleJoysticks(reader)
		}

		s.HandleEndCap(reader)
	}

	ended := false
	for !ended {
		_, _ = reader.ReadByte()

		includeEnding := true

		for {
			//pos, _ := reader.Seek(0, 1)
			//fmt.Printf("Pos111 %x\n", pos)

			key, _ := reader.ReadByte()
			time, _ := reader.ReadByte()

			// For some cursed reason it's possible for joysticks to immediately continue after s byte
			// TODO ensure s check is correct for all replays (have encountered 0x00 and 0x01 for preContinueKey[1])
			if (key & 0b00000100) == 0b00000100 {
				// If s happens, seek back 2 bytes
				key = 0x80
				time = 0x00
				reader.Seek(-2, 1)
			}

			s.HandleTimePassed(int(time))

			if key == 0x00 && time == 0x40 {
				// Fully break
				break
			} else if key == 0x00 && time == 0x10 {
				// Only seen in one joystick run, end 1 byte early
				includeEnding = false
				ended = true
				break
			} else if key == 0x00 && time == 0x01 {
				// Read an input and continue
				s.HandleInput(reader)
				continue
			} else {
				endFrameCheck := false
				for {
					//pos, _ := reader.Seek(0, 1)
					//fmt.Printf("Pos %x\n", pos)

					var joysticks byte = 0

					joysticks, _ = reader.ReadByte()
					continueKey, _ := reader.ReadByte()

					if continueKey == 0x01 {
						// Read in input
						// (When joysticks, not sure about the input)
						//s.handleTimePassed(0x1)
						s.HandleInput(reader)

						if (joysticks & 0b00000100) == 0b00000100 {
							// Joysticks
							s.HandleJoysticks(reader)

							checkReadAgain, _ := reader.ReadByte()
							// Go backwards one byte
							reader.Seek(-1, 1)

							if checkReadAgain == 0x80 {
								break
							}
						} else {
							break
						}
					} else if continueKey == 0x41 {
						// Read in input and break
						s.HandleInput(reader)
						endFrameCheck = true
						break
					} else if continueKey == 0x10 {
						// End of the file
						//handleTimePassed(0x10)
						endFrameCheck = true
						includeEnding = false
						ended = true
						break
					} else {
						//s.handleTimePassed(int(continueKey))
						// Joysticks (usually 0x00)
						if (joysticks & 0b00000100) == 0b00000100 {
							//s.handleTimePassed(1)
							s.HandleJoysticks(reader)
						}

						if continueKey == 0x40 {
							// Stop repeating input readings
							endFrameCheck = true
							break
						} else {
							// Try to continue joysticks
							checkReadAgain, _ := reader.ReadByte()
							// Go backwards one byte
							reader.Seek(-1, 1)

							if checkReadAgain == 0x80 {
								break
							}
						}
					}
				}

				if endFrameCheck {
					break
				}
			}
		}

		if includeEnding {
			s.HandleEndCap(reader)
		}
	}

	// End of file, always different
	unk9 := make([]byte, 0x4)
	binary.Read(reader, binary.LittleEndian, unk9)

	currentPosition, _ := reader.Seek(0, 1)
	fmt.Printf("End %d %d\n", currentPosition, reader.Size())
	fmt.Printf("All frames %d\n", s.totalFrames)

	if currentPosition == reader.Size() {
		return nil
	} else {
		return fmt.Errorf("Replay did not end properly")
	}
}

func (s *Replay) GetTASText() string {
	output := ""
	currentTime := 1
	lastTime := 1
	currentState := Entry{}
	for _, entry := range s.entries {
		if entry.entry_type == Time {
			currentTime += int(entry.frames)

			for i := lastTime; i <= currentTime; i++ {
				// Write out a frame
				inputString := "NONE"
				if len(currentState.inputs) != 0 {
					var inputStrings []string
					for _, input := range currentState.inputs {
						inputStrings = append(inputStrings, s.InputToName(input))
					}
					inputString = strings.Join(inputStrings, ";")
				}

				output += fmt.Sprintf("%d %s %d;%d 0;0\n", i, inputString, currentState.joystick_x*2, currentState.joystick_y*2)
			}

			lastTime = currentTime
		} else if entry.entry_type == Joysticks {
			currentState.joystick_x = entry.joystick_x
			currentState.joystick_y = entry.joystick_y
		} else if entry.entry_type == Inputs {
			currentState.inputs = entry.inputs
		}
	}
	return output
}

func main() {
	//replays := []string{"auto_long_9.699.bin", "auto_long_holdrun_9.699.bin", "auto_long_holdrun2_9.699.bin", "auto_long_holdthenrelease_9.699.bin", "auto_long_holdthenrelease16_9.699.bin", "auto_right.bin", "joysticks_complicated.bin", "kaizo_10.761.bin", "long_27.051.bin", "long_level.bin", "long_level_dpad.bin", "right.bin", "right_backandforth_11.000.bin", "right_dpad.bin", "right_dpad_complicated.bin", "right_dpad_complicated2.bin", "right_dpad_pause.bin", "right_dpad_run.bin", "right_joystick_run.bin", "right_joystick_run2.bin", "right_long.bin", "right_long_5.466.bin", "right_long_dpad.bin", "right_long_dpad_pause.bin", "right_long_pause.bin", "right_pause.bin", "right_spinjump_2.866.bin", "up_down_left_right_joysticks.bin", "up_down_left_right_joysticks2.bin"}
	replays := []string{"testreplay.bin"}

	for _, replayName := range replays {
		replay, err := os.ReadFile(replayName)
		if err != nil {
			panic(err)
		}

		reader := &Replay{}
		err = reader.Load(replay)
		if err != nil {
			panic(err)
		}

		os.WriteFile("script0-1.txt", []byte(reader.GetTASText()), 0644)
	}

}
