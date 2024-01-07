package spellingbeecheat

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func SpellingBeeCheat(centerLetter string, allLetters string) {
	center := strings.ToLower(centerLetter)
	all := strings.Split(strings.ToLower(allLetters), "")

	r := strings.NewReplacer(all[0], "", all[1], "", all[2], "", all[3], "", all[4], "", all[5], "", all[6], "")
	alphabet := r.Replace("abcdefghijklmnopqrstuvwxyz")

	file, err := os.Open("words_alpha.txt")
	if err != nil {
		log.Fatal(err)
	}
	scan := bufio.NewScanner(file)
	scan.Split(bufio.ScanLines)

	for scan.Scan() {
		words := strings.ToLower(scan.Text())

		if len(words) > 3 && !strings.Contains(words, "'") && strings.Contains(words, center) && !strings.ContainsAny(words, alphabet) {
			fmt.Println(words)
		}

	}

	file.Close()
}
