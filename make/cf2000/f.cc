#include <iostream>
#include <algorithm>
#include <vector>
#include <tuple>

using namespace std;

int main() {
    int tt; cin >> tt;
    for (int t = 0; t < tt; t++) {
        int n, k; cin >> n >> k;
        vector<tuple<int, int>> squares(n);
        for (int i = 0; i < n; i++) {
            int a, b; cin >> a >> b;
            squares[i] = make_tuple(a, b);
        }

        // Operations needed to use first i squares to reach score at least k
        vector<vector<int>> dp(n, vector<int>(k+1, -1));
        for (int i = 0; i < n; i++) {
            dp[i][0] = 0;
            vector<int> costs (k+1, -1);
            costs[0] = 0;
            auto [w, h] = squares[i];
            for (int j = 1; j <= k; j++) {
                if (w == 1 && h == 1) {
                    costs[j] = costs[j-1] + 1;
                    if (j+1 <= k) {
                        costs[j+1] = costs[j];
                    }
                    break;
                }
                costs[j] = costs[j-1] + min(w, h);
                if (w == min(w,h)) {
                    h--;
                } else {
                    w--;
                }
            }

            for (int j = 1; j <= k; j++) {
                dp[i][j] = costs[j];
                if (i == 0) {
                    continue;
                }
                for (int s = 0; s <= j; s++) {
                    if (costs[s] < 0 || dp[i-1][j-s] < 0) {
                        continue;
                    }
                    if (dp[i][j] == -1) {
                        dp[i][j] = costs[s] + dp[i-1][j-s];
                    } else {
                        dp[i][j] = min(dp[i][j], costs[s] + dp[i-1][j-s]);
                    }
                }
            }
        }

        if (dp[n-1][k] < 0) {
            cout << -1 << "\n";
        } else {
            cout << dp[n-1][k] << "\n";
        }
    }
}