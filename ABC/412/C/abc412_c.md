配点 : $300$ 点

### 問題文

$1$ から $N$ までの番号がついた $N$ 個のドミノがあります。ドミノ $i$ の大きさは $S_i$ です。  
いくつかのドミノを左右一列に並べたあとにドミノを倒すことを考えます。ドミノ $i$ が右に向けて倒れる時、ドミノ $i$ のすぐ右に置かれているドミノの大きさが $2 S_i$ 以下ならばそのドミノも右に向けて倒れます。

あなたは $2$ 個以上のドミノを選んで左右一列に並べることにしました。ただし、ドミノの並べ方は次の条件を満たす必要があります。 

  * 一番左のドミノはドミノ $1$ である。
  * 一番右のドミノはドミノ $N$ である。
  * ドミノ $1$ のみを右に向けて倒した時に、最終的にドミノ $N$ も右に向けて倒れる。



条件を満たすドミノの並べ方は存在しますか？また、存在する場合は最小で何個のドミノを並べる必要がありますか？

$T$ 個のテストケースが与えられるので、それぞれについて問題を解いてください。

### 制約

  * $1 \leq T \leq 10^5$
  * $2 \leq N \leq 2 \times 10^5$
  * $1 \leq S_i \leq 10^9$
  * 全てのテストケースに対する $N$ の総和は $2 \times 10^5$ 以下
  * 入力される値は全て整数



* * *

### 入力

入力は以下の形式で標準入力から与えられる。ここで $\mathrm{case}_i$ は $i$ 番目のテストケースを意味する。
    
    
    T
    \mathrm{case}_1
    \mathrm{case}_2
    \vdots
    \mathrm{case}_T

各テストケースは以下の形式で与えられる。
    
    
    N
    S_1 S_2 \dots S_N

### 出力

$T$ 行出力せよ。$i$ 行目には $i$ 番目のテストケースの答えを出力せよ。  
各テストケースでは、条件を満たすドミノの並べ方が存在しない場合は `-1` を、存在する場合は並べるドミノの最小個数を出力せよ。

* * *

### 入力例 1
    
    
    3
    4
    1 3 2 5
    2
    1 100
    10
    298077099 766294630 440423914 59187620 725560241 585990757 965580536 623321126 550925214 917827435

### 出力例 1
    
    
    4
    -1
    3

$1$ 番目のテストケースについて、ドミノを左から順にドミノ $1$, ドミノ $3$, ドミノ $2$, ドミノ $4$ の順に並べることで問題文の条件を満たすことができます。特に $3$ 番目の条件については、ドミノ $1$ のみを右に向けて倒した時に以下の順にドミノが倒れます。

  * ドミノ $1$ の右にはドミノ $3$ が置かれている。ドミノ $3$ の大きさ $S_3 = 2$ は $S_1 \times 2 = 1 \times 2 = 2$ 以下であるから、ドミノ $3$ も右に向けて倒れる。
  * ドミノ $3$ の右にはドミノ $2$ が置かれている。ドミノ $2$ の大きさ $S_2 = 3$ は $S_3 \times 2 = 2 \times 2 = 4$ 以下であるから、ドミノ $2$ も右に向けて倒れる。
  * ドミノ $2$ の右にはドミノ $4$ が置かれている。ドミノ $4$ の大きさ $S_4 = 5$ は $S_2 \times 2 = 3 \times 2 = 6$ 以下であるから、ドミノ $4$ も右に向けて倒れる。



$3$ 個以下のドミノを並べて問題文の条件を達成することはできないので、答えは $4$ 個です。