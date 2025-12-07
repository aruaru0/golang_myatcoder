配点 : $575$ 点

### 問題文

$N$ 個の英小文字からなる文字列 $S_1,S_2,\dots,S_N$ が与えられます。

$(1,2,\dots,N)$ の順列 $P=(P_1,P_2,\dots,P_N)$ としてありえるもの全てについて、以下の通りに生成される文字列を書き並べます。

  * $S_{P_1},S_{P_2},\dots,S_{P_N}$ をこの順に連結する。



書き並べられた $N!$ 個の文字列を辞書順に並べた列を $A_1,A_2,\dots,A_{N!}$ とします。  
$A_2$ を出力してください。

$T$ 個のテストケースが与えられるので、それぞれについて答えを求めてください。

### 制約

  * $1 \le T \le 1.5 \times 10^5$
  * $2 \le N \le 3 \times 10^5$
  * $T,N$ は整数
  * $S_i$ は英小文字からなる長さ $1$ 以上 $10^6-1$ 以下の文字列
  * ひとつの入力について、 $N$ の総和は $3 \times 10^5$ を超えない
  * ひとつの入力について、全てのテストケースにおける $|S_i|$ の総和は $10^6$ を超えない



* * *

### 入力

入力は以下の形式で標準入力から与えられる。  
$\text{case}_i$ は $i$ 番目のテストケースを表す。
    
    
    T
    \text{case}_1
    \text{case}_2
    \vdots
    \text{case}_T

各テストケースは以下の形式で与えられる。
    
    
    N
    S_1
    S_2
    \vdots
    S_N

### 出力

$T$ 行出力せよ。

$i$ 行目には $i$ 番目のテストケースについて、答えを出力せよ。

* * *

### 入力例 1
    
    
    3
    3
    abc
    ac
    ahc
    4
    aaa
    a
    aaaa
    a
    15
    ks
    sy
    k
    ysk
    yks
    ky
    ksy
    sk
    syk
    s
    kys
    sky
    ys
    yk
    y

### 出力例 1
    
    
    abcahcac
    aaaaaaaaa
    kksksykykysskskyssyksyyksykyskyys

この入力には $3$ 個のテストケースが含まれています。

$1$ 番目のテストケースについて、 $S=($ `abc`, `ac`, `ahc` $)$ です。  
$A=($ `abcacahc`, `abcahcac`, `acabcahc`, `acahcabc`, `ahcabcac`, `ahcacabc` $)$ なので、 $A_2=$ `abcahcac` を出力します。