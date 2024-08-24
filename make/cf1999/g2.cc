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
            int dist = high - low;
            int m1 = low + dist / 3;
            int m2 = low + 2 * dist / 3;

            printf("? %d %d\n", m1, m2);
            fflush(stdout);
            int a; cin >> a;
            if (a >= (m1+1)*(m2+1)) {
                high = m1;
            } else if (a >= m1*(m2+1)) {
                high = m2;
                low = m1;
            } else {
                low = m2;
            }
        }
        printf("! %d\n", high);
        fflush(stdout);
    }
    return 0;
}