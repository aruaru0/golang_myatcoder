配点 : $200$ 点

### 問題文

> 連長圧縮（ランレングス圧縮）を復元してください。ただし、長すぎる場合には `Too Long` と出力してください。

$N$ 個の文字と整数の組 $(c_1,l_1),(c_2,l_2),\ldots,(c_N,l_N)$ が与えられます。

$l_1$ 個の文字 $c_1$、$l_2$ 個の文字 $c_2$、$\ldots$、$l_N$ 個の文字 $c_N$ をこの順に連結させた文字列を $S$ とします。

$S$ を出力してください。ただし、$S$ の長さが $100$ を超える場合には代わりに `Too Long` と出力してください。

### 制約

  * $1\leq N\leq 100$
  * $1\leq l_i\leq 10^{18}$
  * $N,l_i$ は整数
  * $c_i$ は英小文字
  * $c_i\neq c_{i+1}$



* * *

### 入力

入力は以下の形式で標準入力から与えられる。
    
    
    N
    c_1 l_1
    c_2 l_2
    \vdots
    c_N l_N

### 出力

$S$ の長さが $100$ 以下なら $S$ を、そうでないなら `Too Long` と出力せよ。

* * *

### 入力例 1
    
    
    8
    m 1
    i 1
    s 2
    i 1
    s 2
    i 1
    p 2
    i 1

### 出力例 1
    
    
    mississippi

$S$ は `mississippi` です。$S$ の長さは $100$ 以下であるため $S$ を出力します。

* * *

### 入力例 2
    
    
    7
    a 1000000000000000000
    t 1000000000000000000
    c 1000000000000000000
    o 1000000000000000000
    d 1000000000000000000
    e 1000000000000000000
    r 1000000000000000000

### 出力例 2
    
    
    Too Long

$S$ の長さは $7\times 10^{18}$ であるため、`Too Long` を出力します。

* * *

### 入力例 3
    
    
    1
    a 100

### 出力例 3
    
    
    aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa

* * *

### 入力例 4
    
    
    6
    g 4
    j 1
    m 4
    e 4
    d 3
    i 4

### 出力例 4
    
    
    ggggjmmmmeeeedddiiii