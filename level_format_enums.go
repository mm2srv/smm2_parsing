package smm2_parsing

// Some of these values come from https://github.com/TheGreatRambler/toost/blob/main/src/LevelParser.hpp
// and others come from https://github.com/liamadvance/smm2-documentation/blob/master/Course%20Format.md

type ObjId uint16

const (
	GOOMBA                  ObjId = 0
	KOOPA                   ObjId = 1
	PIRANHA_FLOWER          ObjId = 2
	HAMMER_BRO              ObjId = 3
	BLOCK                   ObjId = 4
	QUESTION_BLOCK          ObjId = 5
	HARD_BLOCK              ObjId = 6
	GROUND                  ObjId = 7
	COIN                    ObjId = 8
	PIPE                    ObjId = 9
	SPRING                  ObjId = 10
	LIFT                    ObjId = 11
	THWOMP                  ObjId = 12
	BULLET_BILL_BLASTER     ObjId = 13
	MUSHROOM_PLATFORM       ObjId = 14
	BOB_OMB                 ObjId = 15
	SEMISOLID_PLATFORM      ObjId = 16
	BRIDGE                  ObjId = 17
	P_SWITCH                ObjId = 18
	POW                     ObjId = 19
	SUPER_MUSHROOM          ObjId = 20
	DONUT_BLOCK             ObjId = 21
	CLOUD                   ObjId = 22
	NOTE_BLOCK              ObjId = 23
	FIRE_BAR                ObjId = 24
	SPINY                   ObjId = 25
	GOAL_GROUND             ObjId = 26
	GOAL                    ObjId = 27
	BUZZY_BEETLE            ObjId = 28
	HIDDEN_BLOCK            ObjId = 29
	LAKITU                  ObjId = 30
	LAKITU_CLOUD            ObjId = 31
	BANZAI_BILL             ObjId = 32
	ONE_UP                  ObjId = 33
	FIRE_FLOWER             ObjId = 34
	SUPER_STAR              ObjId = 35
	LAVA_LIFT               ObjId = 36
	STARTING_BRICK          ObjId = 37
	STARTING_ARROW          ObjId = 38
	MAGIKOOPA               ObjId = 39
	SPIKE_TOP               ObjId = 40
	BOO                     ObjId = 41
	CLOWN_CAR               ObjId = 42
	SPIKES                  ObjId = 43
	BIG_MUSHROOM            ObjId = 44
	SHOE_GOOMBA             ObjId = 45
	DRY_BONES               ObjId = 46
	CANNON                  ObjId = 47
	BLOOPER                 ObjId = 48
	CASTLE_BRIDGE           ObjId = 49
	JUMPING_MACHINE         ObjId = 50
	SKIPSQUEAK              ObjId = 51
	WIGGLER                 ObjId = 52
	FAST_CONVEYOR_BELT      ObjId = 53
	BURNER                  ObjId = 54
	DOOR                    ObjId = 55
	CHEEP_CHEEP             ObjId = 56
	MUNCHER                 ObjId = 57
	ROCKY_WRENCH            ObjId = 58
	TRACK                   ObjId = 59
	LAVA_BUBBLE             ObjId = 60
	CHAIN_CHOMP             ObjId = 61
	BOWSER                  ObjId = 62
	ICE_BLOCK               ObjId = 63
	VINE                    ObjId = 64
	STINGBY                 ObjId = 65
	ARROW                   ObjId = 66
	ONE_WAY                 ObjId = 67
	SAW                     ObjId = 68
	PLAYER                  ObjId = 69
	BIG_COIN                ObjId = 70
	HALF_COLLISION_PLATFORM ObjId = 71
	KOOPA_CAR               ObjId = 72
	CINOBIO                 ObjId = 73
	SPIKE_BALL              ObjId = 74
	STONE                   ObjId = 75
	TWISTER                 ObjId = 76
	BOOM_BOOM               ObjId = 77
	POKEY                   ObjId = 78
	P_BLOCK                 ObjId = 79
	SPRINT_PLATFORM         ObjId = 80
	SMB2_MUSHROOM           ObjId = 81
	DONUT                   ObjId = 82
	SKEWER                  ObjId = 83
	SNAKE_BLOCK             ObjId = 84
	TRACK_BLOCK             ObjId = 85
	CHARVAARGH              ObjId = 86
	SLIGHT_SLOPE            ObjId = 87
	STEEP_SLOPE             ObjId = 88
	REEL_CAMERA             ObjId = 89
	CHECKPOINT_FLAG         ObjId = 90
	SEESAW                  ObjId = 91
	RED_COIN                ObjId = 92
	CLEAR_PIPE              ObjId = 93
	CONVEYOR_BELT           ObjId = 94
	KEY                     ObjId = 95
	ANT_TROOPER             ObjId = 96
	WARP_BOX                ObjId = 97
	BOWSER_JR               ObjId = 98
	ON_OFF_BLOCK            ObjId = 99
	DOTTED_LINE_BLOCK       ObjId = 100
	WATER_MARKER            ObjId = 101
	MONTY_MOLE              ObjId = 102
	FISH_BONE               ObjId = 103
	ANGRY_SUN               ObjId = 104
	SWINGING_CLAW           ObjId = 105
	TREE                    ObjId = 106
	PIRANHA_CREEPER         ObjId = 107
	BLINKING_BLOCK          ObjId = 108
	SOUND_EFFECT            ObjId = 109
	SPIKE_BLOCK             ObjId = 110
	MECHAKOOPA              ObjId = 111
	CRATE                   ObjId = 112
	MUSHROOM_TRAMPOLINE     ObjId = 113
	PORKUPUFFER             ObjId = 114
	CINOBIC                 ObjId = 115
	SUPER_HAMMER            ObjId = 116
	BULLY                   ObjId = 117
	ICICLE                  ObjId = 118
	EXCLAMATION_BLOCK       ObjId = 119
	LEMMY                   ObjId = 120
	MORTON                  ObjId = 121
	LARRY                   ObjId = 122
	WENDY                   ObjId = 123
	IGGY                    ObjId = 124
	ROY                     ObjId = 125
	LUDWIG                  ObjId = 126
	CANNON_BOX              ObjId = 127
	PROPELLER_BOX           ObjId = 128
	GOOMBA_MASK             ObjId = 129
	BULLET_BILL_MASK        ObjId = 130
	RED_POW_BOX             ObjId = 131
	ON_OFF_TRAMPOLINE       ObjId = 132
)

