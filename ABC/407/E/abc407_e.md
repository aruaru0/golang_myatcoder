配点 : $450$ 点

### 問題文

長さ $2N$ の非負整数列 $A = (A_1, \dots, A_{2N})$ が与えられます。

長さ $2N$ の括弧列 $s$ のスコアを、以下で得られる値として定義します。

  * $s$ の $i$ 文字目が `)` であるようなすべての整数 $i$ について $A_i$ の値を $0$ に変えた場合の、$A$ の要素の総和。



長さ $2N$ の正しい括弧列のスコアとしてありうる最大値を求めてください。

$T$ 個のテストケースが与えられるので、それぞれについて答えを求めてください。

正しい括弧列とは 正しい括弧列とは、 `()` である部分文字列を削除することを $0$ 回以上繰り返して空文字列にできる文字列を指します。 

### 制約

  * $1 \leq T \leq 500$ 
  * $1 \leq N \leq 2 \times 10^5$
  * 各入力ファイルについて、すべてのテストケースの $N$ の総和は $2 \times 10^5$ 以下である。
  * $0 \leq A_i \leq 10^9$ ($1 \leq i \leq 2N$)
  * 入力はすべて整数である。



* * *

### 入力

入力は以下の形式で標準入力から与えられる。
    
    
    T
    \textrm{case}_1
    \textrm{case}_2
    \vdots
    \textrm{case}_T

$\textrm{case}_i$ は $i$ 番目のテストケースを表す。各テストケースは以下の形式で与えられる。
    
    
    N
    A_1
    A_2
    \vdots
    A_{2N}

### 出力

$T$ 行出力せよ。$i$ 行目 ($1 \leq i \leq T$) には $i$ 番目のテストケースに対する答えを出力せよ。

* * *

### 入力例 1
    
    
    2
    3
    400
    500
    200
    100
    300
    600
    6
    1000000000
    1000000000
    1000000000
    1000000000
    1000000000
    1000000000
    1000000000
    1000000000
    1000000000
    1000000000
    1000000000
    1000000000

### 出力例 1
    
    
    1200
    6000000000

$1$ 番目のテストケースにおいて、正しい括弧列として `(())()` を選ぶと、そのスコアは $400+500+0+0+300+0=1200$ となります。

$1200$ よりも高いスコアを得るような正しい括弧列は存在しないので、$1$ 番目のテストケースの答えとして $1200$ を出力します。

この入出力例における $2$ 番目のテストケースのように、答えが 32 bit 整数におさまらないことがあることに注意してください。