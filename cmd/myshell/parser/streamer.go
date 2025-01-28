package parser

import (
	"fmt"
	"io"
	"strings"

	"github.com/codecrafters-io/shell-starter-go/cmd/myshell/autocomplete"
	"github.com/codecrafters-io/shell-starter-go/cmd/myshell/consts"
)

type Streamer struct {
	reader io.Reader
}

func NewStreamer(reader io.Reader) *Streamer {
	return &Streamer{reader: reader}
}

func (s *Streamer) GetNextCommand() (string, error) {
	buffer := make([]byte, 0, 100)
	byteReader := make([]byte, 1)
	isTab := false
	currSuggestions := ([]string)(nil)

	for {
		if _, err := s.reader.Read(byteReader); err != nil {
			return "", err
		}

		switch currByte := byteReader[0]; currByte {
		case 13: // Newline
			cmd := string(buffer)
			return cmd, nil
		case 0x7f: // Delete
			if len(buffer) == 0 {
				break
			}
			buffer = buffer[:len(buffer)-1]
			fmt.Printf("\b \b")
		case 0x03: // Ctrl + C
			return consts.EXIT, nil
		case 0x9:
			// Autocomplete logic
			currStr := string(buffer)
			closestEstimates, err := autocomplete.GetClosestCommands(currStr)
			if err != nil {
				fmt.Println("Error in command hit")
				return "", err
			}

			if isTab {
				buffer = []byte(autocomplete.GetCommonPrefix(currSuggestions))
				fmt.Printf("\r\n%s\r\n$ %s", strings.Join(currSuggestions, " "), string(buffer))
				isTab = false
				continue
			}

			switch len(closestEstimates) {
			case 0:
				fmt.Printf("\a")
				continue
			case 1:
				break
			default:
				isTab = true
				currSuggestions = closestEstimates
				fmt.Printf("\a")
				continue
			}

			closestEstimate := closestEstimates[0]
			remainingEst := closestEstimate[len(buffer):] + " "
			buffer = append(buffer, []byte(remainingEst)...)
			fmt.Printf("%s", remainingEst)
		default:
			buffer = append(buffer, currByte)
			fmt.Print(string(currByte))
		}
	}
}
