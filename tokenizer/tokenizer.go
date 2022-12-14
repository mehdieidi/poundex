package tokenizer

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var PersianChars = map[string]struct{}{
	"آ": {}, "ا": {}, "ب": {}, "پ": {}, "ت": {}, "ث": {}, "ج": {},
	"چ": {}, "ح": {}, "خ": {}, "د": {}, "ذ": {}, "ر": {}, "ز": {},
	"ژ": {}, "س": {}, "ش": {}, "ص": {}, "ض": {}, "ط": {}, "ظ": {},
	"ع": {}, "غ": {}, "ف": {}, "ق": {}, "ک": {}, "گ": {}, "ل": {},
	"م": {}, "ن": {}, "و": {}, "ه": {}, "ی": {}, "ي": {}, "ئ": {},
	"ة": {}, "ك": {}, "ء": {},
}

func ReadLines(file *os.File) (lines []string, err error) {
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return
}

func Tokenize(file *os.File, line string) ([]string, error) {
	tokens := []string{}

	line = strings.ReplaceAll(line, ".", " ")
	line = strings.ReplaceAll(line, ",", " ")
	line = strings.ReplaceAll(line, "-", " ")
	line = strings.ReplaceAll(line, "?", " ")
	line = strings.ReplaceAll(line, "؟", " ")
	line = strings.ReplaceAll(line, "!", " ")
	line = strings.ReplaceAll(line, "!", " ")
	line = strings.ReplaceAll(line, "،", " ")
	line = strings.ReplaceAll(line, "\r", " ")
	line = strings.ReplaceAll(line, "\n", " ")
	line = strings.ReplaceAll(line, ")", " ")
	line = strings.ReplaceAll(line, "(", " ")
	line = strings.ReplaceAll(line, "#", " ")
	line = strings.ReplaceAll(line, "$", " ")
	line = strings.ReplaceAll(line, "]", " ")
	line = strings.ReplaceAll(line, "[", " ")
	line = strings.ReplaceAll(line, "/", " ")

	segments := strings.Split(line, "")

	var strBuilder strings.Builder

	for _, s := range segments {
		if s == "\r" || s == "\n" {
			continue
		}

		if s == " " && strBuilder.Len() != 0 {
			token := strBuilder.String()

			tokens = append(tokens, token)
			file.WriteString(fmt.Sprintf("%s\n", token))

			strBuilder.Reset()
			continue
		}

		_, ok := PersianChars[s]
		if ok {
			strBuilder.WriteString(s)
		}
	}

	if strBuilder.Len() != 0 {
		file.WriteString(fmt.Sprintf("%s\n", strBuilder.String()))
	}

	return tokens, nil
}
