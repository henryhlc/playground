#include <iostream>
#include <vector>

using namespace std;

using ll = long long;
using pll = pair<ll, ll>;

ll invImpl(ll m, ll a, ll b, pll aVec, pll bVec) {
    ll q = a / b;
    ll r = a % b;

    pll bqVec = pll {bVec.first*q % m, bVec.second*q % m};
    pll rVec = pll {(((aVec.first - bqVec.first) % m) + m) % m, (((aVec.second - bqVec.second) % m) + m) % m};

    if (r == 1) {
        return rVec.second;
    }
    return invImpl(m, b, r, bVec, rVec);
}

ll inv(ll m, ll x) {
    if (x == 1) {
        return 1;
    }
    return invImpl(m, m, x, {1,0}, {0,1});
}


int main() {
    int tt; cin >> tt;
    ll m = 1000000007;
    for (int t = 0; t < tt; t++) {
        int n, k; cin >> n >> k;
        int numOnes = 0;
        int numZeros = 0;
        for (int i = 0; i < n; i++) {
            int e; cin >> e;
            if (e == 0) {
                numZeros++;
            } else {
                numOnes++;
            }
        }

        vector<ll> factorial(n+1);
        factorial[0] = 1;
        for (int i = 1; i < factorial.size(); i++) {
            factorial[i] = (factorial[i-1] * i) % m;
        }

        vector<ll> facInv(n+1);
        for (int i = 0; i < factorial.size(); i++) {
            facInv[i] = inv(m, factorial[i]);
        }

        long long res = 0;
        for (int i = 0; i <= k/2; i++) {
            int chooseZeros = i;
            int chooseOnes = k - i;
            if (numZeros < chooseZeros || numOnes < chooseOnes) {
                continue;
            }

            ll countZeros = factorial[numZeros];
            countZeros *= facInv[numZeros-chooseZeros];
            countZeros %= m;
            countZeros *= facInv[chooseZeros];
            countZeros %= m;

            ll countOnes = factorial[numOnes];
            countOnes *= facInv[numOnes-chooseOnes];
            countOnes %= m;
            countOnes *= facInv[chooseOnes];
            countOnes %= m;
            
            res += (countZeros * countOnes) % m;
            res %= m;
        }
        cout << res << "\n";
    }
    return 0;
}