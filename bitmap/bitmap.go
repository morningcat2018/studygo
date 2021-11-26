package bitmap

import (
	"errors"
)

type BitMap struct {
	words        []uint64
	wordsInUse   int
	sizeIsSticky bool
}

const addressBitsPerWord uint = 6 // 64
const bitSize int = 64

func New(value int) *BitMap {
	if value < 0 {
		panic("value can not < 0")
	}
	words := initWords(value)
	return &BitMap{words, 0, true}
}

func initWords(nbits int) []uint64 {
	words := make([]uint64, wordIndex(nbits-1)+1)
	return words
}

func (bitMap *BitMap) Set(value int) error {
	if value < 0 {
		return errors.New("value can not < 0")
	}
	wordIndex := wordIndex(value)
	bitMap.expandTo(wordIndex)
	bitMap.words[wordIndex] |= 1 << uint64(value%bitSize)
	bitMap.checkInvariants()
	return nil
}

func wordIndex(bitIndex int) int {
	return bitIndex >> addressBitsPerWord
}

func (bitMap *BitMap) expandTo(wordIndex int) {
	wordsRequired := wordIndex + 1
	if bitMap.wordsInUse < wordsRequired {
		bitMap.ensureCapacity(wordsRequired)
		bitMap.wordsInUse = wordsRequired
	}
}

func (bitMap *BitMap) ensureCapacity(wordsRequired int) {
	if len(bitMap.words) < wordsRequired {
		request := max(2*len(bitMap.words), wordsRequired)
		bitMap.words = copyOf(bitMap.words, request)
		bitMap.sizeIsSticky = false
	}
}

func copyOf(original []uint64, newLength int) []uint64 {
	if len(original) >= newLength {
		return original
	}
	copy1 := make([]uint64, newLength)
	for index, v := range original {
		copy1[index] = v
	}

	return copy1
}

func (bitMap *BitMap) checkInvariants() {
	//assert(wordsInUse == 0 || words[wordsInUse - 1] != 0);
	//assert(wordsInUse >= 0 && wordsInUse <= words.length);
	//assert(wordsInUse == words.length || words[wordsInUse] == 0);
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func (bitMap *BitMap) Get(value int) (bool, error) {
	if value < 0 {
		return false, errors.New("value can not < 0")
	}
	bitMap.checkInvariants()
	wordIndex := wordIndex(value)
	if wordIndex < bitMap.wordsInUse {
		x := bitMap.words[wordIndex]
		y := uint64(1 << uint(value%bitSize))
		return x&y != 0, nil
	}
	return false, nil
}
