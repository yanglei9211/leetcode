package main

import "fmt"

//Given a 2D board and a word, find if the word exists in the grid.
//
//The word can be constructed from letters of sequentially adjacent cell, where "adjacent" cells are those horizontally or vertically neighboring. The same letter cell may not be used more than once.
//
//For example,
//Given board =
//
//[
//  ['A','B','C','E'],
//  ['S','F','C','S'],
//  ['A','D','E','E']
//]
//word = "ABCCED", -> returns true,
//word = "SEE", -> returns true,
//word = "ABCB", -> returns false.

type ss79 struct {
	board  [][]byte
	word   string
	bs     []byte
	fx, fy []int
	n, m   int
	vis    [][]bool
}

func (s *ss79) init(board [][]byte, word string) {
	s.board = board
	s.word = word
	s.bs = []byte(word)
	s.fx = []int{0, -1, 0, 1}
	s.fy = []int{-1, 0, 1, 0}
	s.n = len(board)
	if s.n > 0 {
		s.m = len(board[0])
	} else {
		s.m = 0
	}
	s.vis = make([][]bool, s.n)
	for i := 0; i < s.n; i++ {
		s.vis[i] = make([]bool, s.m)
	}
}

func (s *ss79) check(x, y int, res []byte) bool {
	//if len(res) == 0 {
	//	return true
	//}

	if s.board[x][y] == res[0] {
		s.vis[x][y] = true
		if len(res) == 1 {
			return true
		}
		for dir := 0; dir < len(s.fx); dir++ {
			nx := x + s.fx[dir]
			ny := y + s.fy[dir]
			if nx >= 0 && nx < s.n && ny >= 0 && ny < s.m && !s.vis[nx][ny] && s.board[nx][ny] == res[1] {
				if s.check(nx, ny, res[1:]) {
					return true
				}
			}
		}
		s.vis[x][y] = false
		return false
	} else {
		return false
	}
}

func (s *ss79) do() bool {
	for x := 0; x < s.n; x++ {
		for y := 0; y < s.m; y++ {
			if s.board[x][y] == s.bs[0] && s.check(x, y, s.bs) {
				return true
			}
		}
	}
	return false
}

func exist(board [][]byte, word string) bool {
	w := ss79{}
	w.init(board, word)
	if w.n == 0 || w.m == 0 {
		return false
	}
	return w.do()
}

func main() {
	bb := [][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}}
	//words := []string{"ABCCED", "SEE", "ABCB"}
	fmt.Println(exist(bb, "SEE"))
}
