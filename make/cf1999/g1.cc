#include <iostream>
#include <cstdio>
#include <vector>

using namespace std;

int main() {
    int tt; cin >> tt;
    for (int t = 0; t < tt; t++) {
        // Incorrectly measures as y+1 if y >= x
        // 1, 2, 3, 5, 6, ...

        int low = 1;
        int high = 1000;
        while (high - low > 1) {
            int mid = (high + low) / 2;
            printf("? %d %d\n", mid, mid);
            fflush(stdout);
            int a; cin >> a;
            if (a > mid * mid) {
                high = mid;
            } else {
                low = mid;
            }
        }
        printf("! %d\n", high);
        fflush(stdout);
    }
    return 0;
}