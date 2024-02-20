package game

import "time"

const (
  NORTH = 0
  SOUTH = 1
  EAST = 2
  WEST = 3
  UP = 4
  DOWN = 5
  FIGHT = 6
  RUN = 7
  MAGIC_AMULET = 8
  INVENTORY = 9
  QUIT = 10
  PICK_UP = 11
  CONSUME = 12

  INPUT_MAP_FILE = "assets/map.txt"
  PLAYER_NAME_REGEX = `^[a-zA-Z]{3,15}$`
  TIMEOUT = 500 * time.Millisecond
  ERROR_CODE = 255

  NR_TREASURES = 4
  NR_MONSTERS = 4
  INITIAL_STRENGTH = 100
  INITIAL_WEALTH = 75
  RUN_SUCCESS_RATE = 80
  FOOD_STRENGTH_VALUE = 5
  ENERGY_CONSUMPTION = 5
  WOUND_DAMAGE = 5
  ROOM_CONTENT = 6

  TORCH_ITEM = 1
  TORCH_PRICE = 15
  AXE_ITEM = 2
  AXE_PRICE = 10
  SWORD_ITEM = 3
  SWORD_PRICE = 20
  FOOD_ITEM = 4
  FOOD_PRICE = 2
  AMULET_ITEM = 5
  AMULET_PRICE = 30
  ARMOR_ITEM = 6
  ARMOR_PRICE = 50

  PRIVATE_MEETING_ROOM = 3
  START = 5
  LIFT = 8
  REAR_VESTIBULE = 9
  EXIT = 10
  TREASURY = 15

  PROMPT = "What do you want to do?  "

  BANNER = `
m     m mmmmmm mmmmm  mmmmmmm     m  mmmm  m      mmmmmm
#  #  # #      #   "# #     #  #  # m"  "m #      #     
" #"# # #mmmmm #mmmm" #mmmmm" #"# # #    # #      #mmmmm
 ## ##" #      #   "m #      ## ##" #    # #      #     
 #   #  #mmmmm #    " #mmmmm #   #   #mm#  #mmmmm #     
                                                        
                                                        
                                   
                 mm   mm   m mmmm  
                 ##   #"m  # #   "m
                #  #  # #m # #    #
                #mm#  #  # # #    #
               #    # #   ## #mmm" 
                                   
                                   
                                                        
m     m   mm   mm   m mmmm   mmmmmm mmmmm  mmmmmm mmmmm 
#  #  #   ##   #"m  # #   "m #      #   "# #      #   "#
" #"# #  #  #  # #m # #    # #mmmmm #mmmm" #mmmmm #mmmm"
 ## ##"  #mm#  #  # # #    # #      #   "m #      #   "m
 #   #  #    # #   ## #mmm"  #mmmmm #    " #mmmmm #    "


`
)

var MONSTERS = map[uint8]Monster{
  255: {name: "Ferocious Werewolf", ferocity: 5},
  254: {name: "Fanatical Fleshgorg", ferocity: 10},
  253: {name: "Maloventy Maldemer", ferocity: 15},
  252: {name: "Devastating Ice-Dragon", ferocity: 20},
}

var ITEM_PRICES = map[uint8]uint16{
  TORCH_ITEM: TORCH_PRICE,
  AXE_ITEM: AXE_PRICE,
  SWORD_ITEM: SWORD_PRICE,
  FOOD_ITEM: FOOD_PRICE,
  AMULET_ITEM: AMULET_PRICE,
  ARMOR_ITEM: ARMOR_PRICE,
}
