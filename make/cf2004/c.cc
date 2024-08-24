#include <iostream>
#include <vector>
#include <algorithm>

using namespace std;

int main() {
    int tt; cin >> tt;
    for (int t = 0; t < tt; ++t) {
        int n, k; cin >> n >> k;
        vector<int> as(n);
        for (int i = 0; i < n; i++) {
            cin >> as[i];
        }
        // total increase <= k
        // both play optimally, they will take in descending order
        // cost is a(n) - a(n-1) + a(n-2) - ...
        sort(as.begin(), as.end());
        int reducable_cost = 0;
        int irreducable_cost = 0;
        for (int i = as.size()-1; i >= 0; i -= 2) {
            if (i > 0) {
                reducable_cost += as[i] - as[i-1];
            } else {
                irreducable_cost += as[i];
            }
        }

        reducable_cost -= min(reducable_cost, k);
        cout << reducable_cost + irreducable_cost << "\n";

    }
    return 0;
}