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
        int n; cin >> n;
        string s; cin >> s;
        vector<int> open;
        int score = 0;
        for (int i = 0; i < s.length(); i++) {
            if (s[i] == ')') {
                score += i - *open.rbegin();
                open.pop_back();
                continue;
            }
            if (s[i] == '(') {
                open.push_back(i);
                continue;
            }
            if (open.size() == 0) {
                open.push_back(i);
                continue;
            }
            score += i - *open.rbegin();
            open.pop_back();
        }
        cout << score << '\n';
    }
    return 0;
}