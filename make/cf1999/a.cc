#include <iostream>
#include <vector>

using namespace std;

int main() {
    int tt; cin >> tt;
    for (int t = 0; t < tt; t++) {
        int n; cin >> n;
        int sum = 0;
        while (n > 0) {
            sum += n % 10;
            n /= 10;
        }
        cout << sum << "\n";
    }
    return 0;
}