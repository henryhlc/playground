#include <iostream>
#include <vector>
#include <map>

using namespace std;

int main() {
    int tt; cin >> tt;
    for (int t = 0; t < tt; t++) {
        int n, m, k; cin >> n >> m >> k;
        cout << min(n,k) * min(m, k) << "\n";
    }
}