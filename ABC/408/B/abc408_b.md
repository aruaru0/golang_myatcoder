配点 : $150$ 点

### 問題文

長さ $N$ の整数列 $A=(A_1,A_2,\ldots,A_N)$ が与えられます。

$A$ に含まれる数を重複を除いて小さい順に出力してください。

### 制約

  * $1\le N\le 100$
  * $1\le A_i\le 100$
  * 入力される値は全て整数



* * *

### 入力

入力は以下の形式で標準入力から与えられる。
    
    
    N
    A_1 A_2 \ldots A_N

### 出力

$A$ に含まれる数を小さい順に $C_1,C_2,\ldots , C_M$ として、以下の形式で出力せよ。
    
    
    M
    C_1 C_2 \ldots C_M

* * *

### 入力例 1
    
    
    4
    3 1 4 1

### 出力例 1
    
    
    3
    1 3 4

$A=(3,1,4,1)$ に含まれる数は小さい順に $1,3,4$ の $3$ つです。したがって、上記のように出力してください。

* * *

### 入力例 2
    
    
    3
    7 7 7

### 出力例 2
    
    
    1
    7

* * *

### 入力例 3
    
    
    8
    19 5 5 19 5 19 4 19

### 出力例 3
    
    
    3
    4 5 19