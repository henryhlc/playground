#include<algorithm>
#include<iostream>
#include<map>
#include<set>
#include<vector>

using namespace std;
using ll = unsigned long long;

int main() {
    int tt; cin >> tt;
    for (int t = 0; t < tt; t++) {
        ll n, q; cin >> n >> q;
        vector<ll> xs(n);
        map<ll,ll> kToCount;
        for (ll i = 0; i < n; i++) {
            cin >> xs[i];

            if (i > 0) {
                // xs[i-1]+1 to xs[i]-1
                // ...xs[i-1] (...) xs[i] ...
                ll l = xs[i] - xs[i-1] - 1;
                ll count = i * (n - i);
                if (kToCount.find(count) == kToCount.end()) {
                    kToCount[count] = l;
                } else {
                    kToCount[count] += l;
                }
            }

            // xs[i]
            // 0...i-1 i i+1...n-1
            ll count = i * (n-1-i) + i + (n-1-i);
            if (kToCount.find(count) == kToCount.end()) {
                kToCount[count] = 1;
            } else {
                kToCount[count] += 1;
            }
        }

        bool first = true;
        for (int i = 0; i < q; i++) {
            ll query; cin >> query;
            if (!first) {
                cout << ' ';
            }
            first = false;
            if (kToCount.find(query) == kToCount.end()) {
                cout << 0;
            } else {
                cout << kToCount[query];
            }
        }
        cout << '\n';
    }
    return 0;
}