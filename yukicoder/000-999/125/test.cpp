#include "bits/stdc++.h"
using namespace std;

long long mod = (long long)1e9 + 7;

int gcd(int a, int b){
	if (b == 0) return a;
	return gcd(b, a%b);
}

long long powmod(long long a, int p){
	if (p == 0) return 1;
	if (p % 2 == 1){
		return (powmod(a, p - 1) * a) % mod;
	}
	long long mid = powmod(a, p / 2);
	return (mid * mid) % mod;
}

long rev(long long a){
	return powmod(a, mod - 2);
}

int main() {
	int K;
	cin >> K;
	vector<int> C(K);
	int sum = 0;
	for (int i = 0; i < K; i++)
	{
		cin >> C[i];
		sum += C[i];
	}
	int g = C[0];
	for (int i = 0; i < K; i++)
	{
		g = gcd(g, C[i]);
	}

	long long ans = 0;

	vector<int> factor;
	vector<long long> num;

	for (int i = g; i >= 1; i--)
	{
		if (g%i != 0) continue;
		int remain = sum / i - 1;
    
        cout << "remain " << remain << " ";

		long long mul = 1;
		long long div = 1;
		for (int j = 0; j < K; j++)
		{
			int end = 0;
			if (j == 0) end = 1;
			for (int l = 0; l < C[j] / i - end; l++)
			{
				mul *= remain--;
				mul %= mod;
				div *= l + 1;
				div %= mod;
			}
		}
		mul *= rev(div);

		for (int j = 0; j < factor.size(); j++)
		{
			if (factor[j] % i != 0) continue;
			mul -= num[j];
		}
		mul %= mod;
		mul += mod;
		mul %= mod;

		num.push_back(mul);
        cout << "num " <<  mul << " ";

		mul *= rev(C[0] / i);
		mul %= mod;
		factor.push_back(i);
		ans += mul;
		ans %= mod;
        cout << "factor " << i  <<" ";
        cout << ans << endl;

	}

	cout << ans << endl;
}