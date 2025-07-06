配点 : $200$ 点

### 問題文

$N$ 種類の文字列 $S _ 1,S _ 2,\ldots,S _ N$ が与えられます。

あなたは、次の操作を $1$ 度だけ行います。

  * **相異なる** 整数 $i,j\ (1\le i\le N,1\le j\le N)$ を選び、$S _ i$ と $S _ j$ をこの順で連結する。



この操作で連結した結果の文字列としてありえるものは何通りあるか求めてください。

選んだ $(i,j)$ が異なっても、連結した文字列が同じ場合は $1$ 通りと数えることに注意してください。

### 制約

  * $2\le N\le100$
  * $N$ は整数
  * $S _ i$ は英小文字からなる長さ $1$ 以上 $10$ 以下の文字列
  * $S _ i\ne S _ j\ (1\le i\lt j\le N)$



* * *

### 入力

入力は以下の形式で標準入力から与えられる。
    
    
    N
    S _ 1
    S _ 2
    \vdots
    S _ N

### 出力

操作の結果できる文字列が何通りあるかを出力せよ。

* * *

### 入力例 1
    
    
    4
    at
    atco
    coder
    der

### 出力例 1
    
    
    11

できる文字列は、`atatco`, `atcoat`, `atcoder`, `atcocoder`, `atder`, `coderat`, `coderatco`, `coderder`, `derat`, `deratco`, `dercoder` の $11$ 通りです。

よって、`11` を出力してください。

* * *

### 入力例 2
    
    
    5
    a
    aa
    aaa
    aaaa
    aaaaa

### 出力例 2
    
    
    7

できる文字列は、`aaa`, `aaaa`, `aaaaa`, `aaaaaa`, `aaaaaaa`, `aaaaaaaa`, `aaaaaaaaa` の $7$ 通りです。

よって、`7` を出力してください。

* * *

### 入力例 3
    
    
    10
    armiearggc
    ukupaunpiy
    cogzmjmiob
    rtwbvmtruq
    qapfzsitbl
    vhkihnipny
    ybonzypnsn
    esxvgoudra
    usngxmaqpt
    yfseonwhgp

### 出力例 3
    
    
    90