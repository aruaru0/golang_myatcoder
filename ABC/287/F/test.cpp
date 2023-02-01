#include <bits/stdc++.h>
using namespace std;
#include <atcoder/all>
using namespace atcoder;
#define rep(i,n) for (int i = 0; i < (n); ++i)
using mint = modint998244353;

int main() {
  int n;
  cin >> n;
  vector<vector<int>> to(n);
  rep(i,n-1) {
    int a, b;
    cin >> a >> b;
    --a; --b;
    to[a].push_back(b);
    to[b].push_back(a);
  }

  using vm = vector<mint>;
  using vvm = vector<vm>;
  auto dfs = [&](auto f, int v, int p=-1) -> vvm {
    vvm dp(2,vm(1,1));
    for (int u : to[v]) {
      if (u == p) continue;
      vvm r = f(f,u,v);
      int rn = r[0].size(), pn = dp[0].size();
      vvm p(2,vm(rn+pn));
      swap(dp,p);
      rep(i,pn)rep(j,rn) {
        dp[1][i+j] += p[1][i]*(r[0][j]+r[1][j]);
        dp[0][i+j] += p[0][i]*r[0][j];
        dp[0][i+j+1] += p[0][i]*r[1][j];
      }
    }
    return dp;
  };

  vvm r = dfs(dfs,0);
  for (int i = 1; i <= n; i++) {
    mint ans;
    if (i < n) ans += r[0][i];
    ans += r[1][i-1];
    cout << ans.val() << '\n';
  }
  return 0;
}
