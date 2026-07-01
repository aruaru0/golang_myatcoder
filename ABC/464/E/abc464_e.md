配点 : $450$ 点

### 問題文

$H \times W$ のグリッドがあります。はじめ、すべてのマスに `A` が書かれています。上から $i$ 行目、左から $j$ 列目のマスを $(i,j)$ で表します。

これから $Q$ 回の操作を順に行います。  
$i$ 回目の操作では、左上のマスを $(1,1)$、 右下のマスを $(R_i,C_i)$ とする長方形に含まれるすべてのマスを、英大文字 $X_i$ で上書きします。

すべての操作を行った後のグリッドを出力してください。

### 制約

  * $1 \le H, W$
  * $H \times W \le 10^6$
  * $1 \le Q \le 2 \times 10^5$
  * $1 \le R_i \le H$
  * $1 \le C_i \le W$
  * $X_i$ は英大文字
  * 入力される数値はすべて整数



* * *

### 入力

入力は以下の形式で標準入力から与えられる。
    
    
    H W Q
    R_1 C_1 X_1
    R_2 C_2 X_2
    \vdots
    R_Q C_Q X_Q

### 出力

$H$ 行出力せよ。$i$ 行目には長さ $W$ の文字列であって、$j$ 文字目が操作後のグリッドにおいて $(i, j)$ に書かれている英大文字であるものを出力せよ。

* * *

### 入力例 1
    
    
    2 3 3
    2 2 B
    1 3 C
    2 1 D

### 出力例 1
    
    
    DCC
    DBA

図のように操作が進みます。 

* * *

### 入力例 2
    
    
    1 7 7
    1 7 E
    1 6 C
    1 5 N
    1 4 A
    1 3 V
    1 2 D
    1 1 A

### 出力例 2
    
    
    ADVANCE

* * *

### 入力例 3
    
    
    10 10 15
    8 9 B
    6 7 C
    5 8 D
    10 6 E
    8 5 F
    3 10 G
    7 3 H
    4 6 I
    3 1 J
    10 2 K
    3 6 L
    3 3 M
    2 5 N
    9 1 O
    1 4 P

### 出力例 3
    
    
    PPPPNLGGGG
    ONNNNLGGGG
    OMMLLLGGGG
    OKIIIIDDBA
    OKHFFEDDBA
    OKHFFECBBA
    OKHFFEBBBA
    OKFFFEBBBA
    OKEEEEAAAA
    KKEEEEAAAA