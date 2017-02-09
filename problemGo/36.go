package main

import "fmt"

func main() {
    x := [][]byte{
        {'1','2','.','4','5','6','7','8','9'},
        {'1','2','3','4','5','6','7','8','9'},
        {'1','2','3','4','5','6','7','8','9'},
        {'1','2','3','4','5','6','7','8','9'},
        {'1','2','3','4','5','6','7','8','9'},
        {'1','2','3','4','5','6','7','8','9'},
        {'1','2','3','4','5','6','7','8','9'},
        {'1','2','3','4','5','6','7','8','9'},
        {'1','2','3','4','5','6','7','8','9'},
    }
    //fmt.Println(genList(0,0,0,1, x))
    //fmt.Println(genGrid(0,0,x))
    fmt.Println(isValidSudoku(x))
}

func isValidSudoku(board [][]byte) bool {
    for i:= 0; i < 3; i++{
        for j := 0; j < 3; j++ {
            line := i * 3 + j
            col := checkNums(genList(line, 0, 0, 1, board))
            row := checkNums(genList(0, line, 1, 0, board))
            grid := checkNums(genGrid(i*3, j*3, board))
            if !(col && row && grid){
                return false
            }
        }
    }
    return true
}

func genList(x, y, fx, fy int, board [][]byte) []byte {
    res := []byte{}
    for i := 0; i < 9; i++ {
        if board[x][y] != '.'{
            res = append(res, board[x][y])
        }
        x += fx
        y += fy
    }
    return res
}

func genGrid(x, y int, board [][]byte) []byte {
    res := []byte{}
    for i := 0; i < 3; i++ {
        for j := 0; j < 3; j++{
            res = append(res, board[x+i][y+j])
        }
    }
    return res
}

func checkNums(nums []byte) bool {
    c := 0
    for _, x := range nums{
        if x != '.'{
            y := (1<<uint(x))
            if (c & y) > 0 {
                return false
            }
            c |= y
        }
    }
    return true
}
