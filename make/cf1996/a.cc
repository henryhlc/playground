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
        int count = n / 4 + (n % 4) / 2;
        cout << count << '\n';
    }
    return 0;
}