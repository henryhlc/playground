#include<algorithm>
#include<iostream>
#include<map>
#include<set>
#include<vector>
#include <functional>

using namespace std;
using ll = long long;

int main() {
    int tt; cin >> tt;
    for (int t = 0; t < tt; t++) {
        int n; cin >> n;
        vector<int> as(n);
        for (int i = 0; i < n; i++) {
            cin >> as[i];
        }
        vector<vector<int>> children(n, vector<int>{});
        for (int i = 0; i < n-1; i++) {
            int p; cin >> p; p--;
            children[p].push_back(i+1);
        }

        function<int(int)> maxmin = [&](int root) -> int {
            if (children[root].empty()) {
                return as[root];
            }
            int childrenMaxmin = -1;
            for (auto c : children[root]) {
                if (childrenMaxmin == -1) {
                    childrenMaxmin = maxmin(c);
                } else {
                    childrenMaxmin = min(maxmin(c), childrenMaxmin);
                }
            }
            if (childrenMaxmin <= as[root]) {
                return childrenMaxmin;
            } else {
                int diff = childrenMaxmin - as[root] - 1;
                return as[root] + (diff + 1) / 2;
            }
        };

        int minc = -1;
        for (auto c : children[0]) {
            if (minc == -1) {
                minc = maxmin(c);
            } else {
                minc = min(minc, maxmin(c));
            }
        }

        cout << as[0] + minc << '\n';
    }
    return 0;
}