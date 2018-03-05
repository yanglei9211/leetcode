package main

import "fmt"

type DLX struct {
	L, R, U, D     []int
	S, col, row    []int
	size, mRow     int
	locX, locY     []int
	head           [][]int
	N              int
	rowNum, colNum int
	ans            []int
	totSol         int
	output         [][]int
}

func (d *DLX) resume(c int) {
	d.L[d.R[c]] = c
	d.R[d.L[c]] = c
	for i := d.U[c]; i != c; i = d.U[i] {
		for j := d.L[i]; j != i; j = d.L[j] {
			d.U[d.D[j]] = j
			d.D[d.U[j]] = j
			d.S[d.col[j]]++
		}
	}
}

func (d *DLX) remove(c int) {
	d.R[d.L[c]] = d.R[c]
	d.L[d.R[c]] = d.L[c]
	for i := d.D[c]; i != c; i = d.D[i] {
		for j := d.R[i]; j != i; j = d.R[j] {
			d.U[d.D[j]] = d.U[j]
			d.D[d.U[j]] = d.D[j]
			d.S[d.col[j]]--
		}
	}
}

func (d *DLX) init(row, n, m int) {
	d.totSol = 0
	d.ans = make([]int, row+1)
	d.rowNum = n
	d.colNum = m
	d.N = row
	maxn := (n + 1) * (m + 1)
	d.head = make([][]int, n+1)
	for i := 0; i <= n; i++ {
		d.head[i] = make([]int, m+1)
	}
	d.S = make([]int, m+1)
	d.L = make([]int, maxn)
	d.R = make([]int, maxn)
	d.U = make([]int, maxn)
	d.D = make([]int, maxn)
	d.locX = make([]int, maxn)
	d.locY = make([]int, maxn)
	d.col = make([]int, maxn)
	d.row = make([]int, maxn)

	for i := 1; i <= m; i++ {
		d.L[i] = i - 1
		d.R[i] = i + 1
		d.U[i] = i
		d.D[i] = i
		d.col[i] = i
		d.row[i] = 0
	}
	d.size = m
	d.mRow = 0
	d.L[0] = m
	d.R[0] = 1
	d.R[m] = 0
	d.U[0] = 0
	d.D[0] = 0
}

func (d *DLX) makehead(c int) int {
	d.size++
	d.S[c]++
	d.col[d.size] = c
	d.row[d.size] = d.mRow
	d.L[d.size] = d.size
	d.R[d.size] = d.size
	d.U[d.size] = c
	d.D[d.size] = d.D[c]
	d.U[d.D[d.size]] = d.size
	d.D[d.U[d.size]] = d.size
	return d.size
}

func (d *DLX) addcol(id, c int) {
	d.size++
	size := d.size
	d.S[c]++
	d.col[size] = c
	d.row[size] = d.mRow

	d.L[size] = id
	d.R[size] = d.R[id]
	d.L[d.R[size]] = size
	d.R[d.L[size]] = size

	d.U[size] = c
	d.D[size] = d.D[c]
	d.U[d.D[size]] = size
	d.D[d.U[size]] = size
}

func (d *DLX) addrow(i, j int) {
	var id int
	d.mRow++
	d.locX[d.mRow] = i
	d.locY[d.mRow] = j
	id = d.makehead(i + 1)
	d.head[i][j] = id
	d.addcol(id, j+d.N+1)
	d.addcol(id, i+j+(d.N<<1)+1)
	d.addcol(id, i-j+d.N-1+(d.N<<2))
}

func (d *DLX) show() {
	mp := make([][]int, d.rowNum+1)
	for i := 0; i <= d.rowNum; i++ {
		mp[i] = make([]int, d.colNum+1)
	}
	for i := d.R[0]; i != 0; i = d.R[i] {
		mp[d.row[i]][d.col[i]] = i
		for j := d.D[i]; j != i; j = d.D[j] {
			mp[d.row[j]][d.col[j]] = j
		}
	}
	for i := 0; i <= d.rowNum; i++ {
		if i == 1 {
			fmt.Println("-----------------------------------------------------")
		}
		for j := 0; j <= d.colNum; j++ {
			fmt.Printf("%d\t", mp[i][j])
		}
		fmt.Println()
	}
	fmt.Println()
	fmt.Println()
	fmt.Println()
}

func (d *DLX) saveAnsNum() {
	d.totSol++
}

func (d *DLX) saveAns() {
	res := make([]int, d.N+1)
	for i := 0; i <= d.N; i++ {
		res[i] = d.locY[d.ans[i]]
	}
	d.output = append(d.output, res)
}

func (d *DLX) dfs(k int) {
	if k == d.N {
		//for _, i := range(d.ans) {
		//	fmt.Printf("%d, ", d.locY[i])
		//}
		//fmt.Println()
		//d.saveAnsNum()
		d.saveAns()
		return
	}
	c := 0
	minn := 1 << 29
	for i := d.R[0]; i <= d.N; i = d.R[i] {
		if d.S[i] == 0 {
			return
		}
		if d.S[i] < minn {
			minn = d.S[i]
			c = i
		}
	}
	d.remove(c)
	for i := d.D[c]; i != c; i = d.D[i] {
		d.ans[c] = d.row[i]
		for j := d.R[i]; j != i; j = d.R[j] {
			d.remove(d.col[j])
		}
		d.dfs(k + 1)
		for j := d.L[i]; j != i; j = d.L[j] {
			d.resume(d.col[j])
		}
	}
	d.resume(c)
}

func totalNQueens(n int) int {
	dlx := DLX{}
	dlx.init(n, n*n, 6*n-2)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			dlx.addrow(i, j)
		}
	}
	dlx.dfs(0)
	return dlx.totSol
}

func solveNQueens(n int) [][]string {
	dlx := DLX{}
	dlx.init(n, n*n, 6*n-2)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			dlx.addrow(i, j)
		}
	}
	dlx.dfs(0)
	fmt.Println(dlx.output)
	res := [][]string{}
	for _, arr := range dlx.output {
		one := []string{}
		for idx, r := range arr {
			if idx == 0 {
				continue
			}
			bs := ""
			for i := 1; i <= dlx.N; i++ {
				if i == r+1 {
					bs += "Q"
				} else {
					bs += "."
				}
			}
			one = append(one, bs)
		}
		res = append(res, one)
	}
	return res
}

func main() {
	ss := solveNQueens(6)
	fmt.Println(len(ss))
	for _, s := range ss {
		for _, r := range s {
			fmt.Println(r)
		}
		fmt.Println()
		fmt.Println()
	}
	//n := 4
	//nn := n*n
	//m := 6*n-2
	//dlx := DLX{}
	//dlx.init(n, nn ,m)
	//for i := 0; i < n; i ++ {
	//	for j :=0 ; j < n; j ++ {
	//		dlx.addrow(i,j)
	//	}
	//}
	////dlx.show()
	////dlx.remove(1)
	////dlx.show()
	//dlx.dfs(0)
	////fmt.Println(dlx.totSol)
	//fmt.Println(dlx.output)
}
