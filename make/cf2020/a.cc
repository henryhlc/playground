#include<algorithm>
#include<iostream>
#include<map>
#include<set>
#include<vector>

using namespace std;
using ll = long long;

int main() {
    int tt; cin >> tt;
    for (int t = 0; t < tt; t++) {
        ll n, k; cin >> n >> k; 
       
        if (k == 1) {
            cout << n << '\n';
            continue;
        }

        ll p = 1;
        while (p * k <= n) {
            p *= k;
        }

        int ops = 0;
        while (n > 0) {
            while (n < p) {
                p /= k;
            }
            ops += n / p;
            n %= p;
        }
        cout << ops << '\n';
    }
    return 0;
}