#include <iostream>
#include <vector>

using namespace std;

int main() {
    int tt; cin >> tt;
    for (int t = 0; t < tt; t++) {
        int n, s, m; cin >> n >> s >> m;
        vector<int> open(n);
        vector<int> close(n);
        for (int i = 0; i < n; i++) {
            cin >> open[i] >> close[i];
        }
        int intervalOpen = 0;
        string res = "NO";
        for (int i = 0; i < n; i++) {
            int intervalClose = open[i];
            if (intervalClose - intervalOpen >= s) {
                res = "YES";
                break;
            }
            intervalOpen = close[i];
        }
        if (m - intervalOpen >= s) {
            res = "YES";
        }
        cout << res << "\n";
    }
    return 0;
}