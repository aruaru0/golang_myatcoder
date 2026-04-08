配点 : $450$ 点

### 問題文

> あなたの今日のラッキーギリシャ文字はシグマです。シグマを $2$ つも使ったこの問題を解けば、きっと幸運が舞い込むことでしょう。

長さ $N$ の正整数列 $A = (A_1, \cdots, A_N)$ および長さ $M$ の正整数列 $B = (B_1, \cdots, B_M)$ が与えられます。

$\displaystyle \sum_{i=1}^{N} \sum_{j=1}^{M} A_i \cdot B_j \cdot (i \bmod j)$ の値を $998244353$ で割ったあまりを求めてください。

### 制約

  * $1 \leq N,M \leq 5 \times 10^5$
  * $1 \leq A_i, B_j \leq 5 \times 10^5$
  * 入力される値はすべて整数



* * *

### 入力

入力は以下の形式で標準入力から与えられる。
    
    
    N M
    A_1 A_2 \cdots A_N
    B_1 B_2 \cdots B_M

### 出力

答えを $1$ 行に出力せよ。

* * *

### 入力例 1
    
    
    6 4
    1 6 9 2 3 1
    1 10 3 7

### 出力例 1
    
    
    508

以下の $24$ 個の値の合計は $508$ です。

  * $A_1 \cdot B_1 \cdot (1 \bmod 1) = 1 \cdot 1 \cdot 0 = 0$
  * $A_1 \cdot B_2 \cdot (1 \bmod 2) = 1 \cdot 10 \cdot 1 = 10$
  * $A_1 \cdot B_3 \cdot (1 \bmod 3) = 1 \cdot 3 \cdot 1 = 3$
  * $A_1 \cdot B_4 \cdot (1 \bmod 4) = 1 \cdot 7 \cdot 1 = 7$
  * $A_2 \cdot B_1 \cdot (2 \bmod 1) = 6 \cdot 1 \cdot 0 = 0$
  * $A_2 \cdot B_2 \cdot (2 \bmod 2) = 6 \cdot 10 \cdot 0 = 0$
  * $A_2 \cdot B_3 \cdot (2 \bmod 3) = 6 \cdot 3 \cdot 2 = 36$
  * $A_2 \cdot B_4 \cdot (2 \bmod 4) = 6 \cdot 7 \cdot 2 = 84$
  * $A_3 \cdot B_1 \cdot (3 \bmod 1) = 9 \cdot 1 \cdot 0 = 0$
  * $A_3 \cdot B_2 \cdot (3 \bmod 2) = 9 \cdot 10 \cdot 1 = 90$
  * $A_3 \cdot B_3 \cdot (3 \bmod 3) = 9 \cdot 3 \cdot 0 = 0$
  * $A_3 \cdot B_4 \cdot (3 \bmod 4) = 9 \cdot 7 \cdot 3 = 189$
  * $A_4 \cdot B_1 \cdot (4 \bmod 1) = 2 \cdot 1 \cdot 0 = 0$
  * $A_4 \cdot B_2 \cdot (4 \bmod 2) = 2 \cdot 10 \cdot 0 = 0$
  * $A_4 \cdot B_3 \cdot (4 \bmod 3) = 2 \cdot 3 \cdot 1 = 6$
  * $A_4 \cdot B_4 \cdot (4 \bmod 4) = 2 \cdot 7 \cdot 0 = 0$
  * $A_5 \cdot B_1 \cdot (5 \bmod 1) = 3 \cdot 1 \cdot 0 = 0$
  * $A_5 \cdot B_2 \cdot (5 \bmod 2) = 3 \cdot 10 \cdot 1 = 30$
  * $A_5 \cdot B_3 \cdot (5 \bmod 3) = 3 \cdot 3 \cdot 2 = 18$
  * $A_5 \cdot B_4 \cdot (5 \bmod 4) = 3 \cdot 7 \cdot 1 = 21$
  * $A_6 \cdot B_1 \cdot (6 \bmod 1) = 1 \cdot 1 \cdot 0 = 0$
  * $A_6 \cdot B_2 \cdot (6 \bmod 2) = 1 \cdot 10 \cdot 0 = 0$
  * $A_6 \cdot B_3 \cdot (6 \bmod 3) = 1 \cdot 3 \cdot 0 = 0$
  * $A_6 \cdot B_4 \cdot (6 \bmod 4) = 1 \cdot 7 \cdot 2 = 14$



* * *

### 入力例 2
    
    
    20 20
    36625 195265 98908 111868 111868 47382 147644 472464 472464 416653 111868 195265 327972 327972 262769 75439 381156 451275 36625 195265
    327972 111868 416653 177330 340019 262769 47382 262769 47382 340019 47382 262769 327972 327972 359676 381156 327972 36625 451275 381156

### 出力例 2
    
    
    58141644