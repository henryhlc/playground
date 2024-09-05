#include <iostream>
#include <vector>
#include <algorithm>

using namespace std;

int main() {
    int tt; cin >> tt;
    for (int t = 0; t < tt; t++) {
        int n; cin >> n;
        vector<int> a(n);
        for (int i = 0; i < n; i++) {
            cin >> a[i];
        }
        sort(a.begin(), a.end());
        cout << a[n/2+1-1] << '\n';
    }
    return 0;
}