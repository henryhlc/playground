#include <iostream>
#include <vector>

using namespace std;

int main() {
    int tt; cin >> tt;
    for (int t = 0; t < tt; t++) {
        int n; cin >> n;
        int best = 0;
        vector<int> freq(n+1, 0);
        for (int i = 0; i < n; i++) {
            int a; cin >> a;
            freq[a]++;
            best = max(best, freq[a]);
        }
        cout << n - best << '\n';
    }
    return 0;
}