package game

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

func (g Game) checkDeath() {
	if g.player.strength <= 0 {
		g.displayScoreAndExit("\nYou have died.........")
	}
}

func (g *Game) doMove(direction uint8) {
	g.currentRoom = g.rooms[g.currentRoom].room[direction]
}

func (g *Game) inventory() {
  if g.player.wealth <= 0 {
    fmt.Println("You have no money!")
    time.Sleep(TIMEOUT)
    return
  }
  fmt.Printf("\nYou have %d$. You can buy:\n\n", g.player.wealth)
  fmt.Printf("  1 - FLAMING TORCH (%d$)\n", TORCH_PRICE)
  fmt.Printf("  2 - AXE (%d$)\n", AXE_PRICE)
  fmt.Printf("  3 - SWORD (%d$)\n", SWORD_PRICE)
  fmt.Printf("  4 - FOOD (%d$ per unit)\n", FOOD_PRICE)
  fmt.Printf("  5 - MAGIC AMULET (%d$)\n", AMULET_PRICE)
  fmt.Printf("  6 - SUIT ARMOR (%d$)\n\n", ARMOR_PRICE)
  fmt.Printf("  0 - RETURN TO ADVENTURE\n\n")

  prompt := "Insert the number of the item you need:  "
  var itemN uint8 = ERROR_CODE
  var qty uint8 = ERROR_CODE
  for fmt.Print(prompt); g.scanner.Scan(); fmt.Print(prompt) {
    choice := strings.TrimSpace(g.scanner.Text())
    itemUi8, ok := parseUint8(choice)
    if !ok || itemUi8 > uint8(len(ITEM_PRICES)) {
    	fmt.Println("I do not recognize this item, please master, say it again.", )
      continue
    }
    itemN = itemUi8
    break
  }
	if itemN == ERROR_CODE {
		fmt.Println("Nothing has been done! 1")
		return
	}
  if itemN == 0 {
    time.Sleep(TIMEOUT)
    return
  }
  if itemN == FOOD_ITEM {
  	prompt = "How many units of food?  "
  	for fmt.Print(prompt); g.scanner.Scan(); fmt.Print(prompt) {
  		qtyStr := strings.TrimSpace(g.scanner.Text())
			qtyUi8, ok := parseUint8(qtyStr)
			if !ok {
    		fmt.Println("This unit has not been invented yet! Try again.")
      	continue
			}
    	qty = qtyUi8
			break
    }
    if qty == ERROR_CODE {
			fmt.Println("Nothing has been done!")
			return
    }
  } else {
    qty = 1
  }
  g.player.BuyItem(itemN, qty)
	time.Sleep(TIMEOUT)
}

func (g *Game) tryToRunAway() {
	thereIsNoMonster := g.rooms[g.currentRoom].room[ROOM_CONTENT] < 252 
	if thereIsNoMonster {
		fmt.Println("There is nothing to fight here...")
		return
	}
	var cmd uint8 = ERROR_CODE
	if prob := rand.Intn(100); prob > RUN_SUCCESS_RATE {
		prompt := "Which way do you want to flee?  "
		for fmt.Print(prompt); g.scanner.Scan(); fmt.Print(prompt) {
			direction := strings.ToLower(g.scanner.Text())
			cmd = parseCmd(direction)
			if g.checkValidDirection(cmd) {
				break
			}
			fmt.Println("I cannot let you move in that direction, wanderer!")
		}
		g.doMove(cmd)
	} else {
		fmt.Println("NO! YOU MUST STAND AND FIGHT!!!")
		g.fight()
	}
}

func (g *Game) pickUpTreasure() {
	roomContent := g.rooms[g.currentRoom].room[ROOM_CONTENT]
	if roomContent == 0 || roomContent > 251 {
		fmt.Println("There is no treasure to pick up here!")
	} else if !g.player.light {
		fmt.Println("It's too dark to see anything in the room!")
	} else {
		g.player.PickUpTreasure(uint16(roomContent))
		g.rooms[g.currentRoom].room[ROOM_CONTENT] = 0
	}
}

func (g *Game) consumeFood() {
	if g.player.food < 1 {
		fmt.Println("You have no food!")
		return
	}
	time.Sleep(TIMEOUT)
	fmt.Println("You have", g.player.food, "units of food")
	prompt := "How many do you want to eat?  "
	for fmt.Print(prompt); g.scanner.Scan(); fmt.Print(prompt) {
		qtyU64, err := strconv.ParseUint(g.scanner.Text(), 10, 16)
		if err != nil {
			fmt.Println("You cannot eat", qtyU64, "units of food!")
			continue
		}
		qty := uint16(qtyU64)
		if qty == 0 || qty > g.player.food {
			fmt.Println("You don't have enough food!")
			continue
		}
		g.player.ConsumeFood(qty)
		time.Sleep(TIMEOUT)
		fmt.Println("You have recovered", qty*FOOD_STRENGTH_VALUE, "health points")
		break
	}
}
 
