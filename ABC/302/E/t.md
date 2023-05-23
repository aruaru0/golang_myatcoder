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
        日本語](/contests/abc302/tasks/abc302_e?lang=ja)
    -   [![](//img.atcoder.jp/assets/top/img/flag-lang/en.png)
        English](/contests/abc302/tasks/abc302_e?lang=en)
-   [Sign
    Up](/register?continue=https%3A%2F%2Fatcoder.jp%2Fcontests%2Fabc302%2Ftasks%2Fabc302_e)
-   [Sign
    In](/login?continue=https%3A%2F%2Fatcoder.jp%2Fcontests%2Fabc302%2Ftasks%2Fabc302_e)
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
[ E - Isolation
[Editorial](/contests/abc302/tasks/abc302_e/editorial){.btn .btn-default
.btn-sm} ]{.h2}
[[![](//img.atcoder.jp/assets/top/img/flag-lang/ja.png)]{data-lang="ja"}
/
[![](//img.atcoder.jp/assets/top/img/flag-lang/en.png)]{data-lang="en"}]{#task-lang-btn
.pull-right}

------------------------------------------------------------------------

Time Limit: 2 sec / Memory Limit: 1024 MB

::: {#task-statement}
[ [ ]{.lang-ja}]{.lang}

配点 : `425`{.variable} 点

::: part
::: section
### 問題文

最初 `N`{.variable} 頂点 `0`{.variable} 辺の無向グラフがあり、各頂点には
`1`{.variable} から `N`{.variable} まで番号がついています。\
`Q`{.variable}
個のクエリが与えられるので、順に処理し、各クエリの後における「他のどの頂点とも辺で結ばれていない頂点」の数を出力してください。

`i`{.variable} 個目のクエリは `\mathrm{query}_i`{.variable}
であり、各クエリは次の `2`{.variable} 種類いずれかです。

-   `1 u v`: 頂点 `u`{.variable} と頂点 `v`{.variable}
    を辺で結ぶ。このクエリが与えられる直前の時点で、頂点 `u`{.variable}
    と頂点 `v`{.variable} は辺で結ばれていない事が保証される。

-   `2 v` : 頂点 `v`{.variable}
    と他の頂点を結ぶ辺をすべて削除する。（頂点 `v`{.variable}
    自体は削除しない。）
:::
:::

::: part
::: section
### 制約

-   `2 \leq N\leq 3\times 10^5`{.variable}
-   `1 \leq Q\leq 3\times 10^5`{.variable}
-   `1`{.variable}
    番目の種類のクエリにおいて、`1\leq u,v\leq N`{.variable},
    `u\neq v`{.variable}
-   `2`{.variable}
    番目の種類のクエリにおいて、`1\leq v\leq N`{.variable}
-   `1`{.variable} 番目の種類のクエリの直前の時点で、そのクエリの
    `u,v`{.variable} について頂点 `u`{.variable} と頂点 `v`{.variable}
    は辺で結ばれていない。
-   入力はすべて整数
:::
:::

------------------------------------------------------------------------

::: io-style
::: part
::: section
### 入力

入力は以下の形式で標準入力から与えられる。

    N Q
    \mathrm{query}_1
    \mathrm{query}_2
    \vdots
    \mathrm{query}_Q
:::
:::

::: part
::: section
### 出力

`Q`{.variable} 行出力せよ。\
`i`{.variable} 行目 `(1\leq i\leq Q)`{.variable} には、`i`{.variable}
個目のクエリを処理した後の「他のどの頂点とも辺で結ばれていない頂点」の数を出力せよ。
:::
:::
:::

------------------------------------------------------------------------

::: part
::: section
### 入力例 1

    3 7
    1 1 2
    1 1 3
    1 2 3
    2 1
    1 1 2
    2 2
    1 1 2
:::
:::

::: part
::: section
### 出力例 1

    1
    0
    0
    1
    0
    3
    1

`1`{.variable} 個目のクエリの後で、頂点 `1`{.variable} と頂点
`2`{.variable} は互いに結ばれており、頂点 `3`{.variable}
のみが他のどの頂点とも辺で結ばれていません。\
よって、`1`{.variable} 行目には `1`{.variable} を出力します。

また、`3`{.variable} 個目のクエリの後でどの相異なる `2`{.variable}
頂点の間も辺で結ばれていますが、`4`{.variable} 個目のクエリによって、
頂点 `1`{.variable} と他の頂点を結ぶ辺、すなわち 頂点 `1`{.variable}
と頂点 `2`{.variable} を結ぶ辺および頂点 `1`{.variable} と頂点
`3`{.variable} を結ぶ辺が削除されます。 この結果として、頂点
`2`{.variable} と頂点 `3`{.variable} は互いに結ばれているが、頂点
`1`{.variable} は他のどの頂点とも辺で結ばれていない状態となります。\
よって、`3`{.variable} 行目には `0`{.variable} を、`4`{.variable}
行目には `1`{.variable} を出力します。
:::
:::

------------------------------------------------------------------------

::: part
::: section
### 入力例 2

    2 1
    2 1
:::
:::

::: part
::: section
### 出力例 2

    2

`2`{.variable}
番目の種類のクエリを行う直前の時点で、すでにその頂点と他の頂点を結ぶ辺が
`1`{.variable} 本も存在しないこともあります。
:::
:::

[ ]{.lang-en}

Score : `425`{.variable} points

::: part
::: section
### Problem Statement

There is an undirected graph with `N`{.variable} vertices numbered
`1`{.variable} through `N`{.variable}, and initially with `0`{.variable}
edges.\
Given `Q`{.variable} queries, process them in order. After processing
each query, print the number of vertices that are not connected to any
other vertices by an edge.

The `i`{.variable}-th query, `\mathrm{query}_i`{.variable}, is of one of
the following two kinds.

-   `1 u v`: connect vertex `u`{.variable} and vertex `v`{.variable}
    with an edge. It is guaranteed that, when this query is given,
    vertex `u`{.variable} and vertex `v`{.variable} are not connected by
    an edge.

-   `2 v`: remove all edges that connect vertex `v`{.variable} and the
    other vertices. (Vertex `v`{.variable} itself is not removed.)
:::
:::

::: part
::: section
### Constraints

-   `2 \leq N\leq 3\times 10^5`{.variable}
-   `1 \leq Q\leq 3\times 10^5`{.variable}
-   For each query of the first kind, `1\leq u,v\leq N`{.variable} and
    `u\neq v`{.variable}.
-   For each query of the second kind, `1\leq v\leq N`{.variable}.
-   Right before a query of the first kind is given, there is no edge
    between vertices `u`{.variable} and `v`{.variable}.
-   All values in the input are integers.
:::
:::

------------------------------------------------------------------------

::: io-style
::: part
::: section
### Input

The input is given from Standard Input in the following format:

    N Q
    \mathrm{query}_1
    \mathrm{query}_2
    \vdots
    \mathrm{query}_Q
:::
:::

::: part
::: section
### Output

Print `Q`{.variable} lines.\
The `i`{.variable}-th line `(1\leq i\leq Q)`{.variable} should contain
the number of vertices that are not connected to any other vertices by
an edge.
:::
:::
:::

------------------------------------------------------------------------

::: part
::: section
### Sample Input 1

    3 7
    1 1 2
    1 1 3
    1 2 3
    2 1
    1 1 2
    2 2
    1 1 2
:::
:::

::: part
::: section
### Sample Output 1

    1
    0
    0
    1
    0
    3
    1

After the first query, vertex `1`{.variable} and vertex `2`{.variable}
are connected to each other by an edge, but vertex `3`{.variable} is not
connected to any other vertices.\
Thus, `1`{.variable} should be printed in the first line.

After the third query, all pairs of different vertices are connected by
an edge.\
However, the fourth query asks to remove all edges that connect vertex
`1`{.variable} and the other vertices, specifically to remove the edge
between vertex `1`{.variable} and vertex `2`{.variable}, and another
between vertex `1`{.variable} and vertex `3`{.variable}. As a result,
vertex `2`{.variable} and vertex `3`{.variable} are connected to each
other, while vertex `1`{.variable} is not connected to any other
vertices by an edge.\
Thus, `0`{.variable} and `1`{.variable} should be printed in the third
and fourth lines, respectively.
:::
:::

------------------------------------------------------------------------

::: part
::: section
### Sample Input 2

    2 1
    2 1
:::
:::

::: part
::: section
### Sample Output 2

    2

When the query of the second kind is given, there may be no edge that
connects that vertex and the other vertices.
:::
:::
:::
:::
:::

------------------------------------------------------------------------

::: {.a2a_kit .a2a_kit_size_20 .a2a_default_style .pull-right a2a-url="https://atcoder.jp/contests/abc302/tasks/abc302_e?lang=en" a2a-title="E - Isolation"}
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
