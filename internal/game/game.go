package game

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func (g Game) Start() {
	for {
		if g.currentRoom == EXIT {
			g.winTheGame()
		}
		g.describeRoom()
		if g.currentRoom == LIFT {
			time.Sleep(TIMEOUT * 3)
			g.currentRoom = REAR_VESTIBULE
		  fmt.Println("\n" + strings.Repeat("-", 50))
			continue
		}
		fmt.Println("\n" + strings.Repeat("-", 50))
		g.player.DisplayStatus()
		fmt.Println("\n" + strings.Repeat("-", 50))
		cmd := g.getInput()
		g.execCmd(cmd)
	}
}

func (g Game) displayScoreAndExit(msg string) {
	time.Sleep(TIMEOUT)
	fmt.Println(msg)
	g.displayScore()
	os.Exit(0)
}

func (g *Game) increaseTally() {
	g.tally++
}

func (g Game) winTheGame() {
	fmt.Println("You've done it!!")
	time.Sleep(TIMEOUT)
	fmt.Println("You have succeeded,", g.player.playerName)
	time.Sleep(TIMEOUT)
	fmt.Println("You managed to get out of the castle!")
	g.displayScoreAndExit("Well done!!!")
}

func (g Game) checkValidDirection(cmd uint8) bool {
	// the command is not related to movement
	if cmd > 5 {
		return true
	}
	if g.rooms[g.currentRoom].room[cmd] == 0 {
		var msg string
		switch cmd {
		case NORTH:
			msg = "No exit that way\n"
		case SOUTH:
			msg = "There is no exit south\n"
		case EAST:
			msg = "You cannot go in that direction\n"
		case WEST:
			msg = "You cannot move through solid stone\n"
		case UP:
			msg = "There is no way up from here\n"
		case DOWN:
			msg = "Cannot descend from here\n"
		}
		fmt.Println(msg)
		return false
	}
	return true
}

func (g *Game) execCmd(cmd uint8) {
	if !g.checkValidDirection(cmd) {
		return
	}

	switch cmd {
	case NORTH:
		fallthrough
	case SOUTH:
		fallthrough
	case EAST:
		fallthrough
	case WEST:
		fallthrough
	case UP:
		fallthrough
	case DOWN:
		g.doMove(cmd)
	case FIGHT:
		g.fight()
	case RUN:
		g.tryToRunAway()
	case MAGIC_AMULET:
		g.teleportWithAmulet()
	case PICK_UP:
		g.pickUpTreasure()
	case INVENTORY:
		g.inventory()
	case CONSUME:
		g.consumeFood()
	case QUIT:
		g.quit()
	default:
		// if you press CTRL-D after a space, this will prevent the crash
		if cmd != ERROR_CODE {
			panic(fmt.Sprintf("INVALID COMMAND: %v", cmd))
		} else {
			// if you have pressed CTRL-D you exit the game
			fmt.Printf("\nSo long, and thanks for all the fish my master!\n\n")
			os.Exit(0)
		}
	}

	g.increaseTally()
}

func parseCmd(cmd string) uint8 {
	cmd = strings.ToLower(cmd)
	if cmd == "n" || cmd == "north" {
		return NORTH
	}
	if cmd == "s" || cmd == "south" {
		return SOUTH
	}
	if cmd == "e" || cmd == "east" {
		return EAST
	}
	if cmd == "w" || cmd == "west" {
		return WEST
	}
	if cmd == "u" || cmd == "up" {
		return UP
	}
	if cmd == "d" || cmd == "down" {
		return DOWN
	}
	if cmd == "f" || cmd == "fight" {
		return FIGHT
	}
	if cmd == "r" || cmd == "run" {
		return RUN
	}
	if cmd == "m" || cmd == "ma" || cmd == "amulet" || cmd == "magic" {
		return MAGIC_AMULET
	}
	if cmd == "i" || cmd == "inv" || cmd == "inventory" {
		return INVENTORY
	}
	if cmd == "q" || cmd == "quit" || cmd == "exit" {
		return QUIT
	}
	if cmd == "p" || cmd == "pick" {
		return PICK_UP
	}
	if cmd == "c" || cmd == "consume" || cmd == "eat" {
		return CONSUME
	}
	return ERROR_CODE
}

func (g Game) describeRoom() {
	if !g.player.light {
		fmt.Println("It is too dark to see anything.")
	}

	fmt.Printf("\n%s\n", g.rooms[g.currentRoom].description)

	roomContent := g.rooms[g.currentRoom].room[ROOM_CONTENT]
	hasTreasure := roomContent > 0 && roomContent < 251 
	hasMonster := roomContent > 251
	if hasTreasure {
		fmt.Printf("There is treasure here worth $%v.\n", roomContent)
	} else if hasMonster {
		fmt.Println("DANGER...THERE IS A MONSTER HERE....")
		time.Sleep(TIMEOUT)
		fmt.Println("It is a", MONSTERS[roomContent].name)
		fmt.Println("The danger level is:", MONSTERS[roomContent].ferocity)
	}
	time.Sleep(TIMEOUT)
}

func (g Game) displayScore() {
	finalScore := 3*g.tally +
		5*int(g.player.strength) +
		2*int(g.player.wealth) +
		int(g.player.food) +
		30*int(g.player.monsterKilled)
	fmt.Println("\n\nYour final score is:", finalScore)
}
