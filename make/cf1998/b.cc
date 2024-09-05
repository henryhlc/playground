#include <iostream>
#include <vector>
#include <algorithm>

using namespace std;

int main() {
    int tt; cin >> tt;
    for (int t = 0; t < tt; t++) {
        int n; cin >> n;
        vector<int> ps(n);
        for (int i = 0; i < n; i++) {
            cin >> ps[i];
        }
        bool first = true;
        for (int i = 0; i < n; i++) {
            int out = ps[i] + 1;
            if (out == n + 1) {
                out = 1;
            }
            if (!first) {
                cout << ' ';
            }
            first = false;
            cout << out;
        }
        cout << "\n";
    }
    return 0;
}