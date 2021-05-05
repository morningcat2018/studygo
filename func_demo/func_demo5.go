package func_demo

import "fmt"

/*
你有50枚金币，需要分配给以下几个人：Matthew,Sarah,Augustus,Heidi,Emilie,Peter,Giana,Adriano,Aaron,Elizabeth。
分配规则如下：
a. 名字中每包含1个'e'或'E'分1枚金币
b. 名字中每包含1个'i'或'I'分2枚金币
c. 名字中每包含1个'o'或'O'分3枚金币
d: 名字中每包含1个'u'或'U'分4枚金币
写一个程序，计算每个用户分到多少金币，以及最后剩余多少金币？
程序结构如下，请实现 ‘dispatchCoin’ 函数
*/
var (
	coins = 50
	users = []string{
		"Matthew", "Sarah", "Augustus", "Heidi", "Emilie", "Peter", "Giana", "Adriano", "Aaron", "Elizabeth", "uUu",
	}
	distribution = make(map[string]int, len(users))
)

func Demo5Main() {
	left := dispatchCoin()
	fmt.Println("剩下：", left)
}

func dispatchCoin() int {
	for _, name := range users {
		coin := calCoin(name)
		if coin > coins {
			panic("硬币不足")
		}
		distribution[name] = coin
		coins = coins - coin
	}
	fmt.Println(distribution)
	return coins
}

func calCoin(name string) int {
	var coin = 0
	for _, sc := range name {
		if sc == 'e' || sc == 'E' {
			coin += 1
		} else if sc == 'i' || sc == 'I' {
			coin += 2
		} else if sc == 'o' || sc == 'O' {
			coin += 3
		} else if sc == 'u' || sc == 'U' {
			coin += 4
		}
	}
	return coin
}
