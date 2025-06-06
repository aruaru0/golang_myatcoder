配点 : $400$ 点

### 問題文

AtCoder 国には動物園が $N$ 園あり、$1$ から $N$ の番号がついています。動物園 $i$ の入園料は $C_i$ 円です。

鈴木さんは $M$ 種類の動物、動物 $1,\ldots,M$ が好きです。  
動物 $i$ は $K_i$ 園の動物園 $A_{i,1},\dots, A_{i,K_i}$ で見ることができます。

$M$ 種類の動物全てを $2$ 度以上ずつ見るために必要な入園料の合計の最小値を求めてください。  
なお、同じ動物園を複数回訪れた場合、その動物園の動物は訪れた回数だけ見たとみなします。

### 制約

  * $1\leq N \leq 10$
  * $1\leq M \leq 100$
  * $0\leq C_i \leq 10^9$
  * $1\leq K_i \leq N$
  * $1 \leq A_{i,j} \leq N$
  * $j \neq j' \Longrightarrow A_{i,j}\neq A_{i,j'}$
  * 入力は全て整数



* * *

### 入力

入力は以下の形式で標準入力から与えられる。
    
    
    N M
    C_1 \dots C_N
    K_1 A_{1,1} \dots A_{1,K_1}
    \vdots
    K_M A_{M,1} \dots A_{M,K_M}

### 出力

答えを出力せよ。

* * *

### 入力例 1
    
    
    4 3
    1000 300 700 200
    3 1 3 4
    3 1 2 4
    2 1 3

### 出力例 1
    
    
    1800

以下のようにすることで、$1800$ 円で動物 $1,2,3$ を $2$ 度以上ずつ見ることができます。

  * 動物園 $3$ に行く。入園料 $700$ 円を払い、動物 $1,3$ を見る。
  * 動物園 $3$ に行く。入園料 $700$ 円を払い、動物 $1,3$ を見る。
  * 動物園 $4$ に行く。入園料 $200$ 円を払い、動物 $1,2$ を見る。
  * 動物園 $4$ に行く。入園料 $200$ 円を払い、動物 $1,2$ を見る。



* * *

### 入力例 2
    
    
    7 6
    500 500 500 500 500 500 1000
    3 1 2 7
    3 2 3 7
    3 3 4 7
    3 4 5 7
    3 5 6 7
    3 6 1 7

### 出力例 2
    
    
    2000

動物園 $7$ に $2$ 度行くことで、合計 $2000$ 円で動物 $1,2,3,4,5,6$ を $2$ 度ずつ見ることができます。