func (g *Game) teleportWithAmulet() {
	time.Sleep(TIMEOUT)
	n := len(g.rooms)
	newRoom := rand.Intn(n)
	for newRoom == START || newRoom == EXIT {
		newRoom = rand.Intn(n)
	}
	g.currentRoom = uint8(newRoom)
}

func (g Game) quit() {
	prompt := "Are you sure you want to abandon the game? (Y|n)  "
	for fmt.Print(prompt); g.scanner.Scan(); fmt.Print(prompt) {
		answer := strings.ToLower(g.scanner.Text())
		if answer == "" || answer == "y" || answer == "yes" {
			msg := fmt.Sprintf(
				"\nBye bye, %s! It has been a pleasure!",
				g.player.playerName,
			)
			g.displayScoreAndExit(msg)
		} else if answer == "n" || answer == "no" {
			break
		}
		fmt.Println("I cannot understand thy request, wanderer! Answer (y)es or (no)")
	}
}

func (g Game) applyWeaponAndArmorBonus(ferocityFactor *uint8) {
	if g.player.suit {
		fmt.Println("Your armor increase your chance of success!")
		*ferocityFactor = 3*(*ferocityFactor/4)
	}
	if !g.player.axe && !g.player.sword {
		fmt.Println("You must fight with bare hands!")
		*ferocityFactor += *ferocityFactor/5
	} else if g.player.axe && !g.player.sword {
		fmt.Println("You have only an axe to fight with!")
		*ferocityFactor = 4*(*ferocityFactor/5)
	} else if !g.player.axe && g.player.sword {
		fmt.Println("You must fight with your sword!")
		*ferocityFactor = 3*(*ferocityFactor/4)
	} else { // player has both sword and axe
		prompt := "Which weapon? (1 - AXE, 2 - SWORD)  "
		var choice uint8 = ERROR_CODE
		for fmt.Print(prompt); g.scanner.Scan(); fmt.Print(prompt) {
  		choiceStr := strings.TrimSpace(g.scanner.Text())
			choiceUi8, ok := parseUint8(choiceStr)
			if !ok || choice == 0 || choice > 2 {
				fmt.Println("You don't have this weapon, my master! Say it again!")
				continue
			}
			choice = choiceUi8
			break
		}
		if choice == ERROR_CODE {
			fmt.Println("My master, let me choose thy weapon!")
			rnd := rand.Intn(100)
			if rnd > 50 {
				fmt.Println("You have only an axe to fight with!")
				*ferocityFactor = 4*(*ferocityFactor/5)
			} else {
				fmt.Println("You must fight with your sword!")
				*ferocityFactor = 3*(*ferocityFactor/4)
			}
		} else if choice == 1 {
			fmt.Println("You have only an axe to fight with!")
			*ferocityFactor = 4*(*ferocityFactor/5)
		} else if choice == 2 {
			fmt.Println("You must fight with your sword!")
			*ferocityFactor = 3*(*ferocityFactor/4)
		}
	}
}

func (g *Game) fight() {
	monsterIdx := g.rooms[g.currentRoom].room[6]
	thereIsNoMonster := monsterIdx < 252 
	if thereIsNoMonster {
		fmt.Println("There is nothing to fight here...")
		return
	}

	fmt.Println("Press any key to fight...")
	g.scanner.Scan()

	monster := MONSTERS[monsterIdx]
	ferocityFactor := monster.ferocity
	g.applyWeaponAndArmorBonus(&ferocityFactor)
	time.Sleep(TIMEOUT)

	randomSuccess := func(prob int) bool {
		return rand.Intn(100) > prob
	}

	for {
		if randomSuccess(50) {
			fmt.Println(monster.name, "attacks!")
			time.Sleep(TIMEOUT)
			if randomSuccess(50) {
				fmt.Println("The monster wounds you!")
				g.player.GetWounded()
				g.checkDeath()
			} else {
				fmt.Println("You avoided the attack!")
			}
		} else {
			fmt.Println("You attack!")
			time.Sleep(TIMEOUT)
			if randomSuccess(50) {
				fmt.Println("You manage to hit the monster!")
				ferocityFactor = 5*(ferocityFactor/6)
			} else {
				fmt.Println("The monster prevented the damage!")
			}
		}
		time.Sleep(TIMEOUT)
		if randomSuccess(65) {
			break
		}
	}

	if uint8(rand.Intn(16)) > ferocityFactor {
		fmt.Println("...and you managed to kill the", monster.name, "!!!")
		g.player.monsterKilled++
	} else {
		fmt.Println("The", monster.name, "defeated you!")
		g.player.strength /= 2
	}
	g.rooms[g.currentRoom].room[6] = 0
	time.Sleep(TIMEOUT)
	fmt.Println()
}
