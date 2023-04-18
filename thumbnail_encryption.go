package smm2_parsing

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/binary"
	"fmt"
	"image"
	"image/jpeg"
)

var thumbnailTable = []uint32{
	0x39b399d2, 0xfae40b38, 0x851bc213, 0x8cb4e3d9,
	0x7ed1c46a, 0xe8050462, 0xd8d24f76, 0xb52886fc,
	0x67890bf0, 0xf5329cb0, 0xd597fb28, 0x2b8ee0ea,
	0x47574c51, 0x0f7569d9, 0xcf1163ae, 0xe4a153bf,
	0xd1fae468, 0xd4c64738, 0x360106f5, 0xdd7eb113,
	0xc296f3e2, 0x2c58f258, 0x79b554e1, 0x85df9d06,
	0xaa307330, 0x01410f69, 0xb2f2c573, 0x82b93eb1,
	0xf351a11c, 0x63098693, 0x885b5da5, 0x8872a8ed,
	0xacd9cb13, 0xed7fbcad, 0xe6a41ec2, 0x5f44e79f,
	0x8346f5b5, 0x389fe6ed, 0x507124b5, 0xe9b23eaa,
	0x577113f0, 0xa95ed917, 0x2f62d158, 0x47843f86,
	0xc65637d0, 0x2f272052, 0xba4a4cc4, 0xb5f146f6,
	0x501b87a7, 0x51fc3a93, 0x6ede3f02, 0x3d265728,
	0x9b809440, 0x75b89229, 0xf6a280cc, 0x8537fa68,
	0x5b5ed19a, 0x6fc05bb6, 0xf4ef5261, 0xaa1b7d4f,
	0xfcb26110, 0x00ad3d74, 0xc0e73a4b, 0xf132e7c7,
}

func UnpackJpegThumbnail(buf []byte) ([]byte, error) {
	if len(buf) != 0x1c000 {
		return []byte{}, fmt.Errorf("invalid buf size %d != %d", len(buf), 0x1c000)
	}

	reader := bytes.NewReader(buf)
	img, _, err := image.Decode(reader)
	if err != nil {
		return nil, err
	}

	out := &bytes.Buffer{}
	err = jpeg.Encode(out, img, &jpeg.Options{Quality: 65})
	if err != nil {
		return nil, err
	}

	return out.Bytes(), nil
}

// no clue why but (some) dump world thumbnails fail otherwise
func RepackWorldThumbnailSilly(buf []byte) ([]byte, error) {
	reader := bytes.NewReader(buf)
	img, _, err := image.Decode(reader)
	if err != nil {
		return nil, err
	}

	out := &bytes.Buffer{}
	err = jpeg.Encode(out, img, &jpeg.Options{Quality: 65})
	if err != nil {
		return nil, err
	}

	return out.Bytes(), nil
}

func RepackThumbnailUntilFit(buf []byte) ([]byte, error) {
	fmt.Printf("WARNING repacking thumbnail because jpeg is too large %d > %d\n", len(buf), 0x1BF9C)
	reader := bytes.NewReader(buf)
	img, _, err := image.Decode(reader)
	if err != nil {
		return nil, err
	}

	out := &bytes.Buffer{}
	for _, quality := range []int{95, 85, 75, 65, 55, 45, 20} {
		out.Reset()
		err = jpeg.Encode(out, img, &jpeg.Options{Quality: quality})
		if err != nil {
			return nil, err
		}

		if out.Len() < 0x1BF9C {
			return out.Bytes(), nil
		}
	}

	return out.Bytes(), fmt.Errorf("RepackThumbnailUntilFit unable to shrink enough %v > %v", out.Len(), 0x1C000)
}

// Add neccesary data at the end of the thumbnail
func EncryptJpegThumbnail(buf []byte) ([]byte, error) {
	// return if already an encrypted thumbnail
	if len(buf) == 0x1C000 && bytes.Equal(buf[0x1Bf9C:(0x1BF9C+4)], []byte{0x9C, 0xBF, 0x01, 0x00}) {
		return buf, nil
	}

	// 0x1BF9C is unecrypted thumbnail, 0x1C000 is already encrypted but will be reencrypted just in case
	if len(buf) > 0x1BF9C {
		var err error
		buf, err = RepackThumbnailUntilFit(buf)
		if err != nil {
			return []byte{}, err
		}
	}

	bufNew := make([]byte, 0x1C000)
	copy(bufNew, buf)

	// Technically random bytes, we make it deterministic here
	randomSeed := []byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}
	r := &Random{
		binary.LittleEndian.Uint32(randomSeed[0:4]),
		binary.LittleEndian.Uint32(randomSeed[4:8]),
		binary.LittleEndian.Uint32(randomSeed[8:12]),
		binary.LittleEndian.Uint32(randomSeed[12:16]),
	}

	sha256Key := new(bytes.Buffer)
	createKey(r, thumbnailTable, 0x10, sha256Key)

	mac := hmac.New(sha256.New, sha256Key.Bytes())
	mac.Write(bufNew[:0x1BF9C])

	// Add some required bytes
	copy(bufNew[0x1Bf9C:], []byte{0x9C, 0xBF, 0x01, 0x00})

	// Add mac
	copy(bufNew[0x1BFA0:], mac.Sum(nil))

	// Add random seed
	copy(bufNew[0x1BFC0:], randomSeed)

	// Add padding bytes
	copy(bufNew[0x1BFD0:], make([]byte, 0x30))

	return bufNew, nil
}
