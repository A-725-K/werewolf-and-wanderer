package game

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func readMazeFromFile(inputPath string) (rooms []Room, nrRooms uint8) {
  f, err := os.Open(inputPath)
  if err != nil {
    panic(fmt.Sprintf("Cannot open map input file: %v", err))
  }
  scanner := bufio.NewScanner(f)
  i := 0
  for scanner.Scan() {
    i++
    line := scanner.Text()
    fields := strings.Split(line, ":")
    if len(fields) != 2 {
      panic("Malformed input file 1:" + strconv.Itoa(i))
    }
    description := fields[1]
    if description == "" {
      panic("Malformed input file 2:" + strconv.Itoa(i))
    }
    description = strings.ReplaceAll(description, "@", "\n")
    fields = strings.Split(fields[0], ",")
    if len(fields) != 6 {
      panic("Malformed input file 3:" + strconv.Itoa(i))
    }
    var room Room
    for _, n := range fields {
      ui8, err := strconv.ParseUint(n, 10, 8)
      if err != nil {
        panic("Malformed input file 4:" + strconv.Itoa(i))
      }
      room.room = append(room.room, uint8(ui8))
    }
    room.room = append(room.room, 0)
    room.description = description
    rooms = append(rooms, room)
  }
  nrRooms = uint8(len(rooms))
  return
}

func initTreasure(rooms *[]Room, nrRooms uint8, fixed bool) {
  genRandom := func (n int) uint8 {
    return uint8(rand.Intn(100)+n)
  }
	if fixed {
		(*rooms)[PRIVATE_MEETING_ROOM].room[ROOM_CONTENT] = genRandom(100)
		(*rooms)[TREASURY].room[ROOM_CONTENT] = genRandom(100)
		return
	}

	for i := 0; i < NR_TREASURES; i++ {
		roomIdx := START
		for roomIdx == START || roomIdx == EXIT || (*rooms)[roomIdx].room[ROOM_CONTENT] > 0 {
			roomIdx = rand.Intn(int(nrRooms))
		}

		treasure := genRandom(10)
		(*rooms)[roomIdx].room[ROOM_CONTENT] = treasure
	}
}

func initMonsters(rooms *[]Room, nrRooms uint8) {
	monsterId := uint8(255)
	for i := 0; i < NR_MONSTERS; i++ {
		roomIdx := START
		for roomIdx == START || roomIdx == EXIT || (*rooms)[roomIdx].room[ROOM_CONTENT] > 0 {
			roomIdx = rand.Intn(int(nrRooms))
		}

		(*rooms)[roomIdx].room[ROOM_CONTENT] = monsterId
		monsterId--
	}
}

func getPlayerName(scanner *bufio.Scanner) (name string) {
	prompt := "What is thy name, explorer?  "
	for fmt.Print(prompt); scanner.Scan(); fmt.Print(prompt) {
    name = strings.TrimSpace(scanner.Text())
    lenOk := len(name) >= 3 && len(name) <= 15
    charsValid, _ := regexp.MatchString(PLAYER_NAME_REGEX, name)
    if lenOk && charsValid {
      break
    }
    fmt.Println("I cannot accept it, explorer! Call it again!")
  }
  if err := scanner.Err(); err != nil {
    fmt.Fprintln(os.Stderr, err)
  } else if len(name) == 0 {
  	fmt.Println("\nThe game, my master, is over even before it could start!")
  	os.Exit(0)
  }
  fmt.Println()
	return
}

func initPlayer(scanner *bufio.Scanner) Player {
	return Player{
  	playerName: getPlayerName(scanner),
  	monsterKilled: 0,
  	food: 0,
  	strength: INITIAL_STRENGTH,
  	wealth: INITIAL_WEALTH,
  	sword: false,
  	amulet: false,
  	axe: false,
  	suit: false,
  	light: false,
	}
}

func InitGame() Game {
	rand.Seed(time.Now().UnixNano())
	scanner := bufio.NewScanner(os.Stdin)

  rooms, nrRooms := readMazeFromFile(INPUT_MAP_FILE)
  initTreasure(&rooms, nrRooms, false)
  initMonsters(&rooms, nrRooms)
  initTreasure(&rooms, nrRooms, true)

  player := initPlayer(scanner)

  return Game{
    rooms: rooms,
    nrRooms: nrRooms,
    currentRoom: START,
    player: player,
  	tally: 0,
  	ferocityFactor: 0,
  	scanner: scanner,
  	prompt: "What do you want to do?  ",
  }
}
