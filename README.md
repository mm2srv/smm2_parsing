# smm2_parsing
Super Mario Maker 2 file parsing in go. Includes level, thumbnail and replay parsing.

## API
### Level encryption
```go
func DecryptLevel(buf []byte) ([]byte, error)
```
Decrypt BCD binary.

```go
func EncryptLevel(buf []byte) ([]byte, error)
```
Encrypt BCD binary into format accepted by the game.

```go
func Compress(data []byte) ([]byte, error)
```
Zlib compression. Zlib compression on a decrypted BCD generally makes it 1% of the size.

```go
func Decompress(buf []byte) ([]byte, error)
```
Zlib decompression.

### Level parsing
```go
func (s *BCD) Load(buf []byte) error
```
Load encrypted BCD into BCD struct.

```go
func (s *BCD) LoadDecrypted(buf []byte) error
```
Load decrypted BCD into BCD struct.

```go
func (s *BCD) Save() ([]byte, error)
```
Save encrypted BCD from BCD struct.

```go
func (s *BCD) SaveDecrypted() ([]byte, error)
```
Save decrypted BCD from BCD struct.

### Thumbnail encryption
```go
func EncryptJpegThumbnail(buf []byte) ([]byte, error)
```
Encrypt thumbnail to be accepted by the game. Will ignore already encrypted thumbnails. Encrypted thumbnails contain some data at the end but are still viewable by standard image viewers.

```go
func RepackThumbnailUntilFit(buf []byte) ([]byte, error)
```
Thumbnails accepted by the game must be less than 0x1BF9C bytes, this continually tries to repack the JPEG until it is under this limit and returns an error if quality 20 is not enough to make the JPEG fit under the size limit.

```go
func UnpackJpegThumbnail(buf []byte) ([]byte, error)
```
Unpack arbitrary image into JPEG. Simply performs `image.Decode(reader)`. 

```go
func RepackWorldThumbnailSilly(buf []byte) ([]byte, error)
```
Some thumbnails fail for some reason.

### Replay parsing
```go
func (s *Replay) Load(buf []byte) error
```
Load WR, first clear or upload replay into list.

```go
func (s *Replay) GetTASText() string
```
Get nx-tas compatible tas script from replay. Due to slight differences there will be desyncs.

## Examples
```go
import (
	"os"
	"github.com/TheGreatRambler/smm2_parsing"
)

encrypted, err := os.ReadFile("test.bcd")
if err != nil {
    return err
}
decrypted, err := smm2_parsing.DecryptLevel(encrypted)
if err != nil {
    return err
}
compressed, err := smm2_parsing.Compress(decrypted)
if err != nil {
    return err
}
err = os.WriteFile("test_compressed.bcd", compressed, 0644)
if err != nil {
    return err
}
```
Read encrypted BCD from disk, decrypt it, compress it and write it to disk.

```go
import (
	"os"
	"github.com/TheGreatRambler/smm2_parsing"
)

decrypted, err := os.ReadFile("thumbnail.jpg")
if err != nil {
    return err
}
encrypted, err := smm2_parsing.EncryptJpegThumbnail(decrypted)
if err != nil {
    return err
}
err = os.WriteFile("thumbnail_encrypted.jpg", encrypted, 0644)
if err != nil {
    return err
}
```
Read thumbnail from disk and encrypt it so it is accepted by the game.

```go
import (
	"os"
	"github.com/TheGreatRambler/smm2_parsing"
)

decrypted, err := os.ReadFile("decrypted.bcd")
if err != nil {
    return err
}
bcd := &smm2_parsing.BCD{}
if bcd.LoadDecrypted(decrypted) != nil {
    return err
}

bcd.Header.UploadId = 0
bcd.Header.ManagementFlags = (level.Header.ManagementFlags &^ (1 << 16))
bcd.Header.ManagementFlags = (level.Header.ManagementFlags &^ (1 << 4))
bcd.Header.ManagementFlags = (level.Header.ManagementFlags &^ (1 << 6))

newDecrypted, err := bcd.Save()
if err != nil {
	return err
}
err = os.WriteFile("no_upload_flag.bcd", newDecrypted, 0644)
if err != nil {
    return err
}
```
Read decrypted BCD from disk, remove upload flag and write it back to disk.

```go
import (
	"os"
    "fmt"
	"github.com/TheGreatRambler/smm2_parsing"
)

replayFile, err := os.ReadFile("replay.bin")
if err != nil {
    return err
}
replay := &smm2_parsing.Replay{}
if replay.Load(replayFile) != nil {
    return err
}
for _, entry := range replay.entries {
    // Iterate through replay entries
    if entry.entry_type == Time {
        fmt.Printf("Time: %d frames\n", entry.frames)
	} else if entry.entry_type == Joysticks {
        fmt.Printf("Joysticks: %d x %d y\n", entry.joystick_x, entry.joystick_y)
	} else if entry.entry_type == Inputs {
        // Iterate through inputs to construct a string
        var inputString string
		if len(entry.inputs) == 0 {
			inputString = "NONE"
		} else {
            for i, input := range entry.inputs {
                inputString += replay.InputToName(input)
                if i != len(entry.inputs) - 1 {
                    inputString += " "
                }
            }
		}
		fmt.Printf("Inputs: %s\n", inputString)
	}
}
```
Load WR, first clear or upload replay into list and print that list.

## Credits
* [simontime](https://github.com/simontime) for [C level decryption code](https://github.com/simontime/SMM2CourseDecryptor)
* [Liam](https://github.com/liamadvance) for [documentation on the level format](https://github.com/liamadvance/smm2-documentation)
* [Ji](https://github.com/JiXiaomai) for [creating working code to parse the level format](https://github.com/JiXiaomai/SMM2LevelViewer)
* Kramer, [Wizulus](https://twitter.com/wizulus) and [TheGreatRambler](https://twitter.com/tgr_code/) for making this library