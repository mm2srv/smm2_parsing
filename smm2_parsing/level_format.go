package smm2_parsing

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"log"
	"os"

	"honnef.co/go/spew"
)

type Header struct {
	YStart                  uint8  // Starting Y position of level
	YGoal                   uint8  // Y position of goal (in tiles)
	XGoal                   uint16 //	X position of goal (in tiles * 10)
	TimeLimit               uint16 // Time limit; defaults to 300
	ClearConditionMagnitude uint16 // Target amount required for clear condition
	CreationYear            uint16 //	Creation year
	CreationMonth           uint8  //	Creation month
	CreationDay             uint8  //	Creation day
	CreationHour            uint8  //	Creation hour
	CreationMinute          uint8  //	Creation minute
	AutoscrollSpeed         uint8
	ClearConditionCategory  uint8  //	Clear condition type (more below)
	ClearConditionObject    uint32 // Clear condition object (CRC32 of more below)
	UnkGameVer              uint32
	ManagementFlags         uint32 // Management flags: &1 always seems to be set, &2 shows that a level has passed its Clear Check, &0x10 shows that the level may not be uploaded
	ClearAttemmpts          uint32
	ClearCheckTime          uint32 // Time taken in Clear Check (units unknown), or 0xFFFFFFFF if the level has not been cleared
	CreationId              uint32 // Initialised to a random value when a level is created
	UploadId                uint64
	GameVersion             uint32
	Unk1                    [0xBD]byte
	GameStyle               [0x2]byte // Game style: M1, M3, MW, WU, 3W
	Unk2                    uint8
	Name                    [0x42]byte // Course name, null-terminated, UCS-2
	Description             [0xCA]byte //	Course description, null-terminated, UCS-2 (game only lets you enter up to 75, but there's space for 100)
}

type Object struct {
	X      uint32
	Y      uint32
	Unk1   uint16
	Width  uint8
	Height uint8
	Flag   uint32
	CFlag  uint32
	Ex     uint32
	Id     uint16
	CId    uint16
	LId    uint16
	SId    uint16
}

type Sound struct {
	Id   uint8
	X    uint8
	Y    uint8
	Unk1 uint8
}

type SnakeNode struct {
	Index     uint16
	Direction uint16
	Id        uint32
}

type Snake struct {
	Index     uint8
	NodeCount uint8
	Unk1      uint16
	Nodes     [120]SnakeNode
}

type ClearPipeNode struct {
	Type      uint8
	Index     uint8
	X         uint8
	Y         uint8
	Width     uint8
	Height    uint8
	Unk1      uint8
	Direction uint8
}

type ClearPipe struct {
	Index     uint8
	NodeCount uint8
	Unk       uint16
	Nodes     [36]ClearPipeNode
}

type PiranhaCreeperNode struct {
	Unk1      uint8
	Direction uint8
	Unk2      uint16
}

type PiranhaCreeper struct {
	Unk1      uint8
	Index     uint8
	NodeCount uint8
	Unk2      uint8
	Nodes     [20]PiranhaCreeperNode
}

type ExclamationBlockNode struct {
	Unk1      uint8
	Direction uint8
	Unk2      uint16
}

type ExclamationBlock struct {
	Unk1      uint8
	Index     uint8
	NodeCount uint8
	Unk2      uint8
	Nodes     [10]ExclamationBlockNode
}

type TrackBlockNode struct {
	Unk1      uint8
	Direction uint8
	Unk2      uint16
}

type TrackBlock struct {
	Unk1      uint8
	Index     uint8
	NodeCount uint8
	Unk2      uint8
	Nodes     [10]TrackBlockNode
}

type Ground struct {
	X            uint8
	Y            uint8
	Id           uint8
	BackgroundId uint8
}

type Track struct {
	Unk1  uint16
	Flags uint8
	X     uint8
	Y     uint8
	Type  uint8
	LId   uint16
	Unk2  uint16
	Unk3  uint16
}

type Icicle struct {
	X    uint8
	Y    uint8
	Type uint8
	Unk1 uint8
}

type LevelArea struct {
	Theme                     uint8
	AutoscrollType            uint8
	BoundaryType              uint8
	Orientation               uint8
	LiquidEndHeight           uint8
	LiquidMode                uint8
	LiquidSpeed               uint8
	LiquidStartHeight         uint8
	BoundaryRight             uint32
	BoundaryTop               uint32
	BoundaryLeft              uint32
	BoundaryBottom            uint32
	UnkFlag                   uint32
	ObjectCount               uint32
	SoundEffectCount          uint32
	SnakeBlockCount           uint32
	ClearPipeCount            uint32
	PiranhaCreeperCount       uint32
	ExclamationMarkBlockCount uint32
	TrackBlockCount           uint32
	Unk1                      uint32
	GroundCount               uint32
	TrackCount                uint32
	IceCount                  uint32
	Objects                   [2600]Object
	Sounds                    [300]Sound
	Snakes                    [5]Snake
	ClearPipes                [200]ClearPipe
	PiranhaCreepers           [10]PiranhaCreeper
	ExclamationBlocks         [10]ExclamationBlock
	TrackBlocks               [10]TrackBlock
	Ground                    [4000]Ground
	Tracks                    [1500]Track
	Icicles                   [300]Icicle
	Unk2                      [0xDBC]byte
}

