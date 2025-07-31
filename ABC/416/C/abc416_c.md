配点 : $300$ 点

### 問題文

$N$ 個の文字列 $S_1,\ldots,S_N$ が与えられます。

全ての要素が $1$ 以上 $N$ 以下であるような長さ $K$ の数列 $(A_1,\ldots,A_K)$ に対し、 文字列 $f(A_1,\ldots,A_K)$ を $S_{A_1}+S_{A_2}+\dots+S_{A_K}$ と定めます。ここで `+` は文字列の連結を表します。

$N^K$ 個の数列全てについての $f(A_1,\dots,A_K)$ を辞書順に並べたとき、小さい方から $X$ 番目の文字列を求めてください。

### 制約

  * $1\leq N \leq 10$
  * $1\leq K \leq 5$
  * $1\leq X \leq N^K$
  * $S_i$ は英小文字からなる長さ $10$ 以下の文字列
  * $N,K,X$ は全て整数



* * *

### 入力

入力は以下の形式で標準入力から与えられる。
    
    
    N K X
    S_1
    \vdots
    S_N

### 出力

答えを出力せよ。

* * *

### 入力例 1
    
    
    3 2 6
    abc
    xxx
    abc

### 出力例 1
    
    
    abcxxx

  * $f(1,1)=$ `abcabc`
  * $f(1,2)=$ `abcxxx`
  * $f(1,3)=$ `abcabc`
  * $f(2,1)=$ `xxxabc`
  * $f(2,2)=$ `xxxxxx`
  * $f(2,3)=$ `xxxabc`
  * $f(3,1)=$ `abcabc`
  * $f(3,2)=$ `abcxxx`
  * $f(3,3)=$ `abcabc`



であり、これらを辞書順に並べた `abcabc`, `abcabc`, `abcabc`, `abcabc`, `abcxxx`, `abcxxx`, `xxxabc`, `xxxabc`, `xxxxxx` の $6$ 番目は `abcxxx` です。

* * *

### 入力例 2
    
    
    5 5 416
    a
    aa
    aaa
    aa
    a

### 出力例 2
    
    
    aaaaaaa