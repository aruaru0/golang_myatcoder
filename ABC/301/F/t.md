::: {#modal-contest-start .modal .fade tabindex="-1" role="dialog"}
::: {.modal-dialog role="document"}
::: modal-content
::: modal-header
[×]{aria-hidden="true"}

#### Contest started {#contest-started .modal-title}
:::

::: modal-body
Panasonic Programming Contest 2023（AtCoder Beginner Contest 301） has
begun.
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
Panasonic Programming Contest 2023（AtCoder Beginner Contest 301） has
ended.
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
-   [Panasonic Programming Contest 2023（AtCoder Beginner Contest
    301）](/contests/abc301){.contest-title}

```{=html}
<!-- -->
```
-   [![](//img.atcoder.jp/assets/top/img/flag-lang/en.png) English
    []{.caret}](#){.dropdown-toggle toggle="dropdown" role="button"
    aria-haspopup="true" aria-expanded="false"}
    -   [![](//img.atcoder.jp/assets/top/img/flag-lang/ja.png)
        日本語](/contests/abc301/tasks/abc301_f?lang=ja)
    -   [![](//img.atcoder.jp/assets/top/img/flag-lang/en.png)
        English](/contests/abc301/tasks/abc301_f?lang=en)
-   [Sign
    Up](/register?continue=https%3A%2F%2Fatcoder.jp%2Fcontests%2Fabc301%2Ftasks%2Fabc301_f)
-   [Sign
    In](/login?continue=https%3A%2F%2Fatcoder.jp%2Fcontests%2Fabc301%2Ftasks%2Fabc301_f)
:::
:::

::: {#main-container .container style="padding-top:50px;"}
::: row
::: {#contest-nav-tabs .col-sm-12 .mb-2 .cnvtb-fixed}
<div>

[Contest Duration: [2023-05-13
21:00:00+0900](http://www.timeanddate.com/worldclock/fixedtime.html?iso=20230513T2100&p1=248){target="blank"} -
[2023-05-13
22:45:00+0900](http://www.timeanddate.com/worldclock/fixedtime.html?iso=20230513T2245&p1=248){target="blank"}
(local time) (105 minutes)]{.small} [[Back to Home](/home)]{.small}

</div>

-   [[]{.glyphicon .glyphicon-home aria-hidden="true"}
    Top](/contests/abc301)
-   [[]{.glyphicon .glyphicon-tasks aria-hidden="true"}
    Tasks](/contests/abc301/tasks)
-   [[]{.glyphicon .glyphicon-question-sign aria-hidden="true"}
    Clarifications []{#clar-badge
    .badge}](/contests/abc301/clarifications)
-   [[]{.glyphicon .glyphicon-list aria-hidden="true"}
    Results[]{.caret}](#){.dropdown-toggle toggle="dropdown"
    role="button" aria-haspopup="true" aria-expanded="false"}
    -   [[]{.glyphicon .glyphicon-globe aria-hidden="true"} All
        Submissions](/contests/abc301/submissions)
-   [[]{.glyphicon .glyphicon-sort-by-attributes-alt aria-hidden="true"}
    Standings](/contests/abc301/standings)
-   [[]{.glyphicon .glyphicon-sort-by-attributes-alt aria-hidden="true"}
    Virtual Standings](/contests/abc301/standings/virtual)
-   [[]{.glyphicon .glyphicon-book aria-hidden="true"}
    Editorial](/contests/abc301/editorial)
-   [[]{.glyphicon .glyphicon-pushpin
    aria-hidden="true"}](javascript:void(0)){#fix-cnvtb}
:::

::: col-sm-12
[ F - Anti-DDoS
[Editorial](/contests/abc301/tasks/abc301_f/editorial){.btn .btn-default
.btn-sm} ]{.h2}
[[![](//img.atcoder.jp/assets/top/img/flag-lang/ja.png)]{data-lang="ja"}
/
[![](//img.atcoder.jp/assets/top/img/flag-lang/en.png)]{data-lang="en"}]{#task-lang-btn
.pull-right}

------------------------------------------------------------------------

Time Limit: 2 sec / Memory Limit: 1024 MB

::: {#task-statement}
[ [ ]{.lang-ja}]{.lang}

配点 : `500`{.variable} 点

::: part
::: section
### 問題文

英大文字・英小文字からなる長さ `4`{.variable} の文字列で、以下の
`2`{.variable} 条件をともに満たすものを `DDoS`
型文字列と呼ぶことにします。

-   `1,2,4`{.variable} 文字目が英大文字で、`3`{.variable}
    文字目が英小文字である
-   `1,2`{.variable} 文字目が等しい

例えば `DDoS`, `AAaA` は `DDoS` 型文字列であり、`ddos`, `IPoE` は `DDoS`
型文字列ではありません。

英大文字・英小文字および `?` からなる文字列 `S`{.variable}
が与えられます。 `S`{.variable} に含まれる `?`
を独立に英大文字・英小文字に置き換えてできる文字列は、`S`{.variable}
に含まれる `?` の個数を `q`{.variable} として `52^q`{.variable}
通りあります。 このうち `DDoS` 型文字列を部分列に含まないものの個数を
`998244353`{.variable} で割ったあまりを求めてください。
:::
:::

::: part
::: section
### 注記

文字列の**部分列**とは、文字列から `0`{.variable}
個以上の文字を取り除いた後、残りの文字を元の順序で連結して得られる文字列のことをいいます。\
例えば、`AC` は `ABC` の部分列であり、`RE` は `ECR`
の部分列ではありません。
:::
:::

::: part
::: section
### 制約

-   `S`{.variable} は英大文字・英小文字および `?` からなる
-   `S`{.variable} の長さは `4`{.variable} 以上
    `3\times 10^5`{.variable} 以下
:::
:::

------------------------------------------------------------------------

::: io-style
::: part
::: section
### 入力

入力は以下の形式で標準入力から与えられる。

    S
:::
:::

::: part
::: section
### 出力

答えを出力せよ。
:::
:::
:::

------------------------------------------------------------------------

::: part
::: section
### 入力例 1

    DD??S
:::
:::

::: part
::: section
### 出力例 1

    676

`?` の少なくとも一方が英小文字のとき、`DDoS`
型文字列を部分列に含みます。
:::
:::

------------------------------------------------------------------------

::: part
::: section
### 入力例 2

    ????????????????????????????????????????
:::
:::

::: part
::: section
### 出力例 2

    858572093

`998244353`{.variable} で割ったあまりを求めてください。
:::
:::

------------------------------------------------------------------------

::: part
::: section
### 入力例 3

    ?D??S
:::
:::

::: part
::: section
### 出力例 3

    136604
:::
:::

[ ]{.lang-en}

Score : `500`{.variable} points

::: part
::: section
### Problem Statement

A `DDoS`-type string is a string of length `4`{.variable} consisting of
uppercase and lowercase English letters satisfying both of the following
conditions.

-   The first, second, and fourth characters are uppercase English
    letters, and the third character is a lowercase English letter.
-   The first and second characters are equal.

For instance, `DDoS` and `AAaA` are `DDoS`-type strings, while neither
`ddos` nor `IPoE` is.

You are given a string `S`{.variable} consisting of uppercase and
lowercase English letters and `?`. Let `q`{.variable} be the number of
occurrences of `?` in `S`{.variable}. There are `52^q`{.variable}
strings that can be obtained by independently replacing each `?` in
`S`{.variable} with an uppercase or lowercase English letter. Among
these strings, find the number of ones that do not contain a `DDoS`-type
string as a subsequence, modulo `998244353`{.variable}.
:::
:::

::: part
::: section
### Notes

A **subsequence** of a string is a string obtained by removing zero or
more characters from the string and concatenating the remaining
characters without changing the order.\
For instance, `AC` is a subsequence of `ABC`, while `RE` is not a
subsequence of `ECR`.
:::
:::

::: part
::: section
### Constraints

-   `S`{.variable} consists of uppercase English letters, lowercase
    English letters, and `?`.
-   The length of `S`{.variable} is between `4`{.variable} and
    `3\times 10^5`{.variable}, inclusive.
:::
:::

------------------------------------------------------------------------

::: io-style
::: part
::: section
### Input

The input is given from Standard Input in the following format:

    S
:::
:::

::: part
::: section
### Output

Print the answer.
:::
:::
:::

------------------------------------------------------------------------

::: part
::: section
### Sample Input 1

    DD??S
:::
:::

::: part
::: section
### Sample Output 1

    676

When at least one of the `?`s is replaced with a lowercase English
letter, the resulting string will contain a `DDoS`-type string as a
subsequence.
:::
:::

------------------------------------------------------------------------

::: part
::: section
### Sample Input 2

    ????????????????????????????????????????
:::
:::

::: part
::: section
### Sample Output 2

    858572093

Find the count modulo `998244353`{.variable}.
:::
:::

------------------------------------------------------------------------

::: part
::: section
### Sample Input 3

    ?D??S
:::
:::

::: part
::: section
### Sample Output 3

    136604
:::
:::
:::
:::
:::

------------------------------------------------------------------------

::: {.a2a_kit .a2a_kit_size_20 .a2a_default_style .pull-right a2a-url="https://atcoder.jp/contests/abc301/tasks/abc301_f?lang=en" a2a-title="F - Anti-DDoS"}
[]{.a2a_button_facebook} []{.a2a_button_twitter}
[]{.a2a_button_telegram} [](https://www.addtoany.com/share){.a2a_dd}
:::
:::

------------------------------------------------------------------------
:::

::: {.container style="margin-bottom: 80px;"}
-   [Rule](/contests/abc301/rules)
-   [Glossary](/contests/abc301/glossary)

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
