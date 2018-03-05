package main

import "fmt"

//Given an array of words and a length L, format the text such that each line has exactly L characters and is fully (left and right) justified.
//
//You should pack your words in a greedy approach; that is, pack as many words as you can in each line. Pad extra spaces ' ' when necessary so that each line has exactly L characters.
//
//Extra spaces between words should be distributed as evenly as possible. If the number of spaces on a line do not divide evenly between words, the empty slots on the left will be assigned more spaces than the slots on the right.
//
//For the last line of text, it should be left justified and no extra space is inserted between words.
//
//For example,
//words: ["This", "is", "an", "example", "of", "text", "justification."]
//L: 16.
//
//Return the formatted lines as:
//[
//   "This    is    an",
//   "example  of text",
//   "justification.  "
//]
//Note: Each word is guaranteed not to exceed L in length.
//

func fullJustify(words []string, maxWidth int) []string {
	if len(words) == 0 {
		return []string{}
	}
	res := make([]string, 0, 0)
	curCnt := 0
	curLen := 0
	temp := []string{}
	temp = append(temp, words[0])
	curCnt = 1
	curLen = len(words[0])

	genLine := func(tps []string, curLen int) string {
		cs := ""
		cnt := len(tps) - 1
		resLen := maxWidth - curLen
		perSpace := make([]int, cnt)
		var od, pp int
		if cnt > 0 {
			od = resLen % cnt
			pp = resLen / cnt
		} else {
			od = 0
			pp = 0
		}

		for i := 0; i < cnt; i++ {
			if i < od {
				perSpace[i] = 2 + pp
			} else {
				perSpace[i] = 1 + pp
			}
		}
		for i := 0; i < len(temp); i++ {
			cs += temp[i]

			if i >= cnt {
				continue
			}
			for j := 0; j < perSpace[i]; j++ {
				cs += " "
			}
		}
		curCsLen := len(cs)
		for i := 0; i < maxWidth-curCsLen; i++ {
			cs += " "
		}
		return cs
	}

	for k := 1; k < len(words); k++ {
		s := words[k]
		if curLen+1+len(s) <= maxWidth {
			curCnt++
			curLen += (1 + len(s))
			temp = append(temp, s)
		} else {

			cs := genLine(temp, curLen)

			res = append(res, cs)
			temp = make([]string, 0)
			temp = append(temp, s)
			curCnt = 1
			curLen = len(s)
		}
	}
	if len(temp) > 0 {
		//cs := genLine(temp, curLen)
		//res = append(res, cs)
		cs := ""
		for i, s := range temp {
			cs += s
			if i < len(temp)-1 {
				cs += " "
			}
		}
		cl := maxWidth - len(cs)
		for i := 0; i < cl; i++ {
			cs += " "
		}
		res = append(res, cs)
	}
	return res
}

func main() {
	s := []string{"This", "is", "an", "example", "of", "text", "justification."}
	//s := []string{"example", "of", "text"}
	res := fullJustify(s, 16)
	fmt.Println("----------------")
	for _, r := range res {
		fmt.Println(r)
	}
}