type ClearConId uint32

const (
	CLEARCON_NONE                                                                                              ClearConId = 0
	REACH_THE_GOAL_WITHOUT_LANDING_AFTER_LEAVING_THE_GROUND                                                    ClearConId = 137525990
	REACH_THE_GOAL_AFTER_DEFEATING_AT_LEAST_ALL_MECHAKOOPA                                                     ClearConId = 199585683
	REACH_THE_GOAL_AFTER_DEFEATING_AT_LEAST_ALL_CHEEP_CHEEP                                                    ClearConId = 272349836
	REACH_THE_GOAL_WITHOUT_TAKING_DAMAGE                                                                       ClearConId = 375673178
	REACH_THE_GOAL_AS_BOOMERANG_MARIO                                                                          ClearConId = 426197923
	REACH_THE_GOAL_WHILE_WEARING_A_SHOE                                                                        ClearConId = 436833616
	REACH_THE_GOAL_AS_FIRE_MARIO                                                                               ClearConId = 713979835
	REACH_THE_GOAL_AS_FROG_MARIO                                                                               ClearConId = 744927294
	REACH_THE_GOAL_AFTER_DEFEATING_AT_LEAST_ALL_LARRY                                                          ClearConId = 751004331
	REACH_THE_GOAL_AS_RACCOON_MARIO                                                                            ClearConId = 900050759
	REACH_THE_GOAL_AFTER_DEFEATING_AT_LEAST_ALL_BLOOPER                                                        ClearConId = 947659466
	REACH_THE_GOAL_AS_PROPELLER_MARIO                                                                          ClearConId = 976173462
	REACH_THE_GOAL_WHILE_WEARING_A_PROPELLER_BOX                                                               ClearConId = 994686866
	REACH_THE_GOAL_AFTER_DEFEATING_AT_LEAST_ALL_SPIKE                                                          ClearConId = 998904081
	REACH_THE_GOAL_AFTER_DEFEATING_AT_LEAST_ALL_BOOM_BOOM                                                      ClearConId = 1008094897
	REACH_THE_GOAL_WHILE_HOLDING_A_KOOPA_SHELL                                                                 ClearConId = 1051433633
	REACH_THE_GOAL_AFTER_DEFEATING_AT_LEAST_ALL_PORCUPUFFER                                                    ClearConId = 1061233896
	REACH_THE_GOAL_AFTER_DEFEATING_AT_LEAST_ALL_CHARVAARGH                                                     ClearConId = 1062253843
	REACH_THE_GOAL_AFTER_DEFEATING_AT_LEAST_ALL_BULLET_BILL                                                    ClearConId = 1079889509
	REACH_THE_GOAL_AFTER_DEFEATING_AT_LEAST_ALL_BULLY_BULLIES                                                  ClearConId = 1080535886
	REACH_THE_GOAL_WHILE_WEARING_A_GOOMBA_MASK                                                                 ClearConId = 1151250770
	REACH_THE_GOAL_AFTER_DEFEATING_AT_LEAST_ALL_HOP_CHOPS                                                      ClearConId = 1182464856
	REACH_THE_GOAL_WHILE_HOLDING_A_RED_POW_BLOCK_OR_REACH_THE_GOAL_AFTER_ACTIVATING_AT_LEAST_ALL_RED_POW_BLOCK ClearConId = 1219761531
	REACH_THE_GOAL_AFTER_DEFEATING_AT_LEAST_ALL_BOB_OMB                                                        ClearConId = 1221661152
	REACH_THE_GOAL_AFTER_DEFEATING_AT_LEAST_ALL_SPINY_SPINIES                                                  ClearConId = 1259427138
	REACH_THE_GOAL_AFTER_DEFEATING_AT_LEAST_ALL_BOWSER_MEOWSER                                                 ClearConId = 1268255615
	REACH_THE_GOAL_AFTER_DEFEATING_AT_LEAST_ALL_ANT_TROOPER                                                    ClearConId = 1279580818
	REACH_THE_GOAL_ON_A_LAKITUS_CLOUD                                                                          ClearConId = 1283945123
	REACH_THE_GOAL_AFTER_DEFEATING_AT_LEAST_ALL_BOO                                                            ClearConId = 1344044032
	REACH_THE_GOAL_AFTER_DEFEATING_AT_LEAST_ALL_ROY                                                            ClearConId = 1425973877
	REACH_THE_GOAL_WHILE_HOLDING_A_TRAMPOLINE                                                                  ClearConId = 1429902736
	REACH_THE_GOAL_AFTER_DEFEATING_AT_LEAST_ALL_MORTON                                                         ClearConId = 1431944825
	REACH_THE_GOAL_AFTER_DEFEATING_AT_LEAST_ALL_FISH_BONE                                                      ClearConId = 1446467058
	REACH_THE_GOAL_AFTER_DEFEATING_AT_LEAST_ALL_MONTY_MOLE                                                     ClearConId = 1510495760
	REACH_THE_GOAL_AFTER_PICKING_UP_AT_LEAST_ALL_1_UP_MUSHROOM                                                 ClearConId = 1656179347
	REACH_THE_GOAL_AFTER_DEFEATING_AT_LEAST_ALL_HAMMER_BRO                                                     ClearConId = 1665820273
	REACH_THE_GOAL_AFTER_HITTING_AT_LEAST_ALL_P_SWITCH_OR_REACH_THE_GOAL_WHILE_HOLDING_A_P_SWITCH              ClearConId = 1676924210
	REACH_THE_GOAL_AFTER_ACTIVATING_AT_LEAST_ALL_POW_BLOCK_OR_REACH_THE_GOAL_WHILE_HOLDING_A_POW_BLOCK         ClearConId = 1715960804
	REACH_THE_GOAL_AFTER_DEFEATING_AT_LEAST_ALL_ANGRY_SUN                                                      ClearConId = 1724036958
	REACH_THE_GOAL_AFTER_DEFEATING_AT_LEAST_ALL_POKEY                                                          ClearConId = 1730095541
	REACH_THE_GOAL_AS_SUPERBALL_MARIO                                                                          ClearConId = 1780278293
	REACH_THE_GOAL_AFTER_DEFEATING_AT_LEAST_ALL_POM_POM                                                        ClearConId = 1839897151
	REACH_THE_GOAL_AFTER_DEFEATING_AT_LEAST_ALL_PEEPA                                                          ClearConId = 1969299694
	REACH_THE_GOAL_AFTER_DEFEATING_AT_LEAST_ALL_LAKITU                                                         ClearConId = 2035052211
	REACH_THE_GOAL_AFTER_DEFEATING_AT_LEAST_ALL_LEMMY                                                          ClearConId = 2038503215
	REACH_THE_GOAL_AFTER_DEFEATING_AT_LEAST_ALL_LAVA_BUBBLE                                                    ClearConId = 2048033177
	REACH_THE_GOAL_WHILE_WEARING_A_BULLET_BILL_MASK                                                            ClearConId = 2076496776
	REACH_THE_GOAL_AS_BIG_MARIO                                                                                ClearConId = 2089161429
	REACH_THE_GOAL_AS_CAT_MARIO                                                                                ClearConId = 2111528319
	REACH_THE_GOAL_AFTER_DEFEATING_AT_LEAST_ALL_GOOMBA_GALOOMBA                                                ClearConId = 2131209407
	REACH_THE_GOAL_AFTER_DEFEATING_AT_LEAST_ALL_THWOMP                                                         ClearConId = 2139645066
	REACH_THE_GOAL_AFTER_DEFEATING_AT_LEAST_ALL_IGGY                                                           ClearConId = 2259346429
	REACH_THE_GOAL_WHILE_WEARING_A_DRY_BONES_SHELL                                                             ClearConId = 2549654281
	REACH_THE_GOAL_AFTER_DEFEATING_AT_LEAST_ALL_SLEDGE_BRO                                                     ClearConId = 2694559007
	REACH_THE_GOAL_AFTER_DEFEATING_AT_LEAST_ALL_ROCKY_WRENCH                                                   ClearConId = 2746139466
	REACH_THE_GOAL_AFTER_GRABBING_AT_LEAST_ALL_50_COIN                                                         ClearConId = 2749601092
	REACH_THE_GOAL_AS_FLYING_SQUIRREL_MARIO                                                                    ClearConId = 2855236681
	REACH_THE_GOAL_AS_BUZZY_MARIO                                                                              ClearConId = 3036298571
	REACH_THE_GOAL_AS_BUILDER_MARIO                                                                            ClearConId = 3074433106
	REACH_THE_GOAL_AS_CAPE_MARIO                                                                               ClearConId = 3146932243
	REACH_THE_GOAL_AFTER_DEFEATING_AT_LEAST_ALL_WENDY                                                          ClearConId = 3174413484
	REACH_THE_GOAL_WHILE_WEARING_A_CANNON_BOX                                                                  ClearConId = 3206222275
	REACH_THE_GOAL_AS_LINK                                                                                     ClearConId = 3314955857
	REACH_THE_GOAL_WHILE_YOU_HAVE_SUPER_STAR_INVINCIBILITY                                                     ClearConId = 3342591980
	REACH_THE_GOAL_AFTER_DEFEATING_AT_LEAST_ALL_GOOMBRAT_GOOMBUD                                               ClearConId = 3346433512
	REACH_THE_GOAL_AFTER_GRABBING_AT_LEAST_ALL_10_COIN                                                         ClearConId = 3348058176
	REACH_THE_GOAL_AFTER_DEFEATING_AT_LEAST_ALL_BUZZY_BEETLE                                                   ClearConId = 3353006607
	REACH_THE_GOAL_AFTER_DEFEATING_AT_LEAST_ALL_BOWSER_JR                                                      ClearConId = 3392229961
	REACH_THE_GOAL_AFTER_DEFEATING_AT_LEAST_ALL_KOOPA_TROOPA                                                   ClearConId = 3437308486
	REACH_THE_GOAL_AFTER_DEFEATING_AT_LEAST_ALL_CHAIN_CHOMP                                                    ClearConId = 3459144213
	REACH_THE_GOAL_AFTER_DEFEATING_AT_LEAST_ALL_MUNCHER                                                        ClearConId = 3466227835
	REACH_THE_GOAL_AFTER_DEFEATING_AT_LEAST_ALL_WIGGLER                                                        ClearConId = 3481362698
	REACH_THE_GOAL_AS_SMB2_MARIO                                                                               ClearConId = 3513732174
	REACH_THE_GOAL_IN_A_KOOPA_CLOWN_CAR_JUNIOR_CLOWN_CAR                                                       ClearConId = 3649647177
	REACH_THE_GOAL_AS_SPINY_MARIO                                                                              ClearConId = 3725246406
	REACH_THE_GOAL_IN_A_KOOPA_TROOPA_CAR                                                                       ClearConId = 3730243509
	REACH_THE_GOAL_AFTER_DEFEATING_AT_LEAST_ALL_PIRANHA_PLANT_JUMPING_PIRANHA_PLANT                            ClearConId = 3748075486
	REACH_THE_GOAL_AFTER_DEFEATING_AT_LEAST_ALL_DRY_BONES                                                      ClearConId = 3797704544
	REACH_THE_GOAL_AFTER_DEFEATING_AT_LEAST_ALL_STINGBY_STINGBIES                                              ClearConId = 3824561269
	REACH_THE_GOAL_AFTER_DEFEATING_AT_LEAST_ALL_PIRANHA_CREEPER                                                ClearConId = 3833342952
	REACH_THE_GOAL_AFTER_DEFEATING_AT_LEAST_ALL_FIRE_PIRANHA_PLANT                                             ClearConId = 3842179831
	REACH_THE_GOAL_AFTER_BREAKING_AT_LEAST_ALL_CRATES                                                          ClearConId = 3874680510
	REACH_THE_GOAL_AFTER_DEFEATING_AT_LEAST_ALL_LUDWIG                                                         ClearConId = 3974581191
	REACH_THE_GOAL_AS_SUPER_MARIO                                                                              ClearConId = 3977257962
	REACH_THE_GOAL_AFTER_DEFEATING_AT_LEAST_ALL_SKIPSQUEAK                                                     ClearConId = 4042480826
	REACH_THE_GOAL_AFTER_GRABBING_AT_LEAST_ALL_COIN                                                            ClearConId = 4116396131
	REACH_THE_GOAL_AFTER_DEFEATING_AT_LEAST_ALL_MAGIKOOPA                                                      ClearConId = 4117878280
	REACH_THE_GOAL_AFTER_GRABBING_AT_LEAST_ALL_30_COIN                                                         ClearConId = 4122555074
	REACH_THE_GOAL_AS_BALLOON_MARIO                                                                            ClearConId = 4153835197
	REACH_THE_GOAL_WHILE_WEARING_A_RED_POW_BOX                                                                 ClearConId = 4172105156
	REACH_THE_GOAL_WHILE_RIDING_YOSHI                                                                          ClearConId = 4209535561
	REACH_THE_GOAL_AFTER_DEFEATING_AT_LEAST_ALL_SPIKE_TOP                                                      ClearConId = 4269094462
	REACH_THE_GOAL_AFTER_DEFEATING_AT_LEAST_ALL_BANZAI_BILL                                                    ClearConId = 4293354249
)

