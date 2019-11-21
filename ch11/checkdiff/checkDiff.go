package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
	"unicode"
)

const chineseDigits = "一二三四五六七八九十"

func specify(str string) string {
	leftTrimed := strings.TrimLeftFunc(str, func(r rune) bool {
		return unicode.IsDigit(r) || r == '.' || strings.ContainsRune(chineseDigits, r)
	})
	chineseBracketsReplaced := strings.Replace(leftTrimed, "（", "(", 1)
	chineseBracketsReplaced = strings.Replace(chineseBracketsReplaced, "）", ")", 1)
	re := regexp.MustCompile(`^.*?\(?([a-zA-Z+-]+)\)?$`)
	return re.ReplaceAllString(chineseBracketsReplaced, "$1")
}

func parseLines(path string) (res []string) {
	f, err := os.Open(path)
	if err != nil {
		fmt.Fprintf(os.Stderr, "check diff : %v", err)
		return
	}

	input := bufio.NewScanner(f)
	for input.Scan() {
		res = append(res, input.Text())
	}

	return
}

func parseCSV(path string) (res []string) {
	for _, line := range parseLines(path) {
		parts := strings.Split(line, ",")
		res = append(res, specify(parts[0]+parts[1]))
	}

	return
}

func parseMd(path string) (res []string) {
	for _, line := range parseLines(path) {
		if parts := strings.Split(line, " "); len(parts) > 1 {
			res = append(res, specify(parts[1]))
		}
	}

	return
}

func setDiff(a []string, b []string) (res []string) {
	for _, aOne := range a {
		has := false
	inner:
		for _, bOne := range b {
			if aOne == bOne {
				has = true
				break inner
			}
		}
		if !has {
			res = append(res, aOne)
		}
	}

	return
}

func main() {
	md, cvs := "/Users/youxin/Downloads/建设方案.md", "/Users/youxin/Downloads/软件明细.csv"
	otherSet, ownSet := parseMd(md), parseCSV(cvs)
	fmt.Fprintln(os.Stdout, "<h1>建设方案中有，而软件明细中没有的如下：</h1>")
	for _, other := range setDiff(otherSet, ownSet) {
		fmt.Fprintf(os.Stdout, "<code style=\"color: blue;\">%s</code><br/>", other)
	}
	fmt.Fprintln(os.Stdout, "<h1>软件明细中有，而建设方案中没有的如下：</h1>")
	for _, own := range setDiff(ownSet, otherSet) {
		fmt.Fprintf(os.Stdout, "<code style=\"color: red;\">%s</code><br/>", own)
	}
	fmt.Fprintln(os.Stdout, "<h6>Powered By Golang</h6>")
}
