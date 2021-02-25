package tools

import (
	"math/rand"
	"time"
)

var (
	thai     = []int64{3585, 3654}
	armenian = []int64{1328, 1423}
	chinese  = []int64{19968, 40869}
	//JapaneseKatakana .
	JapaneseKatakana = []int64{12449, 12531}
	//JapaneseHiragana .
	JapaneseHiragana = []int64{12353, 12435}
	koreanHangul     = []int64{12593, 12686}
	cyrillianRussian = []int64{1025, 1169}
	greek            = []int64{884, 974}
)

func randInt(start, end int64) int64 {
	rand.Seed(time.Now().UnixNano())
	return (start + rand.Int63n(end-start))
}

//GenerateRandomRune 随机名
func GenerateRandomRune(size int, start, end int64) string {
	randRune := make([]rune, size)
	for i := range randRune {
		randRune[i] = rune(randInt(start, end))
	}
	return string(randRune)
}
