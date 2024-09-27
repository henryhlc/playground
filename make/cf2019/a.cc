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
        vector<int> as(n);
        for (int i = 0; i < n; i++) {
            cin >> as[i];
        }
        // odd case and even case

        int odd_max = 0;
        int even_max = 0;
        for (int i = 0; i < n; i++) {
            if (i % 2 == 0) {
                even_max = max(even_max, as[i]);
            } else {
                odd_max = max(odd_max, as[i]);
            }
        }

        cout << max(odd_max + n/2, even_max + (n+1)/2) << '\n';
    }
    return 0;
}