#include <iostream>
#include <vector>

using namespace std;

int main() {
    int t;
    cin >> t;
    for (int tt = 0; tt < t; ++tt) {
        int n;
        cin >> n;
        vector<long long> x(n);
        vector<long long> y(n);
        for (int i = 0; i < n; ++i) {
            cin >> x[i] >> y[i];
        }
        long long sx, sy, tx, ty;
        cin >> sx >> sy >> tx >> ty;

        long long dist = (tx-sx)*(tx-sx) + (ty-sy)*(ty-sy);
        bool possible = true;
        for (int i = 0; i < n; ++i) {
            long long cdist = (x[i]-tx)*(x[i]-tx) + (y[i]-ty)*(y[i]-ty);
            if (cdist <= dist) {
                possible = false;
            }
        }
        if (possible) {
            cout << "YES\n";
        } else {
            cout << "NO\n";
        }
    }
    return 0;
}