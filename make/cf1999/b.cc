#include <iostream>

using namespace std;

int main() {
    int tt; cin >> tt;
    for (int t = 0; t < tt; t++) {
        int a1, a2, b1, b2; cin >> a1 >> a2 >> b1 >> b2;
        int count = 0;

        auto win = [](int x1, int x2, int y1, int y2) -> bool {
            int numWin = 0;
            int numLoss = 0;
            if (x1 > x2) {
                numWin++;
            } else if (x1 < x2) {
                numLoss++;
            }
            if (y1 > y2) {
                numWin++;
            } else if (y1 < y2) {
                numLoss++;
            }
            return numWin > numLoss;
        };

        // a1 vs b1, 
        if (win(a1, b1, a2, b2)) {
            count += 2;
        }
        // a1 vs b2
        if (win(a1, b2, a2, b1)) {
            count += 2;
        }
        cout << count << "\n";
    }
    return 0;
}
