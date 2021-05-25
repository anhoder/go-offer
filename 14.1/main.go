package main

import "fmt"

//给你一根长度为 n 的绳子，请把绳子剪成整数长度的 m 段（m、n都是整数，n>1并且m>1），每段绳子的长度记为 k[0],k[1]...k[m-1] 。请问 k[0]*k[1]*...*k[m-1] 可能的最大乘积是多少？例如，当绳子的长度是8时，我们把它剪成长度分别为2、3、3的三段，此时得到的最大乘积是18。
//
//示例 1：
//
//输入: 2
//输出: 1
//解释: 2 = 1 + 1, 1 × 1 = 1
//示例 2:
//
//输入: 10
//输出: 36
//解释: 10 = 3 + 3 + 4, 3 × 3 × 4 = 36
//提示：
//
//2 <= n <= 58

func main() {
    fmt.Println(cuttingRope(10))
}

func cuttingRope(n int) int {
    if n == 2 {
        return 1
    }
    if n == 3 {
        return 2
    }

    cache := []int{
        0,
        1,
        2,
        3,
    }

    for i := 4; i <= n; i++ {
        for j := 1; j <= i/2; j++ {
            tmp := cache[j] * cache[i-j]
            if i > len(cache)-1 {
                cache = append(cache, tmp)
            } else if tmp > cache[i] {
                cache[i] = tmp
            }
        }
    }

    return cache[n]
}
