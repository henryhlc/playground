#include <iostream>
#include <vector>
#include <algorithm>

using namespace std;

int main() {
    int tt; cin >> tt;
    for (int t = 0; t < tt; ++t) {
        int n; cin >> n;
        vector<int> nums(n);
        for (int i = 0; i < n; ++i) {
            cin >> nums[i];
        }

        sort(nums.begin(), nums.end());
        int dist = 1;
        int second = -1;
        for (int i = 1; i < n; ++i) {
            if (nums[i] != nums[i-1]) {
                dist++;
                if (second == -1) {
                    second = nums[i];
                }
            }
        }
        if (dist <= 2 && second > nums[0] + 1) {
            cout << "YES\n";
        } else {
            cout << "NO\n";
        }
    }
    return 0;
}