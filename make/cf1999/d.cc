#include <iostream>

using namespace std;

int main() {
    int tt; cin >> tt;
    for (int k = 0; k < tt; k++) {
        string s, t; cin >> s >> t;
        int tIter = 0;
        int sIter = 0;
        while (tIter < t.length() && sIter < s.length()) {
            if (s[sIter] == '?') {
                s[sIter] = t[tIter];
                sIter++;
                tIter++;
            } else if (s[sIter] == t[tIter]) {
                sIter++;
                tIter++;
            } else {
                sIter++;
            }
        }
        if (tIter < t.length()) {
            cout << "NO\n";
        } else {
            while (sIter < s.length()) {
                if (s[sIter] == '?') {
                    s[sIter] = 'a';
                }
                sIter++;
            }
            cout << "YES\n" << s << "\n";
        }
    }
    return 0;
}