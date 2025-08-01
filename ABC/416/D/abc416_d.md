配点 : $400$ 点

### 問題文

長さ $N$ の非負整数列 $A=(A_1,A_2,\ldots,A_N), B=(B_1,B_2,\ldots,B_N)$ と正整数 $M$ が与えられます。

$A$ の要素を自由に並び替えることが出来るとき、 $\displaystyle \sum_{i=1}^N \left((A_i+B_i) \bmod M\right)$ としてありうる最小値を求めて下さい。

$T$ 個のテストケースが与えられるので、それぞれについて答えを求めてください。

### 制約

  * $1\le T \le 10^5$
  * $1\le N\le 3\times 10^5$
  * $1\le M\le 10^9$
  * $0\le A_i,B_i < M$
  * 全てのテストケースにおける $N$ の総和は $3\times 10^5$ 以下
  * 入力される値は全て整数



* * *

### 入力

入力は以下の形式で標準入力から与えられる。
    
    
    T
    \text{case}_1
    \text{case}_2
    \vdots
    \text{case}_T

各テストケース $\text{case}_i$ は以下の形式で与えられる。
    
    
    N M
    A_1 A_2 \ldots A_N
    B_1 B_2 \ldots B_N

### 出力

$T$ 行出力せよ。

$j$ 行目には $j$ 番目のテストケースについて、 $\displaystyle \sum_{i=1}^N \left((A_i+B_i) \bmod M\right)$ としてありうる最小値を出力せよ。

* * *

### 入力例 1
    
    
    3
    3 6
    3 1 4
    2 0 1
    1 1000000000
    999999999
    999999999
    10 201
    144 150 176 154 110 187 38 136 111 46
    96 109 73 63 85 1 156 7 13 171

### 出力例 1
    
    
    5
    999999998
    619

$1$ つ目のテストケースについて、 $A$ を $4,3,1$ と並び替えると $(A_i+B_i)\bmod M$ はそれぞれ $0,3,2$ となり、これらの総和は $5$ となります。