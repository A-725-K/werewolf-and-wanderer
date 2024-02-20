package game

import "fmt"

func (p *Player) GetWounded() {
  p.strength -= WOUND_DAMAGE
}

func (p *Player) ConsumeEnergy() bool {
  p.strength -= ENERGY_CONSUMPTION
  if p.strength <= 10 {
    fmt.Printf("WARNING, %s, your strength\nis running low!\n", p.playerName)
    if p.strength < 1 {
      return false
    }
  }
  return true
}

func (p *Player) PickUpTreasure(qty uint16) {
  p.wealth += qty
}

func (p *Player) ConsumeFood(qty uint16) {
  p.food -= qty
  p.strength += FOOD_STRENGTH_VALUE*int16(qty)
}

func (p *Player) BuyItem(itemN uint8, qty uint8) {
  price, ok := ITEM_PRICES[itemN]
  if !ok {
    panic("This item does not exist!!!")
  }

  if price * uint16(qty) > p.wealth {
    fmt.Println("Don't try to fool me, wanderer! You cannot afford this much!")
    return
  }

  var itemStr string
  p.wealth -= price * uint16(qty)
  switch itemN {
  case TORCH_ITEM:
    itemStr = "a torch"
    p.light = true
  case AXE_ITEM:
    itemStr = "an axe"
    p.axe = true
  case SWORD_ITEM:
    itemStr = "a sword"
    p.sword = true
  case AMULET_ITEM:
    itemStr = "an amulet"
    p.amulet = true
  case ARMOR_ITEM:
    itemStr = "an armor"
    p.suit = true
  case FOOD_ITEM:
    itemStr = fmt.Sprintf("%d units of food", qty)
    p.food += uint16(qty)
  default:
    panic("Cannot buy this item!!!")
  }

  fmt.Printf("You bought %s!\n\n", itemStr)
}

func (p Player) DisplayStatus() {
	fmt.Printf("\n%s, your strength is %v.\n", p.playerName, p.strength)
	if p.wealth > 0 {
		fmt.Printf("You have $%v.\n", p.wealth)
	}
	if p.food > 0 {
		fmt.Printf("Your provisioning sack holds %v units of food.\n", p.food)
	}
	if p.suit {
		fmt.Println("You are wearing an armor.")
	}
	if p.isCarryingItems() {
	  fmt.Println("You are carrying:")
	  if p.axe {
		  fmt.Println("  - an axe")
	  }
	  if p.sword {
		  fmt.Println("  - a sword")
	  }
	  if p.amulet {
		  fmt.Println("  - a magic amulet")
	  }
	  fmt.Println()
	}
}

func (p Player) isCarryingItems() bool {
  return p.food > 0 || p.suit || p.axe || p.amulet || p.sword
}
