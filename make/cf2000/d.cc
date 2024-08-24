#include <iostream>
#include <vector>

using namespace std;

void problem(int n, const vector<int>& as, const string& s) {
    vector<long long> psum(n);
    psum[0] = as[0];
    for (int i = 1; i < n; i++) {
        psum[i] = psum[i-1] + as[i];
    }
    int L = 0;
    int R = n - 1;
    long long score = 0;
    while (L < R) {
        while (L < n && s[L] != 'L') {
            L++;
        }
        while (R >= 0 && s[R] != 'R') {
            R--;
        }
        if (L < R) {
            score += psum[R] - psum[L] + as[L];
            L++;
            R--;
        }
    }
    cout << score << "\n";
}

int main() {
    int tt; cin >> tt;
    for (int t = 0; t < tt; t++) {
        int n; cin >> n;
        vector<int> as(n);
        for (int i = 0; i < n; i++) {
            cin >> as[i];
        }
        string s; cin >> s;
        problem(n, as, s);
    }
    return 0;
}