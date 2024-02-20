package game

import "bufio"

type Room struct {
  room []uint8
  description string
}

type Monster struct {
  name string
  ferocity uint8
}

type Player struct {
  playerName string
  monsterKilled uint8
  food uint16
  strength int16
  wealth uint16
  sword bool
  amulet bool
  axe bool
  suit bool
  light bool
}

type Game struct {
  rooms []Room
  nrRooms uint8
  currentRoom uint8
  player Player
  tally int
  ferocityFactor uint8
  scanner *bufio.Scanner
  prompt string
}
