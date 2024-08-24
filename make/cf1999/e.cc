#include <iostream>
#include <vector>

using namespace std;

int main() {
    int tt; cin >> tt;
    for (int t = 0; t < tt; t++) {
        int l, r; cin >> l >> r;

        // l to r inclusive
        // operation: x -> 3x, y -> y/3

        // strategy: make l 0, and then choose l for x

        int ops = 0;
        int upper = 1;  // exclusive
        int num = l;
        while (num != 0) {
            num /= 3;
            upper *= 3;
            ops++;
        }
        int total = ops * 2;  // cost to reduce l to 0 and to bring others back

        int lower = l + 1;  // inclusive
        while (lower <= r) {
            while (upper <= lower) {
                upper *= 3;
                ops += 1;
            }
            int factor = min(upper, r+1) - lower;
            total += factor * ops;
            lower += factor;
        }
        cout << total << "\n";
    }
    return 0;
}