type ClearConCategory uint8

const (
	CATEGORY_NONE    ClearConCategory = 0
	CATEGORY_PARTS   ClearConCategory = 1
	CATEGORY_STATUS  ClearConCategory = 2
	CATEGORY_ACTIONS ClearConCategory = 3
)

type GameVersion uint32

const (
	V1_0_0   GameVersion = 0
	V1_0_1   GameVersion = 1
	V1_1_0   GameVersion = 2
	V2_0_0   GameVersion = 3
	V3_0_0   GameVersion = 4
	V3_0_1   GameVersion = 5
	VUNKNOWN GameVersion = 33
)

type CourseTheme uint8

const (
	OVERWORLD   CourseTheme = 0
	UNDERGROUND CourseTheme = 1
	CASTLE      CourseTheme = 2
	AIRSHIP     CourseTheme = 3
	UNDERWATER  CourseTheme = 4
	GHOST_HOUSE CourseTheme = 5
	SNOW        CourseTheme = 6
	DESERT      CourseTheme = 7
	SKY         CourseTheme = 8
	FOREST      CourseTheme = 9
)

type AutoscrollSpeed uint8

const (
	AUTOSCROLL_X1 AutoscrollSpeed = 0
	AUTOSCROLL_X2 AutoscrollSpeed = 1
	AUTOSCROLL_X3 AutoscrollSpeed = 2
)

