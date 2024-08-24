#include <iostream>
#include <functional>
#include <vector>
#include <set>

using namespace std;

int main() {
    int tt; cin >> tt;
    for (int t = 0; t < tt; t++) {
        int n, q; cin >> n >> q;

        vector<int> parent(n);  // [child-1] = parent
        parent[0] = -1;
        for (int i = 1; i < n; i++) {
            cin >> parent[i];
        }
        vector<int> perm(n);  // [pos] = node
        for (int i = 0; i < n; i++) {
            cin >> perm[i];
        }

        vector<int> posParent(n, -1);
        vector<vector<int>> posChildren(n, vector<int>{});
        vector<int> posNextSibling(n, -1);
        // This part relies on the perfect binary tree assumption.
        // [a, b) half open
        function<void(int,int)> computePosTree = [&](int a, int b) {
            int step = ((b - a) - 1) / 2;
            for (int begin = a + 1; begin < b; begin += step) {
                posChildren[a].push_back(begin);
                posParent[begin] = a;
                int next = begin + step;
                computePosTree(begin, next);
                if (next < b) {
                    posNextSibling[begin] = next;
                }
            }
        };
        computePosTree(0, n);

        auto correctParent = [&](int pos) -> bool {
            int node = perm[pos];
            int expectedNode = parent[node-1];  // -1 for node 1
            int parentPos = posParent[pos];  // -1 for pos 0
            if (parentPos == -1) {
                return expectedNode == -1;
            }
            // cout << "pos: " << pos << "; parentPos: " << parentPos << '\n';
            int actualNode = perm[parentPos];
            return expectedNode == actualNode;
        };

        int wrongParents = 0;
        for (int i = 0; i < n; i++) {
            if (!correctParent(i)) {
                wrongParents++;
            }
        }

        auto localWrongParents = [&](int pos1, int pos2) -> int {
            set<int> toCheck {pos1, pos2};
            for (auto pos : {pos1, pos2}) {
                for (auto c : posChildren[pos]) {
                    toCheck.insert(c);
                }
            }
            int count = 0;
            for (auto p : toCheck) {
                if (!correctParent(p)) {
                    count++;
                }
            }
            return count;
        };

        for (int i = 0; i < q; i++) {
            int pos1, pos2; cin >> pos1 >> pos2;
            pos1--; pos2--;
            int localWrongBefore = localWrongParents(pos1, pos2);
            int temp = perm[pos1];
            perm[pos1] = perm[pos2];
            perm[pos2] = temp;
            int localWrongAfter = localWrongParents(pos1, pos2);
            wrongParents += localWrongAfter - localWrongBefore;
            if (wrongParents == 0) {
                cout << "YES\n";
            } else {
                cout << "NO\n";
            }
        }
    }
    return 0;
}