#include <iostream>
#include <unordered_set>

using namespace std;

int main() {
    int tt; cin >> tt;
    for (int t = 0; t < tt; t++) {
        int nn; cin >> nn;

        unordered_set<int> occupied {};
        bool yes = true;
        for (int n = 0; n < nn; n++) {
            int a; cin >> a;
            if (n > 0 && occupied.count(a-1) == 0 && occupied.count(a+1) == 0) {
                yes = false;
            }
            occupied.insert(a);
        }

        if (yes) {
            cout << "YES\n";
        } else {
            cout << "NO\n";
        }
    }
    return 0;
}