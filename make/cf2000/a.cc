#include <iostream>
#include <vector>

using namespace std;

int main() {
    int tt; cin >> tt;
    for (int t = 0; t < tt; t++) {
        int a; cin >> a;
        // Check
        // 10(n), where n > 2, n does not lead with 0
        vector<int> digits;
        while (a > 0) {
            digits.push_back(a % 10);
            a /= 10;
        }
        bool no = digits.size() < 3
            || digits[digits.size()-1] != 1
            || digits[digits.size()-2] != 0
            || digits[digits.size()-3] == 0
            || (digits.size() == 3 && digits[0] == 1); 

        if (no) {
            cout << "NO\n";
        } else {
            cout << "YES\n";
        }
    }
    return 0;
}