配点 : $450$ 点

### 問題文

整数 $N,M$ と長さ $N$ の整数列 $X=(X_1,X_2,\ldots,X_N)$ 、長さ $M$ の整数列 $Y=(Y_1,Y_2,\ldots,Y_M)$ が与えられます。

以下の条件を全て満たす $N$ 行 $M$ 列の整数行列 $A=(A_{i,j})$ $(1\le i\le N,\ 1\le j\le M)$ が存在するか判定し、存在する場合は一つ求めてください。

  * $1\le A_{i,j} \le N\times M$
  * $A_{i,j}$ の $N\times M$ 個の要素は相異なる
  * $i=1,2,\ldots,N$ に対し $\displaystyle \max_{1\le j\le M} A_{i,j} = X_i$ が成り立つ
  * $j=1,2,\ldots,M$ に対し $\displaystyle \max_{1\le i\le N} A_{i,j} = Y_j$ が成り立つ



$T$ 個のテストケースが与えられるので、それぞれについて答えを求めてください。

### 制約

  * $1\le T\le 10^5$
  * $1\le N,M$
  * 全てのテストケースにおける $N\times M$ の総和は $2\times 10^5$ 以下
  * $1\le X_i,Y_j\le N\times M$
  * 入力される値は全て整数



* * *

### 入力

入力は以下の形式で標準入力から与えられる。
    
    
    T
    \text{case}_1
    \text{case}_2
    \vdots
    \text{case}_T

各テストケースは以下の形式で与えられる。
    
    
    N M
    X_1 X_2 \ldots X_N
    Y_1 Y_2 \ldots Y_M

### 出力

各テストケースに対する答えを順に改行区切りで出力せよ。

各テストケースについて、条件を全て満たす $A$ が存在しない場合は `No` を出力せよ。

そうでない場合、条件を全て満たす $A$ を以下の形式で出力せよ。
    
    
    Yes
    A_{1,1} A_{1,2} \ldots A_{1,M}
    A_{2,1} A_{2,2} \ldots A_{2,M}
    \vdots
    A_{N,1} A_{N,2} \ldots A_{N,M}

条件を満たす $A$ が複数存在する場合、どれを出力しても正答となる。

* * *

### 入力例 1
    
    
    3
    2 3
    5 6
    5 3 6
    3 3
    5 4 6
    6 2 4
    5 4
    18 20 19 14 17
    18 20 14 15

### 出力例 1
    
    
    Yes
    5 1 4
    2 3 6
    No
    Yes
    18 12 4 9
    13 20 1 10
    16 19 6 8
    2 5 14 3
    11 17 7 15

$1$ つ目のテストケースについて考えます。

出力例の $A$ の要素は全て $1$ 以上 $6$ 以下で相異なり、さらに

  * $\displaystyle \max_{1\le j\le 3} A_{1,j} =\max \lbrace 5,1,4\rbrace = 5 = X_1$
  * $\displaystyle\max_{1\le j\le 3} A_{2,j} =\max \lbrace 2,3,6\rbrace = 6 = X_2$
  * $\displaystyle\max_{1\le i\le 2} A_{i,1} =\max \lbrace 5,2\rbrace = 5 = Y_1$
  * $\displaystyle\max_{1\le i\le 2} A_{i,2} =\max \lbrace 1,3\rbrace = 3 = Y_2$
  * $\displaystyle\max_{1\le i\le 2} A_{i,3} =\max \lbrace 4,6\rbrace = 6 = Y_3$



より全ての条件を満たしていることが分かります。

そのほかにも、例えば以下のような出力も正答となります。
    
    
    Yes
    5 3 1
    4 2 6