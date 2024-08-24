#include <iostream>
#include <vector>
#include <algorithm>
#include <map>

using namespace std;

int main() {
    // How many k-by-k squares contains i,j?
    // max, k^2, when i >= k-1; i <= n-1-(k-1); j >= k-1; j <= m-1-(k-1)
    // for i, j
    int tt; cin >> tt;
    for (int t = 0; t < tt; t++) {
        int n,m,k,w; cin >> n >> m >> k >> w;
        vector<long long> heights(w);
        for (int i = 0; i < w; i++) {
            cin >> heights[i];
        }
        sort(heights.begin(), heights.end());

        vector<long long> counts(m*n);
        for (int i = 0; i < n; i++) {
            int r_low = max(i - (k-1), 0);
            int r_high = min(i + (k-1), n-1);
            long long r_count = r_high - r_low + 1 - k + 1;
            for (int j = 0; j < m; j++) {
                int c_low = max(j - (k-1), 0);
                int c_high = min(j + (k-1), m-1);
                long long c_count = c_high - c_low + 1 - k + 1;
                counts[i * m + j] = r_count * c_count;
            }
        }

        sort(counts.begin(), counts.end());

        long long score = 0;
        for (int i = 0; i < w; i++) {
            score += heights[w-1-i] * counts[counts.size()-1-i];
        }
        cout << score << "\n";
    }
    return 0;
}