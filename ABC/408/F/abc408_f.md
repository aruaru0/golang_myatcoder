配点 : $500$ 点

### 問題文

$1$ から $N$ までの番号がつけられた $N$ 個の足場が一列に並んでいます。足場 $i(1 \leq i \leq N)$ の高さは $H_i$ です。

高橋君は足場に乗って移動する遊びをすることにしました。 最初、高橋君は整数 $i(1 \leq i \leq N)$ を自由に選び、足場 $i$ に乗ります。

高橋君はある時点で足場 $i$ に乗っている時、以下の条件を満たす整数 $j(1 \leq j \leq N)$ を選び足場 $j$ に移動することができます。

  * $ H_j \leq H_i - D$ かつ $1 \leq |i-j| \leq R$



高橋君は移動を行えなくなるまで移動を繰り返すとき、移動できる回数の最大値を求めてください。

### 制約

  * $1 \leq N \leq 5 \times 10^5$
  * $1 \leq D,R \leq N$
  * $H$ は $(1,2,\ldots,N)$ の順列
  * 入力はすべて整数



* * *

### 入力

入力は以下の形式で標準入力から与えられる。
    
    
    N D R
    H_1 H_2 \ldots H_N

### 出力

答えを出力せよ。

* * *

### 入力例 1
    
    
    5 2 1
    5 3 1 4 2

### 出力例 1
    
    
    2

高橋君は初めに足場 $1$ に乗り、以下のように足場を移動することができます。

  * $1$ 回目の移動: $H_2 \leq H_1 -D$ かつ $|2-1| \leq R$ であるため足場 $2$ に移動することができる。足場 $1$ から足場 $2$ に移動する。
  * $2$ 回目の移動: $H_3 \leq H_2 -D$ かつ $|3-2| \leq R$ であるため足場 $3$ に移動することができる。足場 $2$ から足場 $3$ に移動する。
  * 足場 $3$ の高さは $1$ であるため、これ以上移動できない。



以上のように $2$ 回移動することができます。また、どのように移動する足場を選んでも $3$ 回以上移動することはできません。よって、$2$ を出力します。

* * *

### 入力例 2
    
    
    13 3 2
    13 7 10 1 9 5 4 11 12 2 8 6 3

### 出力例 2
    
    
    3