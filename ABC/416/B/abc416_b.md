配点 : $250$ 点

### 問題文

`.` および `#` からなる文字列 $S$ が与えられます。

以下の条件を全て満たす文字列 $T$ のうち、 `o` の文字数が最大となるものを一つ求めてください。

  * $T$ の長さは $S$ の長さと等しい。
  * $T$ は `.`、`#`、`o` からなる。
  * $S_i=$ `#` であるとき、またそのときに限り $T_i=$ `#` である。
  * $T_i=T_j=$ `o` $(i < j)$ ならば、 $T_{i+1},\ldots,T_{j-1}$ の中に `#` が $1$ つ以上存在する。



### 制約

  * $S$ は `.` および `#` からなる長さ $1$ 以上 $100$ 以下の文字列



* * *

### 入力

入力は以下の形式で標準入力から与えられる。
    
    
    S

### 出力

条件を全て満たす文字列 $T$ のうち、 `o` の文字数が最大となるものを一つ出力せよ。

そのような文字列が複数ある場合、どれを出力しても正答となる。

* * *

### 入力例 1
    
    
    #..#.

### 出力例 1
    
    
    #o.#o

$T=$ `#o.#o` とすると全ての条件を満たすことが確認できます。

全ての条件を満たす $T$ であって、 `o` の文字数が $2$ より多い文字列は存在しないので `#o.#o` を出力すると正答となります。

この他にも `#.o#o` を出力しても正答となります。

* * *

### 入力例 2
    
    
    #

### 出力例 2
    
    
    #

* * *

### 入力例 3
    
    
    .....

### 出力例 3
    
    
    ..o..

この他にも `o....`、`.o...`、`...o.`、`....o` を出力しても正答となります。

* * *

### 入力例 4
    
    
    ...#..#.##.#.

### 出力例 4
    
    
    o..#.o#o##o#o