#include <iostream>
#include <vector>

using namespace std;

int main() {
    int tt; cin >> tt;
    for (int t = 0; t < tt; t++) {
        int n; cin >> n;
        if (n % 2 == 1) {
            bool first = true;
            for (int i = 1; i <= n; i += 2) {
                if (!first) {
                    cout << ' ';
                }
                first = false;
                cout << i;
            }
            for (int i = n-1; i >= 2; i -= 2) {
                cout << ' ' << i;
            }
            cout << "\n";
        } else {
            cout << -1 << "\n";
        }
    }
    return 0;
}