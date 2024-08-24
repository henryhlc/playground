#include <iostream>
#include <vector>
#include <cstdio>

using namespace std;

int main() {
    int tt; cin >> tt;
    for (int t = 0; t < tt; t++) {
        int n; cin >> n;
        int a = 1;
        int b = 2;
        vector<int> edges;
        while (true) {
            cout << "? " << a << ' ' << b << "\n" << flush;
            int d; cin >> d;
            if (d == -1) {
                return 0;
            }
            if (d == a) {
                edges.push_back(a);
                edges.push_back(b);
                a = 1;
                b++;
            } else {
                a = d;
            }
            if (b == n + 1) {
                cout << "!";
                for (auto e : edges) {
                    cout << ' ' << e;
                }
                cout << "\n" << flush;
                break;
            }
        }
    }
    return 0;
}