::: {#modal-contest-start .modal .fade tabindex="-1" role="dialog"}
::: {.modal-dialog role="document"}
::: modal-content
::: modal-header
[×]{aria-hidden="true"}

#### Contest started {#contest-started .modal-title}
:::

::: modal-body
C++入門 AtCoder Programming Guide for beginners (APG4b) has begun.
:::

::: modal-footer
Close
:::
:::
:::
:::

::: {#modal-contest-end .modal .fade tabindex="-1" role="dialog"}
::: {.modal-dialog role="document"}
::: modal-content
::: modal-header
[×]{aria-hidden="true"}

#### Contest is over {#contest-is-over .modal-title}
:::

::: modal-body
C++入門 AtCoder Programming Guide for beginners (APG4b) has ended.
:::

::: modal-footer
Close
:::
:::
:::
:::

::: {#main-div .float-container}
::: container-fluid
::: navbar-header
[]{.icon-bar}[]{.icon-bar}[]{.icon-bar}

[](/home){.navbar-brand}
:::

::: {#navbar-collapse .collapse .navbar-collapse}
-   [C++入門 AtCoder Programming Guide for beginners
    (APG4b)](/contests/APG4b){.contest-title}

```{=html}
<!-- -->
```
-   [![](//img.atcoder.jp/assets/top/img/flag-lang/en.png) English
    []{.caret}](#){.dropdown-toggle toggle="dropdown" role="button"
    aria-haspopup="true" aria-expanded="false"}
    -   [![](//img.atcoder.jp/assets/top/img/flag-lang/ja.png)
        日本語](/contests/APG4b/tasks/APG4b_ce?lang=ja)
    -   [![](//img.atcoder.jp/assets/top/img/flag-lang/en.png)
        English](/contests/APG4b/tasks/APG4b_ce?lang=en)
-   [Sign
    Up](/register?continue=https%3A%2F%2Fatcoder.jp%2Fcontests%2FAPG4b%2Ftasks%2FAPG4b_ce)
-   [Sign
    In](/login?continue=https%3A%2F%2Fatcoder.jp%2Fcontests%2FAPG4b%2Ftasks%2FAPG4b_ce)
:::
:::

::: {#main-container .container style="padding-top:50px;"}
::: row
::: {#contest-nav-tabs .col-sm-12 .mb-2 .cnvtb-fixed}
<div>

[Contest Duration: [2017-12-13
21:00:00+0900](http://www.timeanddate.com/worldclock/fixedtime.html?iso=20171213T2100&p1=248){target="blank"} -
[4017-12-13
21:00:00+0900](http://www.timeanddate.com/worldclock/fixedtime.html?iso=40171213T2100&p1=248){target="blank"}
(local time)]{.small} [[Back to Home](/home)]{.small}

</div>

-   [[]{.glyphicon .glyphicon-home aria-hidden="true"}
    Top](/contests/APG4b)
-   [[]{.glyphicon .glyphicon-tasks aria-hidden="true"}
    Tasks](/contests/APG4b/tasks)
-   [[]{.glyphicon .glyphicon-question-sign aria-hidden="true"}
    Clarifications []{#clar-badge
    .badge}](/contests/APG4b/clarifications)
-   [[]{.glyphicon .glyphicon-list aria-hidden="true"}
    Results[]{.caret}](#){.dropdown-toggle toggle="dropdown"
    role="button" aria-haspopup="true" aria-expanded="false"}
    -   [[]{.glyphicon .glyphicon-globe aria-hidden="true"} All
        Submissions](/contests/APG4b/submissions)
-   [[]{.glyphicon .glyphicon-book aria-hidden="true"}
    Editorial](/contests/APG4b/editorial)
-   [[]{.glyphicon .glyphicon-pushpin
    aria-hidden="true"}](javascript:void(0)){#fix-cnvtb}
:::

::: col-sm-12
[ EX18 - 2.03 [Editorial](/contests/APG4b/tasks/APG4b_ce/editorial){.btn
.btn-default .btn-sm} ]{.h2}
[[![](//img.atcoder.jp/assets/top/img/flag-lang/ja.png)]{data-lang="ja"}
/
[![](//img.atcoder.jp/assets/top/img/flag-lang/en.png)]{data-lang="en"}]{#task-lang-btn
.pull-right}

------------------------------------------------------------------------

Time Limit: 2 sec / Memory Limit: 256 MB

::: {#task-statement}
[[ ]{.lang-ja}]{.lang}

[説明ページに戻る](https://atcoder.jp/contests/apg4b/tasks/APG4b_t)

::: part
::: section
### 問題文

あるゲーム大会には`N`{.variable}人が参加し`M`{.variable}試合が行われました。各参加者には`1`{.variable}から`N`{.variable}の番号が割り当てられています。

試合に関する情報が与えられるので、`M`{.variable}回の試合がすべて終了した時点での*試合結果の表*を作成し、出力してください。

ただし、同じ参加者のペアについて2回以上試合が行われることはないとします。

------------------------------------------------------------------------

試合に関する情報は以下のような形式で与えられます。

##### 試合に関する情報

``` {.io-style style="font-size:16px"}
試合1で勝った人の番号A_1 試合1で負けた人の番号B_1
試合2で勝った人の番号A_2 試合2で負けた人の番号B_2
\vdots \vdots
試合Mで勝った人の番号A_M 試合Mで負けた人の番号B_M
```

同じ参加者のペアについて2回以上試合が行われることはありません。\
例えば、次のような情報が与えられることはありません。

``` {.io-style style="font-size:16px"}
1 2
2 1
```

------------------------------------------------------------------------

*試合結果の表*とは、縦N行、横N列からなる次のような表`R`{.variable}です。

##### 試合結果の表

``` {.io-style style="font-size:16px"}
R_{1, 1} R_{1, 2} R_{1, 3} \cdots R_{1, N}
R_{2, 1} R_{2, 2} R_{2, 3} \cdots R_{2, N}
R_{3, 1} R_{3, 2} R_{3, 3} \cdots R_{3, N}
\vdots  \vdots  \vdots  \vdots  \ddots  \vdots
R_{N, 1} R_{N, 2} R_{N, 3} \cdots R_{N, N}
```

`R_{i, j}`{.variable}の値は以下のように決まります。

`i`{.variable}番の参加者と`j`{.variable}番の参加者が試合をして、

-   `i`{.variable}番の参加者が勝ったなら`R_{i, j}`{.variable} = `o`
-   `i`{.variable}番の負けたなら`R_{i, j}`{.variable} = `x`

`i`{.variable}番の人と`j`{.variable}番の人が試合を行っていない場合

-   `R_{i, j}`{.variable} = `-`

------------------------------------------------------------------------

以下に具体例を示します。

#### 具体例

-   3人が参加した
-   2試合行われた
-   試合に関する情報は次の通り

``` {.io-style style="font-size:16px"}
1 2
3 1
```

この場合の*試合結果の表*は次のようになります。

``` {.io-style style="font-size:16px"}
- o x
x - -
o - -
```

-   1番の人と2番の人が試合を行い、1番の人が勝ったので、`R_{1, 2}`{.variable}
    = `o`、`R_{2, 1}`{.variable} = `x`
-   3番の人と1番の人が試合を行い、3番の人が勝ったので、`R_{1, 3}`{.variable}
    = `x`、`R_{3, 1}`{.variable} = `o`

ただし、**各行の行末に空白を含まない**ことに注意してください。

------------------------------------------------------------------------

##### サンプルプログラム

``` {.prettyprint .linenums .source-code}
#include <bits/stdc++.h>
using namespace std;

int main() {
  int N, M;
  cin >> N >> M;
  vector<int> A(M), B(M);
  for (int i = 0; i < M; i++) {
    cin >> A.at(i) >> B.at(i);
  }

  // ここにプログラムを追記
  // (ここで"試合結果の表"の2次元配列を宣言)
}
```

##### 行末に空白を含めない出力の仕方

以下は配列の要素を空白区切りで出力し末尾には空白を含めないようにする方法の1例です。

``` {.prettyprint .source-code}
vector<int> a = {1, 2, 3, 4, 5};
for (int i = 0; i < 5; i++) {
  cout << a.at(i);
  if (i == 4) {
    cout << endl; // 末尾なら改行
  }
  else {
    cout << " ";  // それ以外なら空白
  }
}
```

------------------------------------------------------------------------
:::
:::

::: part
::: section
### 制約

-   `1≦N≦100`{.variable}
-   `0≦M≦4950`{.variable}
-   `1≦A_i, B_i≦N (1 ≦ i ≦ M)`{.variable}
-   `A_i \neq B_i (1 ≦ i ≦ M)`{.variable}
-   同じ参加者のペアで2回以上試合が行われることはない
-   入力はすべて整数

------------------------------------------------------------------------
:::
:::

::: part
::: section
### 入力

入力は次の形式で標準入力から与えられます。

``` {.io-style style="font-size:16px"}
N M
A_1 B_1
A_2 B_2
A_3 B_3
: :
A_M B_M
```
:::
:::

::: part
::: section
### 出力

`M`{.variable}試合が終了した時点での*試合結果の表*を出力してください。

ただし、**各行の行末に空白を含まない**ことに注意してください。

------------------------------------------------------------------------

**ジャッジでは以下の入力例以外のケースに関してもテストされる**ことに注意。
:::
:::

::: part
::: section
### 入力例1

    3 2
    1 2
    3 1
:::
:::

::: part
::: section
### 出力例1

    - o x
    x - -
    o - -

-   1番の人と2番の人が試合を行い、1番の人が勝ったので、`R_{1, 2}`{.variable}
    = `o`、`R_{2, 1}`{.variable} = `x`
-   3番の人と1番の人が試合を行い、3番の人が勝ったので、`R_{1, 3}`{.variable}
    = `x`、`R_{3, 1}`{.variable} = `o`
:::
:::

::: part
::: section
### 入力例2

    7 12
    1 5
    5 4
    6 2
    1 7
    2 4
    6 3
    1 3
    6 4
    3 7
    5 7
    4 3
    6 1
:::
:::

::: part
::: section
### 出力例2

    - - o - o x o
    - - - o - x -
    x - - x - x o
    - x o - x x -
    x - - o - - o
    o o o o - - -
    x - x - x - -
:::
:::

::: part
::: section
### 入力例3

    1 0
:::
:::

::: part
::: section
### 出力例3

    -

------------------------------------------------------------------------
:::
:::

::: part
::: section
### ヒント

クリックでヒントを開く

まずは、*試合結果の表*をプログラムで管理するために、N×Nの2次元配列を用意しましょう。
この2次元配列の型は、表の要素となる`-`, `o`,
`x`の型なので、`char`型にすれば良いです。

縦a行、横b列の2次元配列の宣言方法は次の通りでした。

``` {.prettyprint .source-code}
vector<vector<要素の型>> 変数名(a, vector<要素の型>(b));
```

------------------------------------------------------------------------
:::
:::

::: part
::: section
### テスト入出力

**書いたプログラムがACにならず、原因がどうしてもわからないときだけ見てください。**

クリックでテスト入出力を見る

##### テスト入力1

    4 6
    1 2
    1 3
    1 4
    2 3
    2 4
    3 4

##### テスト出力1

    - o o o
    x - o o
    x x - o
    x x x -

------------------------------------------------------------------------

##### テスト入力2

    20 56
    14 18
    8 19
    18 20
    5 15
    17 3
    12 15
    7 3
    14 12
    18 17
    2 12
    4 12
    17 5
    11 10
    14 13
    8 5
    8 1
    16 13
    17 7
    16 18
    20 8
    10 7
    9 20
    17 11
    8 2
    6 4
    9 19
    13 3
    7 15
    13 9
    4 2
    18 7
    20 2
    17 2
    8 18
    5 16
    1 12
    6 1
    11 2
    9 6
    11 15
    17 19
    18 6
    8 17
    15 10
    10 6
    1 18
    15 19
    9 7
    2 14
    4 19
    20 11
    16 12
    5 2
    16 9
    13 2
    12 20

##### テスト出力2

    - - - - - x - x - - - o - - - - - o - -
    - - - x x - - x - - x o x o - - x - - x
    - - - - - - x - - - - - x - - - x - - -
    - o - - - x - - - - - o - - - - - - o -
    - o - - - - - x - - - - - - o o x - - -
    o - - o - - - - x x - - - - - - - x - -
    - - o - - - - - x x - - - - o - x x - -
    o o - - o - - - - - - - - - - - o o o x
    - - - - - o o - - - - - x - - x - - o o
    - - - - - o o - - - x - - - x - - - - -
    - o - - - - - - - o - - - - o - x - - x
    x x - x - - - - - - - - - x o x - - - o
    - o o - - - - - o - - - - x - x - - - -
    - x - - - - - - - - - o o - - - - o - -
    - - - - x - x - - o x x - - - - - - o -
    - - - - x - - - o - - o o - - - - o - -
    - o o - o - o x - - o - - - - - - x o -
    x - - - - o o x - - - - - x - x o - - o
    - - - x - - - x x - - - - - x - x - - -
    - o - - - - - o x - o x - - - - - x - -

------------------------------------------------------------------------

##### テスト入力3

データが大きすぎるため省略

------------------------------------------------------------------------
:::
:::

::: part
::: section
### 解答例

**必ず自分で問題に挑戦してみてから見てください。**\

クリックで解答例を見る

``` {.prettyprint .linenums .source-code}
#include <bits/stdc++.h>
using namespace std;

int main() {
  int N, M;
  cin >> N >> M;
  vector<int> A(M), B(M);
  for (int i = 0; i < M; i++) {
    cin >> A.at(i) >> B.at(i);
  }

  // N×Nのchar型の2次元配列のすべての要素を'-'で初期化
  vector<vector<char>> table(N, vector<char>(N, '-'));

  for (int i = 0; i < M; i++) {
    // 1〜N → 0〜N-1 に変換
    A.at(i)--; B.at(i)--;
    table.at(A.at(i)).at(B.at(i)) = 'o';  // AはBに勝った
    table.at(B.at(i)).at(A.at(i)) = 'x';  // BはAに負けた
  }

  for (int i = 0; i < N; i++) {
    for (int j = 0; j < N; j++) {
      cout << table.at(i).at(j);
      if (j == N - 1) {
        cout << endl;  // 行末なら改行
      }
      else {
        cout << " ";  // 行末でないなら空白を出力
      }
    }
  }
}
```
:::
:::
:::
:::
:::

------------------------------------------------------------------------

::: {.a2a_kit .a2a_kit_size_20 .a2a_default_style .pull-right a2a-url="https://atcoder.jp/contests/APG4b/tasks/APG4b_ce?lang=en" a2a-title="EX18 - 2.03"}
[]{.a2a_button_facebook} []{.a2a_button_twitter}
[]{.a2a_button_telegram} [](https://www.addtoany.com/share){.a2a_dd}
:::
:::

------------------------------------------------------------------------
:::

::: {.container style="margin-bottom: 80px;"}
-   [Rule](/contests/APG4b/rules)
-   [Glossary](/contests/APG4b/glossary)

```{=html}
<!-- -->
```
-   [Terms of service](/tos)
-   [Privacy Policy](/privacy)
-   [Information Protection Policy](/personal)
-   [Company](/company)
-   [FAQ](/faq)
-   [Contact](/contact)

::: text-center
[Copyright Since 2012 ©[AtCoder Inc.](http://atcoder.co.jp) All rights
reserved.]{.small}
:::
:::

::: {#scroll-page-top style="display:none;"}
[]{.glyphicon .glyphicon-arrow-up aria-hidden="true"} Page Top
:::
