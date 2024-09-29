#include<algorithm>
#include<iostream>
#include<map>
#include<set>
#include<vector>

using namespace std;
using ll = long long;

struct Op {
    int a, d, k;
};

struct St {
    int d, off;
    bool operator<(const St other) const {
        if (d != other.d) {
            return d < other.d;
        }
        return off < other.off;
    }
};

struct Ufind {
    vector<int> ps;
    vector<int> ss;

    Ufind(int n) {
        for (int i = 0; i < n; i++) {
            ps.push_back(i);
            ss.push_back(1);
        }
    }

    int f(int i) {
        if (ps[i] == i) {
            return i;
        }
        ps[i] = f(ps[i]);
        return ps[i];
    }

    void u(int i, int j) {
        int gi = f(i);
        int gj = f(j);
        if (gi == gj) {
            return;
        }
        if (ss[gi] < ss[gj]) {
            ps[gi] = gj;
            ss[gj] += ss[gi];
        } else {
            ps[gj] = gi;
            ss[gi] += ss[gj];
        }
    }

    int c() {
        int count = 0;
        for (int i = 0; i < ps.size(); i++) {
            if (ps[i] == i) {
                count++;
            }
        }
        return count;
    }
};

int main() {
    int tt; cin >> tt;
    for (int t = 0; t < tt; t++) {
        int n, m; cin >> n >> m; 

        map<int, vector<Op>> ops;
        for (int i = 0; i < m; i++) {
            Op op; cin >> op.a >> op.d >> op.k;
            if (ops.find(op.a) == ops.end()) {
                ops[op.a] = vector<Op>();
            }
            ops[op.a].push_back(op);
        }

        Ufind uf(n);

        map<St, int> ends;
        for (int i = 1; i <= n; i++) {
            bool hasP = false;
            for (int d = 1; d <= 10; d++) {
                St key {.d = d, .off = i % d};
                if (ends.find(key) != ends.end() && ends[key] >= i)  {
                    uf.u(i-1, i-d-1);
                }
            }
            for (auto op : ops[i]) {
                int end = op.a + op.k * op.d;
                St key {.d = op.d, .off = i % op.d};
                if (ends.find(key) == ends.end()) {
                    ends[key] = end;
                } else {
                    ends[key] = max(ends[key], end);
                }
            }
        }
        cout << uf.c() << '\n';
    }
    return 0;
}