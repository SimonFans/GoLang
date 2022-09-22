package Basic

import (
	"bufio"
	"os"
	"regexp"
)

// define regular expression to match the string we want to count freq
/* Example history file
: 1542784278:0;git push
: 1542784308:0;ls
: 1542784310:0;go test
: 1542784314:0;go test -v
: 1542784386:0;which gometalinter
: 1542784314:0;go test -v
*/
var cmdRe = regexp.MustCompile(`;go ([a-z]+)`)

// cmdFreq returns the frequency of "go" subcommand usage in ZSH history
func CmdFreq(filename string) (map[string]int, error) {
	file, err := os.Open("data/" + filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	freqs := make(map[string]int)
	// define a buffer io
	s := bufio.NewScanner(file)
	for s.Scan() {
		// scan each row text
		matches := cmdRe.FindStringSubmatch(s.Text())
		// if len(matches) == 0 {
		// 	continue
		// }
		if len(matches) != 0 {
			cmd := matches[1]
			freqs[cmd]++
		}
	}
	if err := s.Err(); err != nil {
		return nil, err
	}
	return freqs, nil
}
