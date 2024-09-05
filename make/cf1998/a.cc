#include <iostream>

using namespace std;

int main() {
    int tt; cin >> tt;
    for (int t = 0; t < tt; t++) {
        int xc, yc, k; cin >> xc >> yc >> k;
        if (k % 2 == 1) {
            cout << xc << ' ' << yc << '\n';
            k--;
        }
        int offset = 1;
        while (k > 0) {
            cout << xc << ' ' << yc + offset << '\n';
            k--;
            cout << xc << ' ' << yc - offset << '\n';
            k--;
            offset++;
        }
    }
    return 0;
}