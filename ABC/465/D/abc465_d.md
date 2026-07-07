配点 : $400$ 点

### 問題文

整数 $X,Y$ と $2$ 以上の整数 $K$ が与えられます。

変数 $x$ があり、はじめ $x=X$ です。あなたは $x$ に対して以下の操作を $0$ 回以上何回でも行うことができます：

  * $\displaystyle \left\lfloor \frac xK \right\rfloor=y$ または $\displaystyle \left\lfloor \frac yK \right\rfloor=x$ を満たす整数 $y$ を選び、$x$ の値を $y$ に置き換える。



ここで、実数 $z$ に対し $\displaystyle \left\lfloor z \right\rfloor$ は $z$ 以下の最大の整数として定義されます。

$x=Y$ とするために必要な操作回数の最小値を求めてください。ただし、制約下では有限回の操作で $x=Y$ とする方法が必ず存在することが証明できます。

$T$ 個のテストケースが与えられるので、それぞれについて答えを求めてください。

### 制約

  * $1\le T\le 2\times 10^5$
  * $0\le X,Y\le 10^{18}$
  * $2\le K\le 10^{18}$
  * 入力される値は全て整数



* * *

### 入力

入力は以下の形式で標準入力から与えられる。
    
    
    T
    \text{case}_1
    \text{case}_2
    \vdots
    \text{case}_T

$i$ 番目 $(1\le i\le T)$ のテストケース $\text{case}_i$ は以下の形式で与えられる。
    
    
    X Y K

### 出力

各テストケースに対する答えを順に改行区切りで出力せよ。

* * *

### 入力例 1
    
    
    4
    11 9 3
    0 0 2
    842 180 7
    1948706013487601 48019760148910476 89014537

### 出力例 1
    
    
    2
    0
    7
    5

$1$ 番目のテストケースについて考えます。

以下のように操作することで $2$ 回の操作で $x=Y$ とすることができます：

  * $y=3$ を選ぶ。$\displaystyle\left\lfloor\frac{x}K \right\rfloor=\left\lfloor\frac{11}3 \right\rfloor=3$ よりこの選択は合法である。そして、$x$ の値を $3$ に置き換える。
  * $y=9$ を選ぶ。$\displaystyle\left\lfloor\frac{y}K \right\rfloor=\left\lfloor\frac{9}3 \right\rfloor=3$ よりこの選択は合法である。そして、$x$ の値を $9$ に置き換える。



$2$ 回未満の操作で $x=Y$ とすることはできないので、$1$ 行目には $2$ を出力してください。