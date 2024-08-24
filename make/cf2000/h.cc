#include <iostream>
#include <vector>
#include <map>
#include <set>

using namespace std;

constexpr int kMaxLength = 2000000;

int main() {
    int tt; cin >> tt;
    for (int t = 0; t < tt; t++) {
        // <=2e6 elements
        // 2e5 operations
        // add, may split a run into 2
        // remove, may merge 2 runs into 1
        // query, find the first run with size >= k


        map<int, set<int>> lengthToStarts;
        map<int, int> endToStart;
        map<int, int> startToLength;

        auto printIntervals = [&]() {
            bool first = true;
            for (auto [k,v] : startToLength) {
                if (!first) {
                    cout << " ";
                }
                cout << k << "-" << k+v-1 << ";";
                first = false;
            }
            cout << "\n";
            for (auto& [k,v] : lengthToStarts) {
                cout << k << ": ";
                for (auto e : v) {
                    cout << e << " ";
                }
                cout << "\n";
            }
        };

        auto removeRun = [&](int start) {
            int length = startToLength[start];
            int end = start + length - 1;
            startToLength.erase(start);
            endToStart.erase(end);
            lengthToStarts[length].erase(start);
            if (lengthToStarts[length].empty()) {
                lengthToStarts.erase(length);
            }
        };

        auto addRun = [&](int start, int length) {
            int end = start + length - 1;
            startToLength[start] = length;
            endToStart[end] = start;
            if (lengthToStarts.find(length) == lengthToStarts.end()) {
                lengthToStarts[length] = {start};
            } else {
                lengthToStarts[length].insert(start);
            }
        };

        addRun(1, kMaxLength + 1);


        auto query = [&](int k) -> int {
            auto iter = lengthToStarts.lower_bound(k);
            int start = -1;
            
            while (iter != lengthToStarts.end()) {
                if (start == -1) {
                    start = *iter->second.begin();
                } else {
                    start = min(start, *iter->second.begin());
                }
                iter++;
            }
            if (start != -1) {
                return start;
            } else {
                return startToLength.rbegin()->first;
            }
        };

        auto insertion = [&](int a) {
            if (startToLength.find(a) != startToLength.end()) {
                int oldStart = a;
                int oldLength = startToLength[a];
                removeRun(oldStart);
                // 1[a]1 
                if (oldLength == 1) {
                    return;
                }
                // 1[a...]1
                int newStart = a + 1;
                int newLength = oldLength - 1;
                addRun(newStart, newLength);
            } else if (endToStart.find(a) != endToStart.end()) {
                // 1[...a]1
                int start = endToStart[a];
                int oldLength = startToLength[start];
                removeRun(start);
                addRun(start, oldLength - 1);
            } else {
                // 1[.a..]1
                auto [end, start] = *endToStart.lower_bound(a);

                removeRun(start);

                int leftStart = start;
                int leftLength = a - start;
                addRun(leftStart, leftLength);

                int rightStart = a + 1;
                int rightLength = end - a;
                addRun(rightStart, rightLength);
           }

        };

        auto deletion = [&](int a) {
            int leftStart = -1;
            if (endToStart.find(a-1) != endToStart.end()) {
                leftStart = endToStart[a-1];
            }
            int rightStart = -1;
            if (startToLength.find(a+1) != startToLength.end()) {
                rightStart = a + 1;
            }
            int newStart = a;
            int newLength = 1;
            if (leftStart >= 0) {
                newStart = leftStart;
                newLength += startToLength[leftStart];
                removeRun(leftStart);
            }
            if (rightStart >= 0) {
                newLength += startToLength[rightStart];
                removeRun(rightStart);
            }
            addRun(newStart, newLength);
        };

        int n; cin >> n;
        for (int i = 0; i < n; i++) {
            int a; cin >> a;
            insertion(a);
            // printIntervals();
        }

        int m; cin >> m;
        bool first = true;
        for (int i = 0; i < m; i++) {
            char c; cin >> c;
            int k; cin >> k;
            switch (c) {
                case '?':
                    // cout << "query " << k << ": " << query(k) << "\n";
                    if (!first) {
                        cout << ' ';
                    }
                    first = false;
                    cout << query(k);
                    break;
                case '+':
                    // cout << "insert " << k << "\n";
                    insertion(k);
                    break;
                case '-':
                    // cout << "remove " << k << "\n";
                    deletion(k);
                    break;
            }
            // printIntervals();
        }
        cout << "\n";
    }
    return 0;
}