package smm2_parsing

import (
	"bytes"
	"compress/zlib"
	"crypto/aes"
	"crypto/cipher"
	"encoding/binary"
	"fmt"
	"hash/crc32"
	"io"

	"github.com/aead/cmac"
)

var bcdTable = []uint32{
	0x7ab1c9d2, 0xca750936, 0x3003e59c, 0xf261014b,
	0x2e25160a, 0xed614811, 0xf1ac6240, 0xd59272cd,
	0xf38549bf, 0x6cf5b327, 0xda4db82a, 0x820c435a,
	0xc95609ba, 0x19be08b0, 0x738e2b81, 0xed3c349a,
	0x045275d1, 0xe0a73635, 0x1debf4da, 0x9924b0de,
	0x6a1fc367, 0x71970467, 0xfc55abeb, 0x368d7489,
	0x0cc97d1d, 0x17cc441e, 0x3528d152, 0xd0129b53,
	0xe12a69e9, 0x13d1bdb7, 0x32eaa9ed, 0x42f41d1b,
	0xaea5f51f, 0x42c5d23c, 0x7cc742ed, 0x723ba5f9,
	0xde5b99e3, 0x2c0055a4, 0xc38807b4, 0x4c099b61,
	0xc4e4568e, 0x8c29c901, 0xe13b34ac, 0xe7c3f212,
	0xb67ef941, 0x08038965, 0x8afd1e6a, 0x8e5341a3,
	0xa4c61107, 0xfbaf1418, 0x9b05ef64, 0x3c91734e,
	0x82ec6646, 0xfb19f33e, 0x3bde6fe2, 0x17a84cca,
	0xccdf0ce9, 0x50e4135c, 0xff2658b2, 0x3780f156,
	0x7d8f5d68, 0x517cbed1, 0x1fcddf0d, 0x77a58c94,
}

func Decompress(buf []byte) ([]byte, error) {
	reader, err := zlib.NewReader(bytes.NewReader(buf))
	if err != nil {
		return nil, err
	}
	defer reader.Close()

	decompressedBuf, err := io.ReadAll(reader)
	if err != nil {
		return nil, err
	}

	return decompressedBuf, nil
}

func Compress(data []byte) ([]byte, error) {
	buf := &bytes.Buffer{}
	w := zlib.NewWriter(buf)

	_, err := w.Write(data)
	if err != nil {
		return []byte{}, err
	}

	err = w.Close()
	if err != nil {
		return []byte{}, err
	}

	return buf.Bytes(), nil
}

func DecryptLevel(buf []byte) ([]byte, error) {
	if len(buf) != 0x5c000 {
		return []byte{}, fmt.Errorf("invalid buf size %d != %d", len(buf), 0x5c000)
	}

	end := 0x5BFD0
	writer := new(bytes.Buffer)

	// Create random instance
	r := &Random{
		binary.LittleEndian.Uint32(buf[end+0x10 : end+0x14]),
		binary.LittleEndian.Uint32(buf[end+0x14 : end+0x18]),
		binary.LittleEndian.Uint32(buf[end+0x18 : end+0x1C]),
		binary.LittleEndian.Uint32(buf[end+0x1C : end+0x20]),
	}

	cmacWant := buf[end+0x20 : end+0x30]
	crcWant := buf[8:12]

	// Construct AES instance
	aesKey := new(bytes.Buffer)
	createKey(r, bcdTable, 0x10, aesKey)

	aesBlock, err := aes.NewCipher(aesKey.Bytes())
	if err != nil {
		return nil, err
	}

	aesMode := cipher.NewCBCDecrypter(aesBlock, buf[end:end+0x10])
	decrypted := make([]byte, 0x5BFC0)
	aesMode.CryptBlocks(decrypted, buf[0x10:0x5BFD0])

	// crc check
	if crc32.ChecksumIEEE(decrypted) != binary.LittleEndian.Uint32(crcWant) {
		return nil, fmt.Errorf("crc invalid")
	}

	// cmac check
	cmacKey := new(bytes.Buffer)
	createKey(r, bcdTable, 0x10, cmacKey)
	cmacBlock, err := aes.NewCipher(cmacKey.Bytes())
	if err != nil {
		return nil, err
	}
	cmacDigest, err := cmac.Sum(decrypted, cmacBlock, 0x10)
	if err != nil {
		return nil, err
	}
	if !bytes.Equal(cmacDigest, cmacWant) {
		return nil, fmt.Errorf("cmac invalid")
	}

	if false {
		// write bcd header
		_, err = writer.Write(buf[:0x10])
		if err != nil {
			return nil, err
		}
	}

	// Decrypted course data
	_, err = writer.Write(decrypted)
	if err != nil {
		return nil, err
	}

	return writer.Bytes(), nil
}

