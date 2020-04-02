#include <bits/stdc++.h>
#define rep(i,n) for (int i = 0; i < (n); ++i)
using namespace std;
typedef long long ll;

ll gcd(ll a, ll b) { return b?gcd(b,a%b):a;}
ll lcm(ll a, ll b) { return a/gcd(a,b)*b;}

int f(int x) {
  int res = 0;
  while (x%2 == 0) {
    x /= 2;
    res++;
  }
  return res;
}

int main() {
  int n, m;
  cin >> n >> m;
  vector<int> a(n);
  rep(i,n) cin >> a[i];

  // a -> a'
  rep(i,n) {
    if (a[i]%2 == 1) {
      cout << 0 << endl;
      return 0;
    }
    a[i] /= 2;
  }

  // a' -> a''
  int t = f(a[0]);
  rep(i,n) {
    if (f(a[i]) != t) {
      cout << 0 << endl;
      return 0;
    }
    a[i] >>= t; // a[i] /= 2^t
  }
  m >>= t;

  ll l = 1;
  rep(i,n) {
    l = lcm(l, a[i]);
    if (l > m) {
      cout << 0 << endl;
      return 0;
    }
  }

  m /= l;
  int ans = (m+1)/2;
  cout << ans << endl;
  return 0;
}
