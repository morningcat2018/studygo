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
		"Matthew", "Sarah", "Augustus", "Heidi", "Emilie", "Peter", "Giana", "Adriano", "Aaron", "Elizabeth", "love",
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
		distribution[name] = coin
		if coin > coins {
			fmt.Println(distribution)
			panic("硬币不足")
		}
		coins = coins - coin
	}
	fmt.Println(distribution)
	return coins
}

func calCoin(name string) int {
	var coin = 0
	for _, sc := range name {
		coin += calCoinDetail2(sc)
	}
	return coin
}

func calCoinDetail(sc rune) int {
	if sc == 'e' || sc == 'E' {
		return 1
	} else if sc == 'i' || sc == 'I' {
		return 2
	} else if sc == 'o' || sc == 'O' {
		return 3
	} else if sc == 'u' || sc == 'U' {
		return 4
	}
	return 0
}

func calCoinDetail2(sc rune) int {
	var coin = 0
	switch sc {
	case 'e', 'E':
		coin = 1
	case 'i', 'I':
		coin = 2
	case 'o', 'O':
		coin = 3
	case 'u', 'U':
		coin = 4
	}
	return coin
}
