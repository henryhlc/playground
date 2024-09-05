#include <iostream>

using namespace std;

string other(char a) {
    if (a == 'a') {
        return "b";
    }
    return "a";
}

int main() {
    int tt; cin >> tt;
    for (int t = 0; t < tt; t++) {
        string s; cin >> s;
        string res = "";
        for (int i = 1; i < s.length(); i++) {
            if (s[i] == s[i-1]) {
                res = s.substr(0,i) + other(s[i]) + s.substr(i);
                break;
            }
        }
        if (res == "") {
            res = other(s[0]) + s;
        }
        cout << res << '\n';
    }
    return 0;
}