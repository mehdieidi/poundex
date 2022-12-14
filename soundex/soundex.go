package soundex

import (
	"fmt"
	"strconv"
	"strings"
)

var SoundexMap = map[string]int{
	"ب": 1, "پ": 1, "ف": 1,
	"س": 2, "ص": 2, "ث": 2, "ج": 2, "ز": 2, "ض": 2, "ظ": 2, "ذ": 2, "ژ": 2, "چ": 2, "ک": 2, "ك": 2, "غ": 2, "ق": 2, "گ": 2, "خ": 2, "ش": 2,
	"د": 3, "ت": 3, "ط": 3,
	"ل": 4,
	"م": 5, "ن": 5, "ر": 6,
}

var FirstLetterMap = map[string]string{
	"ب": "B", "پ": "P", "ت": "T", "ط": "T", "س": "S", "ص": "S", "ث": "S", "ج": "J", "چ": "C", "ر": "R",
	"ة": "H", "ه": "H", "ح": "H", "خ": "X", "د": "D", "ز": "Z", "ض": "Z", "ظ": "Z", "ذ": "Z", "ژ": "Z", "ش": "S",
	"غ": "G", "ک": "K", "ك": "K", "گ": "G", "ق": "G", "ف": "F", "ل": "L", "م": "M", "ن": "N", "و": "V", "ع": "A",
	"ا": "A", "آ": "A", "أ": "A", "إ": "A", "ی": "Y", "ئ": "Y", "ي": "Y", "ء": "A",
}

func Get(persianWord string) (soundex string, err error) {
	if len(persianWord) <= 2 {
		return "", ErrPersianWordTooSmall
	}

	persianChars := strings.Split(persianWord, "")

	f, ok := FirstLetterMap[persianChars[0]]
	if !ok {
		fmt.Println("err:", persianWord, "char:", persianChars[0])
		return "", ErrInvalidPersianWord
	}

	soundex = f

	for i := 1; i < len(persianChars) && i < 4; i++ {
		s, ok := SoundexMap[persianChars[i]]
		if !ok {
			continue
		}

		soundex += strconv.Itoa(s)
	}

	for len(soundex) < 4 {
		soundex += "0"
	}

	return
}
