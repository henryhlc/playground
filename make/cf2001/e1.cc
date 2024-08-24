#include <iostream>
#include <vector>

using namespace std;

int main() {
    int tt; cin >> tt;
    for (int t = 0; t < tt; t++) {
        int n, k, p; cin >> n >> k >> p;

        // conditions
        // - added to k nodes
        // - each node value is the number of times it is added in the subtree
        // - first pop must be deterministic

        // number of operations 0:k, max level 1:n
        vector<vector<long long>> deterministic(k+1, vector<long long>(n+1, 0));
        vector<vector<long long>> all(k+1, vector<long long>(n+1, 0));

        for (int i = 0; i <= k; i++) {
            all[i][1] = 1;  // 1 level in the tree
            deterministic[i][1] = 1;
        }

        for (int j = 2; j <= n; j++) {
        // number of levels (>=2) in the tree

            // number of add in lower levels, 0 to k
            vector<long long> toChildrenDeterministic(k+1,0);
            vector<long long> toChildrenAll(k+1,0);

            toChildrenDeterministic[0] = 0;
            toChildrenAll[0] = 1;

            for (int t = 1; t <= k; t++) {
            // number of add in lower levels >=1
                for (int l = 0; l <= t; l++) {
                // number of elements given to the left
                    int r = t - l;
                    toChildrenAll[t] += all[l][j-1] * all[r][j-1];
                    toChildrenAll[t] %= p;
                    if (l < r) {
                        toChildrenDeterministic[t] += all[l][j-1] * deterministic[r][j-1];
                        toChildrenDeterministic[t] %= p;
                    } else if (r < l) {
                        toChildrenDeterministic[t] += deterministic[l][j-1] * all[r][j-1];
                        toChildrenDeterministic[t] %= p;
                    }
                }
            }

            all[0][j] = 1; // 0 operations
            deterministic[0][j] = 1;
            for (int i = 1; i <= k; i++) {
            // number of operations (>=1)

                for (int h = 0; h <= i; h++) {
                // number of adds applied to root 
                    int c = i - h;  // number of adds applied to children
                    deterministic[i][j] += toChildrenDeterministic[c];
                    deterministic[i][j] %= p;
                    all[i][j] += toChildrenAll[c];
                    all[i][j] %= p;
                }

            }
        }
        cout << deterministic[k][n] << "\n";
    }
    return 0;
}