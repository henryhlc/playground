#include <iostream>
#include <vector>
#include <map>
#include <deque>

using namespace std;

int main() {
    int tt; cin >> tt;
    for (int t = 0; t < tt; t++) {
        int n; cin >> n;
        vector<int> as(n);
        vector<int> lastOcc(n+1, -1);
        for (int i = 0; i < n; i++) {
            cin >> as[i];
            lastOcc[as[i]] = i;
        }

        vector<int> seq;
        map<int, deque<int>> numToIdx;
        vector<bool> used(n+1, false);

        int right = 0;
        bool findmax = true;
        int minIdx = 0;
        while (right < n) {
            // find the index that the next chosen element must be before
            while (right < n) {
                int a = as[right];
                if (!used[a]) {
                    if (numToIdx.find(a) == numToIdx.end()) {
                        numToIdx[a] = {};
                    }
                    numToIdx[a].push_back(right);
                    if (lastOcc[a] == right) {
                        break;
                    }
                }
                right++;
            }

            while (numToIdx.size() > 0) {
                int a = findmax? numToIdx.rbegin()->first : numToIdx.begin()->first;
                deque<int>& ids = numToIdx[a];

                while (!ids.empty() && ids.front() < minIdx) {
                    ids.pop_front();
                }
                if (ids.empty()) {
                    numToIdx.erase(a);
                } else {
                    seq.push_back(a);
                    used[a] = true;
                    minIdx = ids.front() + 1;
                    numToIdx.erase(a);
                    findmax = !findmax;
                    if (right < n && a == as[right]) {
                        break;
                    } 
                }
            }
            right++;
        }

        cout << seq.size() << '\n';
        bool first = true;
        for (auto v : seq) {
            if (!first) {
                cout << ' ';
            }
            first = false;
            cout << v;
        }
        cout << "\n";
    }
    return 0;
}