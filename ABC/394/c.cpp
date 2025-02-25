#include <bits/stdc++.h>
#include <atcoder/all>

using namespace std;
using ll = long long;
using P = pair<int, int>;

#define rep(i, n) for (int i = 0; i < (int)(n); i++)

int main()
{
    string s;

    cin >> s;

    string t = "";
    rep(i, s.size())
    {
        t += s[i];
        int c = 0;
        while (t.size() >= 2 && t[t.size() - 2] == 'W' && t[t.size() - 1] == 'A')
        {
            t.pop_back();
            t.pop_back();
            t += 'A';
            c++;
        }
        rep(j, c)
        {
            t += 'C';
        }
    }

    cout << t << endl;
    return 0;
}