type BCD struct {
	//VersionNumber1 uint32  // Assumed version number -- must be 1 in v1.0.0
	//VersionNumber2 uint16  // Assumed version number -- must be 0x10 in v1.0.0
	//Padding1       [2]byte // 2 empty bytes
	//CRC32          uint32  // CRC32 of the decrypted level file from offset 0x10 to 0x5BFD0 (non-inclusive)
	//Magic          [4]byte // Magic SCDL (53 53 44 4C)
	Header    Header // header 0x200
	OverWorld LevelArea
	SubWorld  LevelArea
}

func LoadBCD(buf []byte) (*BCD, error) {
	s := &BCD{}
	err := s.Load(buf)
	return s, err
}

func (s *BCD) Load(buf []byte) error {
	var err error
	buf, err = DecryptLevel(buf)
	if err != nil {
		return err
	}
	return binary.Read(bytes.NewReader(buf), binary.LittleEndian, s)
}

func (s *BCD) LoadDecrypted(buf []byte) error {
	return binary.Read(bytes.NewReader(buf), binary.LittleEndian, s)
}

func (s *BCD) Save() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := binary.Write(buf, binary.LittleEndian, s)
	if err != nil {
		return []byte{}, err
	}
	out, err := EncryptLevel(buf.Bytes())
	return out, err
}

func (s *BCD) SaveDecrypted() ([]byte, error) {
	buf := &bytes.Buffer{}
	err := binary.Write(buf, binary.LittleEndian, s)
	if err != nil {
		return []byte{}, err
	}
	return buf.Bytes(), err
}

func RemoveUploadedFlag(rawBCD []byte) ([]byte, error) {
	level, err := LoadBCD(rawBCD)
	if err != nil {
		return []byte{}, err
	}

	level.Header.UploadId = 0
	level.Header.ManagementFlags = (level.Header.ManagementFlags &^ (1 << 16))
	level.Header.ManagementFlags = (level.Header.ManagementFlags &^ (1 << 4))
	level.Header.ManagementFlags = (level.Header.ManagementFlags &^ (1 << 6))

	newBCD, err := level.Save()
	if err != nil {
		return []byte{}, err
	}
	return newBCD, nil
}

func debug_upload_ban() error {
	buf, err := os.ReadFile("data/level_tests/clean_short.bcd")
	if err != nil {
		return err
	}

	level, err := LoadBCD(buf)
	if err != nil {
		return err
	}

	//spew.Dump(level.Header)
	fmt.Println(fmt.Sprintf("%032b", level.Header.ManagementFlags), level.Header.ManagementFlags)

	buf, err = os.ReadFile("data/level_tests/upload_banned.bcd")
	if err != nil {
		return err
	}

	level, err = LoadBCD(buf)
	if err != nil {
		return err
	}

	fmt.Println(fmt.Sprintf("%032b", level.Header.ManagementFlags), level.Header.ManagementFlags)
	level.Header.ManagementFlags = (level.Header.ManagementFlags &^ (1 << 4))
	level.Header.ManagementFlags = (level.Header.ManagementFlags &^ (1 << 6))
	fmt.Println(fmt.Sprintf("%032b", level.Header.ManagementFlags), level.Header.ManagementFlags)

	return nil
}

func parse() error {
	//buf, err := ioutil.ReadFile("data/level_tests/course_data_095.bcd")
	buf, err := os.ReadFile("data/level_tests/clean_short.bcd")
	if err != nil {
		return err
	}

	level, err := LoadBCD(buf)
	if err != nil {
		return err
	}
	if false {
		level.Header.ClearConditionCategory = 2
		level.Header.ClearConditionObject = 46219146
		level.OverWorld.Objects[0].Id = 75 // stone
		//level.OverWorld.Objects[0].Id = 5 // ? block
	}

	spew.Dump(level.Header, level.OverWorld.Objects[:int(level.OverWorld.ObjectCount)])
	out, err := level.Save()
	if err != nil {
		return err
	}

	err = os.WriteFile("data/level_tests/stonetest.bcd", out, 0600)
	if err != nil {
		return err
	}

	if false {
		// round-trip check
		level2, err := LoadBCD(out)
		if err != nil {
			return err
		}
		spew.Dump(level2.Header, level2.OverWorld.Objects[:int(level2.OverWorld.ObjectCount)])
		out2, err := level2.Save()
		if err != nil {
			return err
		}

		spew.Dump(bytes.Equal(out, out2))
	}

	return err
}

func main_off() {
	err := debug_upload_ban()
	if err != nil {
		log.Fatal(err)
	}
}
