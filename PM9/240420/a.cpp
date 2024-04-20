#include <iostream>
using namespace std;

int main()
{
    int x, y;
    cin >> x >> y;

    int ans = 0;
    int xx = 0, yy = 0;

    while (yy < y)
    {
        if (xx < x)
        {
            xx++;
        }
        else
        {
            xx--;
        }
        yy += 2;
        ans++;
    }

    if (xx == x && yy == y)
    {
        cout << ans << endl;
    }
    else
    {
        cout << -1 << endl;
    }
}