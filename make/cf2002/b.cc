#include <iostream>
#include <vector>
#include <map>

using namespace std;

int main() {
    int tt; cin >> tt;
    for (int t = 0; t < tt; t++) {
        int n; cin >> n;
        vector<int> as(n);
        vector<int> bs(n);

        for (int i = 0; i < n; i++) {
            cin >> as[i];
        }
        for (int i = 0; i < n; i++) {
            cin >> bs[i];
        }

        // Bob's only winning strategy is to mimic Alice
        // Possible if two arrays are the same, or in reverse
        bool isSame = true;
        for (int i = 0; i < n; i++) {
            if (as[i] != bs[i]) {
                isSame = false;
                break;
            }
        }
        bool isReverse = true;
        for (int i = 0; i < n; i++) {
            if (as[i] != bs[bs.size()-1-i]) {
                isReverse = false;
                break;
            }
        }
        string winner = isSame || isReverse? "Bob" : "Alice";

        cout << winner << "\n";
    }

}