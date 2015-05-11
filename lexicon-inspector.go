// language-inspector inspects the data for the given language.
package main

import (
	"flag"
	"fmt"
	"path"

	"github.com/BluntSporks/lexicon"
	"github.com/BluntSporks/mapsort"
)

func main() {
	// Parse flags.
	langFile := flag.String("lang", "english", "Name of language file to inspect")
	lexDir := flag.String("lexdir", lex.DefaultDataPath(), "Location of lexicon data directory")
	mode := flag.String("mode", "cntchars", "Mode of program execution")
	strlen := flag.Int("strlen", 1, "Number of runes to include in each substring")
	flag.Parse()

	// Load the language and character counts.
	path := path.Join(*lexDir, *langFile)
	wordList := lex.LoadLang(path)

	// Determine mode of operation.
	if *mode == "cntstrs" {
		// Print the substring counts.
		strCnts := lex.CntSubstrs(wordList, *strlen)
		pairs := mapsort.ByVal(strCnts, false)
		for _, pair := range pairs {
			fmt.Printf("%s,%d\n", pair.Key, pair.Val)
		}
	} else {
		// Print the character counts.
		charCnts := lex.CntChars(wordList)
		for char, cnt := range charCnts {
			fmt.Printf("%#U,%d\n", char, cnt)
		}
	}
}
