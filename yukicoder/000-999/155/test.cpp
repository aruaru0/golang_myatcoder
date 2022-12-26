#include <iostream>
#include <vector>
#include <algorithm>
#include <cstdio>
#define repeat(i,n) for (int i = 0; (i) < (n); ++(i))
#define repeat_from(i,m,n) for (int i = (m); (i) < (n); ++(i))
#define repeat_reverse(i,n) for (int i = (n)-1; (i) >= 0; --(i))
typedef long long ll;
using namespace std;
int main() {
    int n, l; cin >> n >> l;
    l *= 60;
    vector<int> s(n);
    repeat (i,n) {
        int mn, sc; char dummy; cin >> mn >> dummy >> sc;
        s[i] = mn * 60 + sc;
    }
    int max_s = *max_element(s.begin(), s.end());
    vector<vector<ll> > dp(n+1, vector<ll>(l+max_s));
    dp[0][0] = 1;
    repeat (i,n) {
        repeat_reverse (j,n) {
            repeat_reverse (k,l) {
                dp[j+1][k+s[i]] += dp[j][k];
            }
        }
    }
    double acc = 0;
    vector<double> fact(n+1); fact[0] = 1; repeat (i,n) fact[i+1] = (i+1) * fact[i];
    repeat (i,n) {
        vector<vector<ll> > prv = dp;
        repeat (j,n) {
            repeat (k,l) {
                prv[j+1][k+s[i]] -= prv[j][k];
            }
        }
        repeat (j,n) {
            repeat_from (k,max(0,l-s[i]),l) {
                acc += (j+1) * fact[j] * prv[j][k] * fact[n-j-1];
            }
        }
        cout << acc <<endl;
    }
    printf("%.12lf\n", acc / fact[n]);
    return 0;
}