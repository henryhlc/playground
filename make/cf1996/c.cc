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
        int n, q; cin >> n >> q;
        string a, b; cin >> a >> b;

        vector<vector<int>> prefixCountA(a.length(), vector<int>(26));
        vector<vector<int>> prefixCountB(b.length(), vector<int>(26));
        prefixCountA[0] = vector<int>(26, 0);
        prefixCountA[0][a[0]-'a']++;
        prefixCountB[0] = vector<int>(26, 0);
        prefixCountB[0][b[0]-'a']++;
        
        for (int i = 1; i < a.length(); i++) {
            prefixCountA[i] = prefixCountA[i-1];
            prefixCountA[i][a[i]-'a']++;
            prefixCountB[i] = prefixCountB[i-1];
            prefixCountB[i][b[i]-'a']++;
        }

        for (int i = 0; i < q; i++) {
            int l, r; cin >> l >> r;
            l--, r--;
            vector<int> countA(26);
            vector<int> countB(26);
            for (int c = 0; c < 26; c++) {
                countA[c] = prefixCountA[r][c];
                countB[c] = prefixCountB[r][c];
            }
            if (l > 0) {
                for (int c = 0; c < 26; c++) {
                    countA[c] -= prefixCountA[l-1][c];
                    countB[c] -= prefixCountB[l-1][c];
                }
            }
            int same = 0;
            for (int c = 0; c < 26; c++) {
                same += min(countA[c], countB[c]);
            }
            cout << r - l + 1 - same << '\n';
        }
    }
    return 0;
}