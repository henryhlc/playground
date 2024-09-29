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
        ll k; cin >> k;

        // square numbers are off, others are on

        // x such that x^2 - x >= k
        ll low = 0;
        ll high = 1e9 + 1;
        while (high - low > 1) {
            ll mid = (low + high) / 2;
            if (mid * mid - mid >= k) {
                high = mid;
            } else {
                low = mid;
            }
        }

        ll n = high * high - 1; 
        ll count = high * high - high;
        cout << n - (count - k) << '\n';
    }
    return 0;
}