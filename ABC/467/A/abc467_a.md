配点 : $100$ 点

### 問題文

以下の式で計算される値を $\mathrm{BMI}[\mathrm{kg}/\mathrm{m}^2]$ と言います。

  * $\text{体重}[\mathrm{kg}] \div \text{身長}[\mathrm{m}] \div \text{身長}[\mathrm{m}]$



日本では、$\mathrm{BMI}$ が $25 \; \mathrm{kg}/\mathrm{m}^2$ 以上の人は肥満とされます。  
身長 $H[\mathrm{cm}]$、体重 $W[\mathrm{kg}]$ の人が日本で肥満とされるかどうかを判定してください。

### 制約

  * $1 \leq H \leq 300$
  * $1 \leq W \leq 300$
  * 入力される値はすべて整数



* * *

### 入力

入力は以下の形式で標準入力から与えられる。
    
    
    H W

### 出力

身長 $H[\mathrm{cm}]$、体重 $W[\mathrm{kg}]$ の人が日本で肥満とされるならば `Yes` を、そうでないならば `No` を $1$ 行で出力せよ。

* * *

### 入力例 1
    
    
    180 60

### 出力例 1
    
    
    No

身長が $180 \; \mathrm{cm}=1.8 \; \mathrm{m}$ なので、$\mathrm{BMI}$ は $60 \; \mathrm{kg} \div 1.8 \; \mathrm{m} \div 1.8 \; \mathrm{m}=18.5\dots \; \mathrm{kg}/\mathrm{m}^2$ となり、この人は肥満とはされません。

* * *

### 入力例 2
    
    
    182 188

### 出力例 2
    
    
    Yes

* * *

### 入力例 3
    
    
    180 81

### 出力例 3
    
    
    Yes