#include <iostream>
#include <vector>

using namespace std;

// Needs to learn combinatorial game theory for this
int main() {
    int tt; cin >> tt;
    for (int t = 0; t < tt; ++t) {
        int n; cin >> n;
        vector<int> as(n);
        for (int i = 0; i < n; i++) {
            cin >> as[i];
        }

        // odd > 2
        // - play first to play last, k -> 2
        // - play first to play second last, k -> 1
        // even
        // - play first to play last, not possible
        // - play first to play second last, k -> 1

        // odd, to k-1, k-2, 2, 1
        // even, to k-1, 1

        // not just win, but to control the outcome?
        // - we can ignore the case <= 2
        // even, B -> Bob even side play to last, B side to control
        // even, A -> Alice even side play to second last, A side to control
        // odd, B -> Alice odd to 2, A gets control
        // odd, A -> 

        string winner = "Alice";
        if (as[0] % 2 == 0) {
            winner = "Bob";
        }
        
        for (auto i = 1; i < n; i++) {
            if (as[i] <= 2) {
                continue;
            }
            if (winner == "Bob" && as[i] % 2 == 0) {
                // second, second
                winner = "Bob";
            } else if (winner == "Alice" && as[i] % 2 == 0) {
                // second, first
                winner = "Alice";
            } else if (winner == "Bob" && as[i] % 2 == 1) {
                // first, second
                winner = "Alice";
            } else if (winner == "Alice" && as[i] % 2 == 1) {
                // first, first
                winner = "Bob";
            }
        }
        cout << winner << "\n";
    }
    return 0;
}
