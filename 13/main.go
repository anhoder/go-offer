package main

import "fmt"

//地上有一个m行n列的方格，从坐标 [0,0] 到坐标 [m-1,n-1] 。一个机器人从坐标 [0, 0] 的格子开始移动，它每次可以向左、右、上、下移动一格（不能移动到方格外），也不能进入行坐标和列坐标的数位之和大于k的格子。例如，当k为18时，机器人能够进入方格 [35, 37] ，因为3+5+3+7=18。但它不能进入方格 [35, 38]，因为3+5+3+8=19。请问该机器人能够到达多少个格子？
//
//
//示例 1：
//
//输入：m = 2, n = 3, k = 1
//输出：3
//示例 2：
//
//输入：m = 3, n = 1, k = 0
//输出：1
//提示：
//
//1 <= n,m <= 100
//0 <= k <= 20

func movingCount(m int, n int, k int) int {
    count := 0
    reachArea := make([][]bool, m)
    for i := 0; i < m; i++ {
        reachArea[i] = make([]bool, n)
    }

    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if i == 0 && j == 0 {
                reachArea[0][0] = true
                count++
                continue
            }

            if i/10+i%10+j/10+j%10 > k {
                continue
            }

            if (j-1 >= 0 && reachArea[i][j-1]) ||
                (i-1 >= 0 && reachArea[i-1][j]) ||
                (j+1 < len(reachArea[i]) && reachArea[i][j+1]) ||
                (i+1 < len(reachArea) && reachArea[i+1][j]) {
                reachArea[i][j] = true
                count++
                continue
            }

        }
    }

    return count
}

func main() {
    fmt.Println(movingCount(38, 15, 9))
}
