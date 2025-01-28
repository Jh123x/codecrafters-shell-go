package parser

import (
	"fmt"
	"io"

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
	for {
		if _, err := s.reader.Read(byteReader); err != nil {
			return "", err
		}

		switch currByte := byteReader[0]; currByte {
		case 13: // Newline
			fmt.Print("\r\n")
			cmd := string(buffer)
			return cmd, nil
		case 8: // Delete
			if len(buffer) == 0 {
				break
			}

			buffer = buffer[:len(buffer)-1]
		case 0x03: // Ctrl + C
			return consts.EXIT, nil
		case 0x9:
			// Autocomplete logic
			currStr := string(buffer)
			closestEstimate, err := autocomplete.GetClosestCommand(currStr)
			if err != nil {
				return "", err
			}

			if len(closestEstimate) < len(currStr) {
				break
			}

			remainingEst := closestEstimate[len(buffer):] + " "
			buffer = append(buffer, []byte(remainingEst)...)
			fmt.Printf("%s", remainingEst)
		default:
			buffer = append(buffer, currByte)
			fmt.Print(string(currByte))
		}
	}
}