type AutoscrollType uint8

const (
	AUTOSCROLL_NONE AutoscrollType = 0
	SLOW            AutoscrollType = 1
	NORMAL          AutoscrollType = 2
	FAST            AutoscrollType = 3
	CUSTOM          AutoscrollType = 4
)

type BoundaryType uint8

const (
	BUILT_ABOVE_LINE BoundaryType = 0
	BUILT_BELOW_LINE BoundaryType = 0
)

type OrientationType uint8

const (
	HORIZONTAL OrientationType = 0
	VERTICAL   OrientationType = 1
)

type LiquidType uint8

const (
	STATIC             LiquidType = 0
	RISING_OR_FALLING  LiquidType = 1
	RISING_AND_FALLING LiquidType = 2
)

type LiquidSpeed uint8

const (
	NONE      LiquidSpeed = 0
	LIQUID_X1 LiquidSpeed = 1
	LIQUID_X2 LiquidSpeed = 2
	LIQUID_X3 LiquidSpeed = 3
)

type SoundId uint8

const (
	SHOCK                SoundId = 0
	CLATTER              SoundId = 1
	KICK                 SoundId = 2
	APPLAUSE             SoundId = 3
	GLORY                SoundId = 4
	PUNCH                SoundId = 5
	LAUGHTER             SoundId = 6
	BABY                 SoundId = 7
	DING_DONG            SoundId = 8
	BOSS_MUSIC           SoundId = 9
	HEARTBEAT            SoundId = 10
	SCREAM               SoundId = 11
	DRAMA                SoundId = 12
	JUMP                 SoundId = 13
	CHEER                SoundId = 14
	DOOM                 SoundId = 15
	FIREWORKS            SoundId = 16
	HONK_HONK            SoundId = 17
	BZZT                 SoundId = 18
	BONUS_MUSIC          SoundId = 19
	SILENCE              SoundId = 20
	UNKNOWN1             SoundId = 21
	UNKNOWN2             SoundId = 22
	PARTY_POPPERINGS     SoundId = 23
	BOOO                 SoundId = 24
	GUFFAW               SoundId = 25
	NEAR_MISS            SoundId = 26
	UNKNOWN3             SoundId = 27
	UNKNOWN4             SoundId = 28
	OINK                 SoundId = 29
	KUH_THUNK            SoundId = 30
	BEEP                 SoundId = 31
	NINJA_ATTACKGERS     SoundId = 32
	UNKNOWN5             SoundId = 33
	UNKNOWN6             SoundId = 34
	ZAP                  SoundId = 35
	FLASH                SoundId = 36
	YEAH                 SoundId = 37
	AWW                  SoundId = 38
	UNKNOWN7             SoundId = 39
	UNKNOWN8             SoundId = 40
	AUDIENCE             SoundId = 41
	SCATTING             SoundId = 42
	SPARK                SoundId = 43
	TRADITIONAL          SoundId = 44
	ELECTRIC_GUITAR      SoundId = 45
	TWISTY_TURNY         SoundId = 46
	WOOZY                SoundId = 47
	FINAL_BOSS           SoundId = 48
	PEACEFUL             SoundId = 49
	HORROR               SoundId = 50
	SUPER_MARIO_GALAXY   SoundId = 51
	SUPER_MARIO_64       SoundId = 52
	SUPER_MARIO_SUNSHINE SoundId = 53
	SUPER_MARIO_KART     SoundId = 54
	UNKNOWN9             SoundId = 55
)
