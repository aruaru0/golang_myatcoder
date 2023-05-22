::: {#modal-contest-start .modal .fade tabindex="-1" role="dialog"}
::: {.modal-dialog role="document"}
::: modal-content
::: modal-header
[×]{aria-hidden="true"}

#### Contest started {#contest-started .modal-title}
:::

::: modal-body
TOYOTA MOTOR CORPORATION Programming Contest 2023#2 (AtCoder Beginner
Contest 302) has begun.
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
TOYOTA MOTOR CORPORATION Programming Contest 2023#2 (AtCoder Beginner
Contest 302) has ended.
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
-   [TOYOTA MOTOR CORPORATION Programming Contest 2023#2 (AtCoder
    Beginner Contest 302)](/contests/abc302){.contest-title}

```{=html}
<!-- -->
```
-   [![](//img.atcoder.jp/assets/top/img/flag-lang/en.png) English
    []{.caret}](#){.dropdown-toggle toggle="dropdown" role="button"
    aria-haspopup="true" aria-expanded="false"}
    -   [![](//img.atcoder.jp/assets/top/img/flag-lang/ja.png)
        日本語](/contests/abc302/tasks/abc302_d?lang=ja)
    -   [![](//img.atcoder.jp/assets/top/img/flag-lang/en.png)
        English](/contests/abc302/tasks/abc302_d?lang=en)
-   [Sign
    Up](/register?continue=https%3A%2F%2Fatcoder.jp%2Fcontests%2Fabc302%2Ftasks%2Fabc302_d)
-   [Sign
    In](/login?continue=https%3A%2F%2Fatcoder.jp%2Fcontests%2Fabc302%2Ftasks%2Fabc302_d)
:::
:::

::: {#main-container .container style="padding-top:50px;"}
::: row
::: {#contest-nav-tabs .col-sm-12 .mb-2 .cnvtb-fixed}
<div>

[Contest Duration: [2023-05-20
21:00:00+0900](http://www.timeanddate.com/worldclock/fixedtime.html?iso=20230520T2100&p1=248){target="blank"} -
[2023-05-20
22:40:00+0900](http://www.timeanddate.com/worldclock/fixedtime.html?iso=20230520T2240&p1=248){target="blank"}
(local time) (100 minutes)]{.small} [[Back to Home](/home)]{.small}

</div>

-   [[]{.glyphicon .glyphicon-home aria-hidden="true"}
    Top](/contests/abc302)
-   [[]{.glyphicon .glyphicon-tasks aria-hidden="true"}
    Tasks](/contests/abc302/tasks)
-   [[]{.glyphicon .glyphicon-question-sign aria-hidden="true"}
    Clarifications []{#clar-badge
    .badge}](/contests/abc302/clarifications)
-   [[]{.glyphicon .glyphicon-list aria-hidden="true"}
    Results[]{.caret}](#){.dropdown-toggle toggle="dropdown"
    role="button" aria-haspopup="true" aria-expanded="false"}
    -   [[]{.glyphicon .glyphicon-globe aria-hidden="true"} All
        Submissions](/contests/abc302/submissions)
-   [[]{.glyphicon .glyphicon-sort-by-attributes-alt aria-hidden="true"}
    Standings](/contests/abc302/standings)
-   [[]{.glyphicon .glyphicon-sort-by-attributes-alt aria-hidden="true"}
    Virtual Standings](/contests/abc302/standings/virtual)
-   [[]{.glyphicon .glyphicon-book aria-hidden="true"}
    Editorial](/contests/abc302/editorial)
-   [[]{.glyphicon .glyphicon-pushpin
    aria-hidden="true"}](javascript:void(0)){#fix-cnvtb}
:::

::: col-sm-12
[ D - Impartial Gift
[Editorial](/contests/abc302/tasks/abc302_d/editorial){.btn .btn-default
.btn-sm} ]{.h2}
[[![](//img.atcoder.jp/assets/top/img/flag-lang/ja.png)]{data-lang="ja"}
/
[![](//img.atcoder.jp/assets/top/img/flag-lang/en.png)]{data-lang="en"}]{#task-lang-btn
.pull-right}

------------------------------------------------------------------------

Time Limit: 2 sec / Memory Limit: 1024 MB

::: {#task-statement}
[ [ ]{.lang-ja}]{.lang}

配点 : `400`{.variable} 点

::: part
::: section
### 問題文

高橋君は青木君とすぬけ君に **`1`{.variable}
つずつ**贈り物を送ることにしました。\
青木君への贈り物の候補は `N`{.variable} 個あり、 それぞれの価値は
`A_1, A_2, \ldots,A_N`{.variable} です。\
すぬけ君への贈り物の候補は `M`{.variable} 個あり、 それぞれの価値は
`B_1, B_2, \ldots,B_M`{.variable} です。

高橋君は `2`{.variable} 人への贈り物の価値の差が `D`{.variable}
以下になるようにしたいと考えています。

条件をみたすように贈り物を選ぶことが可能か判定し、可能な場合はそのような選び方における贈り物の価値の和の最大値を求めてください。
:::
:::

::: part
::: section
### 制約

-   `1\leq N,M\leq 2\times 10^5`{.variable}
-   `1\leq A_i,B_i\leq 10^{18}`{.variable}
-   `0\leq D \leq 10^{18}`{.variable}
-   入力はすべて整数
:::
:::

------------------------------------------------------------------------

::: io-style
::: part
::: section
### 入力

入力は以下の形式で標準入力から与えられる。

    N M D
    A_1 A_2 \ldots A_N
    B_1 B_2 \ldots B_M
:::
:::

::: part
::: section
### 出力

高橋君が条件をみたすように贈り物を選ぶことができる場合、
条件をみたし、かつ価値の和が最大になるように贈り物を選んだ時の価値の和を出力せよ。
高橋君が条件をみたすように選ぶことができない場合、`-1`{.variable}
を出力せよ。
:::
:::
:::

------------------------------------------------------------------------

::: part
::: section
### 入力例 1

    2 3 2
    3 10
    2 5 15
:::
:::

::: part
::: section
### 出力例 1

    8

高橋君は贈り物の価値の差を `2`{.variable} 以下にする必要があります。\
青木君に価値 `3`{.variable}, すぬけ君に価値 `5`{.variable}
の贈り物を渡すと条件をみたし、価値の和としてはこのときが最大となります。\
よって、`3+5=8`{.variable} を出力します。
:::
:::

------------------------------------------------------------------------

::: part
::: section
### 入力例 2

    3 3 0
    1 3 3
    6 2 7
:::
:::

::: part
::: section
### 出力例 2

    -1

条件をみたすように贈り物を選ぶことは不可能です。
また、同一人物に対して、同じ価値の贈り物が複数存在することもあります。
:::
:::

------------------------------------------------------------------------

::: part
::: section
### 入力例 3

    1 1 1000000000000000000
    1000000000000000000
    1000000000000000000
:::
:::

::: part
::: section
### 出力例 3

    2000000000000000000

答えが `32`{.variable}
bit整数型の範囲に収まらないことがあることに注意してください。
:::
:::

------------------------------------------------------------------------

::: part
::: section
### 入力例 4

    8 6 1
    2 5 6 5 2 1 7 9
    7 2 5 5 2 4
:::
:::

::: part
::: section
### 出力例 4

    14
:::
:::

[ ]{.lang-en}

Score : `400`{.variable} points

::: part
::: section
### Problem Statement

Takahashi has decided to give **one** gift to Aoki and **one** gift to
Snuke.\
There are `N`{.variable} candidates of gifts for Aoki, and their values
are `A_1, A_2, \ldots,A_N`{.variable}.\
There are `M`{.variable} candidates of gifts for Snuke, and their values
are `B_1, B_2, \ldots,B_M`{.variable}.

Takahashi wants to choose gifts so that the difference in values of the
two gifts is at most `D`{.variable}.

Determine if he can choose such a pair of gifts. If he can, print the
maximum sum of values of the chosen gifts.
:::
:::

::: part
::: section
### Constraints

-   `1\leq N,M\leq 2\times 10^5`{.variable}
-   `1\leq A_i,B_i\leq 10^{18}`{.variable}
-   `0\leq D \leq 10^{18}`{.variable}
-   All values in the input are integers.
:::
:::

------------------------------------------------------------------------

::: io-style
::: part
::: section
### Input

The input is given from Standard Input in the following format:

    N M D
    A_1 A_2 \ldots A_N
    B_1 B_2 \ldots B_M
:::
:::

::: part
::: section
### Output

If he can choose gifts to satisfy the condition, print the maximum sum
of values of the chosen gifts. If he cannot satisfy the condition, print
`-1`{.variable}.
:::
:::
:::

------------------------------------------------------------------------

::: part
::: section
### Sample Input 1

    2 3 2
    3 10
    2 5 15
:::
:::

::: part
::: section
### Sample Output 1

    8

The difference of values of the two gifts should be at most
`2`{.variable}.\
If he gives a gift with value `3`{.variable} to Aoki and another with
value `5`{.variable} to Snuke, the condition is satisfied, achieving the
maximum possible sum of values.\
Thus, `3+5=8`{.variable} should be printed.
:::
:::

------------------------------------------------------------------------

::: part
::: section
### Sample Input 2

    3 3 0
    1 3 3
    6 2 7
:::
:::

::: part
::: section
### Sample Output 2

    -1

He cannot choose gifts to satisfy the condition. Note that the
candidates of gifts for a person may contain multiple gifts with the
same value.
:::
:::

------------------------------------------------------------------------

::: part
::: section
### Sample Input 3

    1 1 1000000000000000000
    1000000000000000000
    1000000000000000000
:::
:::

::: part
::: section
### Sample Output 3

    2000000000000000000

Note that the answer may not fit into a `32`{.variable}-bit integer
type.
:::
:::

------------------------------------------------------------------------

::: part
::: section
### Sample Input 4

    8 6 1
    2 5 6 5 2 1 7 9
    7 2 5 5 2 4
:::
:::

::: part
::: section
### Sample Output 4

    14
:::
:::
:::
:::
:::

------------------------------------------------------------------------

::: {.a2a_kit .a2a_kit_size_20 .a2a_default_style .pull-right a2a-url="https://atcoder.jp/contests/abc302/tasks/abc302_d?lang=en" a2a-title="D - Impartial Gift"}
[]{.a2a_button_facebook} []{.a2a_button_twitter}
[]{.a2a_button_telegram} [](https://www.addtoany.com/share){.a2a_dd}
:::
:::

------------------------------------------------------------------------
:::

::: {.container style="margin-bottom: 80px;"}
-   [Rule](/contests/abc302/rules)
-   [Glossary](/contests/abc302/glossary)

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
