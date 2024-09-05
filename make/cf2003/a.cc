#include <iostream>

using namespace std;

int main() {
    int tt; cin >> tt;
    for (int t = 0; t < tt; t++) {
        int n; cin >> n;
        string s; cin >> s;
        // n >= 2
        // s[0] != s[last]
        if (n >= 2 && s[0] != *s.rbegin()) {
            cout << "YES\n";
        } else {
            cout << "NO\n";
        }
    }
    return 0;
}