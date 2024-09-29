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
        ll b, c, d; cin >> b >> c >> d;

        // b c
        // 0 0 -> {0, 1}
        // 0 1 -> {0, 0}
        // 1 0 -> {1, 1}
        // 1 1 -> {1, 0}

        vector<int> bits;
        bool impossible = false;
        while (b > 0 || c > 0 || d > 0) {
            int b0 = b % 2;
            int c0 = c % 2;
            int d0 = d % 2;

            if (b0 == 0 && c0 == 0 && d0 == 0) {
                bits.push_back(0);
            } else if (b0 == 0 && c0 == 0 && d0 == 1) {
                bits.push_back(1);
            } else if (b0 == 0 && c0 == 1 && d0 == 0) {
                bits.push_back(0);
            } else if (b0 == 0 && c0 == 1 && d0 == 1) {
                impossible = true;
            } else if (b0 == 1 && c0 == 0 && d0 == 0) {
                impossible = true;
            } else if (b0 == 1 && c0 == 0 && d0 == 1) {
                bits.push_back(0);
            } else if (b0 == 1 && c0 == 1 && d0 == 0) {
                bits.push_back(1);
            } else if (b0 == 1 && c0 == 1 && d0 == 1) {
                bits.push_back(0);
            }
            if (impossible) {
                break;
            }
            b /= 2, c /= 2, d /= 2;
        }

        if (impossible) {
            cout << -1 << '\n';
            continue;
        }

        ll num = 0;
        for (int i = bits.size()-1; i >= 0; i--) {
            num = num * 2 + bits[i];
        }
        cout << num << '\n';
    }
    return 0;
}