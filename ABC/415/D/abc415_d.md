配点 : $400$ 点

### 問題文

> 本問題の設定は G 問題と類似しています。G 問題とは、最大化の対象および制約の一部が異なります。

不思議なコーラショップがあります。 このショップは直接コーラを売ることはしませんが、コーラの空き瓶と新しい瓶入りコーラとの交換サービスを提供しています。

はじめ、高橋君は瓶入りコーラを $N$ 本持っており、これから以下のいずれかの行動を選んで取ることを好きな回数繰り返すことができます（$0$ 回でもよい）。

  * 持っている瓶入りコーラ $1$ 本を飲む。持っている瓶入りコーラの本数が $1$ 減り、空き瓶の本数が $1$ 増える。 （行動前に $1$ 本も瓶入りコーラを持っていない場合、この行動は取れない。）
  * $1$ 以上 $M$ 以下の整数 $i$ を選ぶ。コーラショップに $A_i$ 本の空き瓶を渡し、引き換えに $B_i$ 本の瓶入りコーラおよび記念のシール $1$ 枚をもらう。 （行動前に持っている空き瓶の本数が $A_i$ 未満であるような $i$ は選ぶことができない。 選ぶことのできる $i$ が存在しない場合、この行動は取れない。）



高橋君はシールが大好きです。うまく行動を取ったとき、最大で何枚のシールを手に入れることができますか？

なお、高橋君は最初空き瓶を $1$ 本も持っておらず、シールも $1$ 枚も持っていないものとします。

### 制約

  * $1\leq N \leq 10^{18}$
  * $1\leq M \leq 2\times 10^5$
  * $1\leq B_i < A_i \leq 10^{18}$
  * 入力は全て整数



* * *

### 入力

入力は以下の形式で標準入力から与えられる。
    
    
    N M
    A_1 B_1
    A_2 B_2
    \vdots
    A_M B_M

### 出力

答えを整数として出力せよ。

* * *

### 入力例 1
    
    
    5 3
    5 1
    4 3
    3 1

### 出力例 1
    
    
    3

以下のような行動を考えます。

  * 最初、高橋君は瓶入りコーラを $5$ 本持っている。
  * 持っている瓶入りコーラ $1$ 本を飲むことを $5$ 回繰り返す。高橋君は空き瓶を $5$ 本持った状態になる。
  * $i=2$ として交換を行い、$4$ 本の空き瓶と引き換えに $3$ 本の瓶入りコーラと $1$ 枚のシールをもらう。高橋君は瓶入りコーラを $3$ 本、空き瓶を $1$ 本、シールを $1$ 枚持った状態になる。
  * 持っている瓶入りコーラ $1$ 本を飲むことを $3$ 回繰り返す。高橋君は空き瓶を $4$ 本、シールを $1$ 枚持った状態になる。
  * $i=2$ として交換を行い、$4$ 本の空き瓶と引き換えに $3$ 本の瓶入りコーラと $1$ 枚のシールをもらう。高橋君は瓶入りコーラを $3$ 本、シールを $2$ 枚持った状態になる。
  * 持っている瓶入りコーラ $1$ 本を飲むことを $3$ 回繰り返す。高橋君は空き瓶を $3$ 本、シールを $2$ 枚持った状態になる。
  * $i=3$ として交換を行い、$3$ 本の空き瓶と引き換えに $1$ 本の瓶入りコーラと $1$ 枚のシールをもらう。高橋君は瓶入りコーラを $1$ 本、シールを $3$ 枚持った状態になる。



これらの行動が終了したあと、高橋君は $3$ 枚のシールを持っています。 高橋君がどのように行動を取っても $4$ 枚以上のシールを手に入れることはできないため、答えは $3$ です。

* * *

### 入力例 2
    
    
    3 3
    5 1
    5 1
    4 2

### 出力例 2
    
    
    0

元々持っている瓶入りコーラを飲むことしかできません。

* * *

### 入力例 3
    
    
    415 8
    327 299
    413 396
    99 67
    108 51
    195 98
    262 180
    250 175
    234 187

### 出力例 3
    
    
    11