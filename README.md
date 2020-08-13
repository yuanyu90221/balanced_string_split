# balance_string_split

## 題目解讀：

### 題目來源:
[split-a-string-in-balanced-strings](https://leetcode.com/problems/split-a-string-in-balanced-strings/)

### 原文:
Balanced strings are those who have equal quantity of 'L' and 'R' characters.

Given a balanced string s split it in the maximum amount of balanced strings.

Return the maximum amount of splitted balanced strings.

### 解讀：

定義 Balanced 字串：一個字串含有等量的'L'跟'R'

給定一個 Balanced 字串 s 

把 s 分割成子字串每個子字串都是 Balanced字串

找出 給定 s 能分割出 最大量的 Balanced子字串


## 初步解法:
### 初步觀察:

由於Balanced 字串具有 L字元個數跟R字元個數相同

因此 如果用一個 正整數 counter 預設為0

從左至右 一旦遇到 R就把counter 加一 遇到L就把counter 減一

則當 counter為零時 代表遇到一個Balanced 字串

舉例:

s: "RLRRLLRLRL"

 -> R count: 1
 -> RL count: 0  => RL為balance 字串
 -> RLR count: 1
 -> RLRR count: 2
 -> RLRRL count: 1
 -> RLRRLL count: 0 =>在RL之後的RRLL 為balance
 -> RLRRLLR count: 1
 -> RLRRLLRL count: 0 => RL
 -> RLRRLLRLR count:１
 -> RLRRLLRLRL count: 0 => RL
 
 s => RL, RRLL, RL, RL => 4
### 初步設計:
given a balanced string s

step 0: let a integer count = 0, amount = 0, sidx = 0

step 1: if sidx > length of s go to step 6

step 2: if s[sidx] == 'R' set count += 1

step 3: if s[sidx] == 'L' set count -= 1

step 4: if count == 0 set amount += 1 

step 5: sidx = sidx + 1 go to step 1

step 6: return amount

## 遇到的困難
### 題目上理解的問題
因為英文不是筆者母語

所以在題意解讀上 容易被英文用詞解讀給搞模糊

### pseudo code撰寫

一開始不習慣把pseudo code寫下來

因此 不太容易把自己的code做解析

### golang table driven test不熟
對於table driven test還不太熟析

所以對於寫test還是耗費不少時間
## 我的github source code
```golang
package balance_string

func balanceStringSplit(s string) int {
	count := 0
	amount := 0
	for _, r := range s {
		if r == 'R' {
			count += 1
		} else {
			count -= 1
		}
		if count == 0 {
			amount += 1
		}
	}
	return amount
}

```
## 測資的撰寫
```golang
package balance_string

import "testing"

func Test_balanceStringSplit(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example1",
			args: args{
				s: "RLRRLLRLRL",
			},
			want: 4,
		},
		{
			name: "Example2",
			args: args{
				s: "RLLLLRRRLR",
			},
			want: 3,
		},
		{
			name: "Example3",
			args: args{
				s: "LLLLRRRR",
			},
			want: 1,
		},
		{
			name: "Example4",
			args: args{
				s: "RLRRRLLRLL",
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := balanceStringSplit(tt.args.s); got != tt.want {
				t.Errorf("balanceStringSplit() = %v, want %v", got, tt.want)
			}
		})
	}
}

```
## my self record
[golang ironman30day 15th day](https://hackmd.io/HSRr63r7S0yB3UnCvznaag?view)
## 參考文章

[golang test](https://ithelp.ithome.com.tw/articles/10204692)