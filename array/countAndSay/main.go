package main

import "fmt"
import "strings"
import "strconv"


/*给定一个正整数 n ，输出外观数列的第 n 项。

「外观数列」是一个整数序列，从数字 1 开始，序列中的每一项都是对前一项的描述。

你可以将其视作是由递归公式定义的数字字符串序列：

    countAndSay(1) = "1"
    countAndSay(n) 是对 countAndSay(n-1) 的描述，然后转换成另一个数字字符串。

前五项如下：

1.     1
2.     11
3.     21
4.     1211
5.     111221
第一项是数字 1 
描述前一项，这个数是 1 即 “ 一 个 1 ”，记作 "11"
描述前一项，这个数是 11 即 “ 二 个 1 ” ，记作 "21"
描述前一项，这个数是 21 即 “ 一 个 2 + 一 个 1 ” ，记作 "1211"
描述前一项，这个数是 1211 即 “ 一 个 1 + 一 个 2 + 二 个 1 ” ，记作 "111221"

要 描述 一个数字字符串，首先要将字符串分割为 最小 数量的组，每个组都由连续的最多 相同字符 组成。
然后对于每个组，先描述字符的数量，然后描述字符，形成一个描述组。要将描述转换为数字字符串，
先将每组中的字符数量用数字替换，再将所有描述组连接起来。*/

// func countAndSay(n int) string{
// 	if n == 1 {
// 		return "1"
// 	}else if n == 2 {
// 		return "11"
// 	}
// 	str := "11"
// 	m := make(map[byte]int)
// 	for i := 0; i < n-2; i++{
// 		ptr := ""
// 		for j := 0; j < len(str)-1; j++ {
// 			if str[j] == str[j+1]{
// 				m[str[j]]++
// 			}else {
// 				ptr = ptr + strconv.Itoa(m[str[j]]+1) + string(str[j])
// 				delete(m, str[j])
// 			}
// 		}
// 		//最后一项没有扫到，且字符串无法修改,用rune转换成切片
// 		strSli := []rune(ptr)
// 		//如果最后一项和倒数第二相一样，map+1；不一样就加一项
// 		if str[len(str)-1] == str[len(str)-2] {
// 			strSli[len(strSli)-2] = strSli[len(strSli)-2] + 1
// 		}else{
// 			strSli = append(strSli, '1', rune(str[len(str)-1]))
// 		}
// 		str = string(strSli)
// 	}
	
// 	return str
// }

func main(){
	fmt.Println(countAndSay(3))
}
func countAndSay(n int) string {
    prev := "1"
    for i := 2; i <= n; i++ {
        cur := &strings.Builder{}
        for j, start := 0, 0; j < len(prev); start = j {
            for j < len(prev) && prev[j] == prev[start] {
                j++
            }
            cur.WriteString(strconv.Itoa(j - start))
            cur.WriteByte(prev[start])
        }
        prev = cur.String()
    }
    return prev
}