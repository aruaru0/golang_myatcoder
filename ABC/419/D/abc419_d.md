配点 : $400$ 点

### 問題文

長さ $N$ の英小文字列 $S, T$ および $M$ 個の整数の組 $(L_1,R_1),(L_2,R_2),\ldots,(L_M,R_M)$ が与えられます。

$i=1,2,\ldots,M$ の順に以下の操作を行います。

  * $S$ の $L_i$ 文字目から $R_i$ 文字目と、$T$ の $L_i$ 文字目から $R_i$ 文字目を入れ替える。
    * 例えば、$S$ が `abcdef`、$T$ が `ghijkl`、$(L_i,R_i)=(3,5)$ のとき、操作後の $S,T$ はそれぞれ `abijkf`、`ghcdel` となる。



$M$ 回の操作を行った後の $S$ を求めてください。

### 制約

  * $1\leq N\leq 5\times 10^5$
  * $1\leq M\leq 2\times 10^5$
  * $S,T$ は長さ $N$ の英小文字列
  * $1\leq L_i\leq R_i\leq N$
  * $N,M,L_i,R_i$ は整数



* * *

### 入力

入力は以下の形式で標準入力から与えられる。
    
    
    N M
    S
    T
    L_1 R_1
    L_2 R_2
    \vdots
    L_M R_M

### 出力

$M$ 回の操作を行った後の $S$ を出力せよ。

* * *

### 入力例 1
    
    
    5 3
    apple
    lemon
    2 4
    1 5
    5 5

### 出力例 1
    
    
    lpple

はじめ $S,T$ はそれぞれ `apple`, `lemon` です。

  * $i=1$ の操作後の $S,T$ はそれぞれ `aemoe`, `lppln` です。
  * $i=2$ の操作後の $S,T$ はそれぞれ `lppln`, `aemoe` です。
  * $i=3$ の操作後の $S,T$ はそれぞれ `lpple`, `aemon` です。



よって、$3$ 回の操作を行った後の $S$ は `lpple` です。

* * *

### 入力例 2
    
    
    10 5
    lemwrbogje
    omsjbfggme
    5 8
    4 8
    1 3
    6 6
    1 4

### 出力例 2
    
    
    lemwrfogje