func EncryptLevel(buf []byte) ([]byte, error) {
	var withoutBcdHeader bool

	if len(buf) == 0x5BFD0-0x10 {
		withoutBcdHeader = true
	} else if len(buf) == 0x5BFD0 {
		withoutBcdHeader = false
	} else {
		return []byte{}, fmt.Errorf("invalid buf size %d != %d (%d)", len(buf), 0x5BFD0, 0x5BFC0-0x10)
	}

	var reader *bytes.Reader
	if withoutBcdHeader {
		reader = bytes.NewReader(buf)
	} else {
		reader = bytes.NewReader(buf[0x10:])
	}

	decrypted := make([]byte, 0x5BFC0)
	_, err := io.ReadFull(reader, decrypted)
	if err != nil {
		return nil, err
	}

	writer := new(bytes.Buffer)

	if withoutBcdHeader {
		var data = []any{
			uint32(0x1),
			uint16(0x10),
			uint16(0x0),
			uint32(crc32.ChecksumIEEE(decrypted)),
			uint8(0x53),
			uint8(0x43),
			uint8(0x44),
			uint8(0x4C),
		}
		for _, v := range data {
			err := binary.Write(writer, binary.LittleEndian, v)
			if err != nil {
				return nil, err
			}
		}
	} else {
		binary.LittleEndian.PutUint32(buf[0x8:(0x8+4)], crc32.ChecksumIEEE(decrypted)) // shouldn't be needed, only when modifing course data
		writer.Write(buf[:0x10])
	}

	// Technically random bytes, we make it deterministic here
	randomSeed := []byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}
	r := &Random{
		binary.LittleEndian.Uint32(randomSeed[0:4]),
		binary.LittleEndian.Uint32(randomSeed[4:8]),
		binary.LittleEndian.Uint32(randomSeed[8:12]),
		binary.LittleEndian.Uint32(randomSeed[12:16]),
	}
	// Also random, but deterministic here
	aesIv := []byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}

	aesKey := new(bytes.Buffer)
	createKey(r, bcdTable, 0x10, aesKey)
	aesBlock, err := aes.NewCipher(aesKey.Bytes())
	if err != nil {
		return nil, err
	}

	aesMode := cipher.NewCBCEncrypter(aesBlock, aesIv)
	encrypted := make([]byte, 0x5BFC0)
	aesMode.CryptBlocks(encrypted, decrypted)

	cmacKey := new(bytes.Buffer)
	createKey(r, bcdTable, 0x10, cmacKey)
	cmacBlock, err := aes.NewCipher(cmacKey.Bytes())
	if err != nil {
		return nil, err
	}

	cmacDigest, err := cmac.Sum(decrypted, cmacBlock, 0x10)
	if err != nil {
		return nil, err
	}

	_, err = writer.Write(encrypted)
	if err != nil {
		return nil, err
	}

	_, err = writer.Write(aesIv)
	if err != nil {
		return nil, err
	}

	_, err = writer.Write(randomSeed)
	if err != nil {
		return nil, err
	}

	_, err = writer.Write(cmacDigest)
	if err != nil {
		return nil, err
	}

	return writer.Bytes(), nil
}

func DecryptReplay(buf []byte) ([]byte, error) {
	return nil, nil
}
