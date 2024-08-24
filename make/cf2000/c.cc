#include <iostream>
#include <vector>
#include <unordered_map>

using namespace std;

string problem(const vector<int>& as, string p) {
    if (p.length() != as.size()) {
        return "NO\n";
    }
    unordered_map<char, int> forward;
    unordered_map<int, char> back;
    for (int i = 0; i < p.length(); i++) {
        bool seen_c = forward.find(p[i]) != forward.end();
        bool seen_a = back.find(as[i]) != back.end();
        if (seen_c != seen_a) {
            return "NO\n";
        } else if (!seen_c && !seen_a) {
            forward[p[i]] = as[i];
            back[as[i]] = p[i];
        } else if (seen_c && seen_a) {
            if (forward[p[i]] != as[i] || back[as[i]] != p[i]) {
                return "NO\n";
            }
        }
    }
    return "YES\n";
}

int main() {
    int tt; cin >> tt;
    for (int t = 0; t < tt; t++) {
        int n; cin >> n;
        vector<int> as(n);
        for (int i = 0; i < n; i++) {
            cin >> as[i];
        }
        int m; cin >> m;
        for (int i = 0; i < m; i++) {
            string p; cin >> p;
            cout << problem(as, p);
        }
    }
}