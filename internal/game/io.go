package game

import (
	"fmt"
	"strconv"
)

func (g Game) getInput() (cmd uint8) {
	for fmt.Print(PROMPT); g.scanner.Scan(); fmt.Print(PROMPT) {
		line := g.scanner.Text()
		if cmd = parseCmd(line); cmd != ERROR_CODE {
			return
		}
		fmt.Println("I am sorry explorer, I did not understand thy command.")
	}
	cmd = ERROR_CODE
	return
}

func parseUint8(s string) (uint8, bool) {
  ui64, err := strconv.ParseUint(s, 10, 8)
  if err != nil {
    return ERROR_CODE, false
  }
  return uint8(ui64), true
}
