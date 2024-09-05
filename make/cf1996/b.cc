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
        int n, k; cin >> n >> k;
        vector<vector<char>> out(n/k, vector<char>(n/k));
        for (int i = 0; i < n / k; i++) {
            string row;
            for (int j = 0; j < k; j++) {
                cin >> row;
            }
            for (int j = 0; j < n / k; j++) {
                out[i][j] = row[j * k];
            }
        }
        for (auto& row : out) {
            for (auto v : row) {
                cout << v;
            }
            cout << '\n';
        }
    }
    return 0;
}