#include <iostream>
#include <utility>
#include <vector>
#include <algorithm>

using namespace std;
using pii = pair<int,int>;
using ll = long long;

int main() {
    int tt; cin >> tt;
    for (int t = 0; t < tt; t++) {
        int n, k; cin >> n >> k;
        vector<pii> numAndId(n);
        for (int i = 0; i < n; i++) {
            cin >> numAndId[i].first;
            numAndId[i].second = i;
        }
        vector<int> movable(n);
        for (int i = 0; i < n; i++) {
            cin >> movable[i];
        }

        sort(numAndId.begin(), numAndId.end());

        auto newMid = [](int n, int i) {
            int newMidPos = (n-1+1)/2-1;
            if (i <= newMidPos) {
                return newMidPos + 1;
            }
            return newMidPos;
        };

        ll bestScore = 0;
        for (int i = 0; i < n; i++) {
            int idx = numAndId[i].second;
            if (!movable[idx] == 1) {
                continue;
            }
            bestScore = max(bestScore, (ll) numAndId[i].first + k + numAndId[newMid(n,i)].first);
        }
        vector<int> choices {(n+1)/2-1-1, (int) numAndId.size()-1};
        for (int i = 0; i < choices.size(); i++) {
            int idx = choices[i];
            if (idx < 0) {
                continue;
            }
            // cout << "idx: " << idx << '\n';
            int currMedianIdx = newMid(n,idx);
            int currMedian = numAndId[currMedianIdx].first;
            ll score = (ll) numAndId[idx].first + currMedian;
            // cout << "score: " << score << "\n";

            vector<int> moveCandidate;
            for (int j = 0; j < currMedianIdx; j++) {
                if (movable[numAndId[j].second] == 1) {
                    moveCandidate.push_back(numAndId[j].first);
                }
            }

            int left = k;
            int toMove = 0;
            int m = currMedianIdx;  // idx of the current median
            while (left > 0) {
                toMove++;
                if (movable[numAndId[m].second] != 1) {
                    if (moveCandidate.empty()) {
                        break;
                    }
                    left -= numAndId[m].first - *moveCandidate.rbegin();
                    moveCandidate.pop_back();
                    if (left < 0) {
                        break;
                    }
                }
                int budgetConstrainedMax = left / toMove;
                int upperIdx = m + 1;
                if (upperIdx == idx) {
                    upperIdx++;
                }
                if (upperIdx < n) {
                    int upperConstrainedMax = numAndId[upperIdx].first - numAndId[m].first;
                    if (upperConstrainedMax < budgetConstrainedMax) {
                        score += upperConstrainedMax;
                        left -= toMove * upperConstrainedMax;
                        m = upperIdx;
                        continue;
                    }
                }
                score += budgetConstrainedMax;
                break;
            }
            bestScore = max(bestScore, score);
        }
        std::cout << bestScore << '\n';
    }
    return 0;
}
