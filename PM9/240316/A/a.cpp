#include <iostream>

using namespace std;

int main() {
    int n;

    cin >> n;

    int ans = 1e8;
    for(int i=0;i < n;i++) {
        int cnt = 0;
        int p;
        cin >> p;
        while(p%10 == 0) {
            p /= 10;
            cnt++;
        }
        ans = min(ans, cnt);
    }

    cout << ans << endl;
}