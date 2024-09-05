#include <iostream>
#include <vector>
#include <deque>
#include <utility>

using namespace std;
using pii = pair<int,int>;

int main() {
    int tt; cin >> tt;
    for (int t = 0; t < tt; t++) {
        int n, m; cin >> n >> m;

        vector<int> maxHopDest(n,-1);
        vector<vector<int>> hopFrom(n, vector<int>{});
        for (int i = 0; i < m; i++) {
            int u, v; cin >> u >> v;
            u--; v--;
            if (v > maxHopDest[u]) {
                maxHopDest[u] = v;
            }
            hopFrom[v].push_back(u);
        }

        vector<int> earliest(n, -1);
        earliest[0] = 0;
        for (int i = 1; i < n; i++) {
            earliest[i] = earliest[i-1] + 1;
            for (auto from : hopFrom[i]) {
                earliest[i] = min(earliest[i], earliest[from] + 1);
            }
        }

        deque<pii> intervals;
        for (int i = 0; i < n; i++) {
            if (maxHopDest[i] == -1) {
                continue;
            }
            int maxS = maxHopDest[i] - earliest[i] - 2;
            int minS = i + 1;
            if (minS <= maxS) {
                intervals.push_back({minS, maxS});
            }
        }

        int maxEnd = -1;
        for (int s = 0; s < n - 1; s++) {
            while (!intervals.empty() && intervals.front().first <= s) {
                maxEnd = max(maxEnd, intervals.front().second);
                intervals.pop_front();
            }
            if (maxEnd >= s) {
                cout << '0';
            } else {
                cout << '1';
            }
        }
        cout << '\n';
    }

    return 0;
}