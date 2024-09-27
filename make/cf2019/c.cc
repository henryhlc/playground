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
        long long n, k; cin >> n >> k; 
        vector<long long> as(n);
        for (int i = 0; i < n; i++) {
            cin >> as[i];
        }

        long long total = 0;
        long long max_a = 0;
        for (auto a : as) {
            total += a;
            max_a = max(max_a, a);
        }

        for (long long h = n; h >= 1; h--) {
            long long n_deck = max(max_a, (total + h - 1) / h);
            if (n_deck * h - total > k) {
                continue;
            } else {
                cout << h << '\n';
                break;
            }
        }
    }
    return 0;
}