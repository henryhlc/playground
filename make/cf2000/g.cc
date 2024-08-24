#include <iostream>
#include <vector>
#include <map>

using namespace std;

struct Edge {
    int to;
    int bus;
    int walk;
};

int main() {
    int tt; cin >> tt;
    for (int t = 0; t < tt; t++) {
        int n, m; cin >> n >> m;
        int t0, t1, t2; cin >> t0 >> t1 >> t2;
        //  t1  t2    t0
        //  v   v     v
        //  1 2 3 4 5 6         
        //  5 4 3 2 1 0
        int start = t0 - t2;  // can take bus as long as exit <= start
        int end = t0 - t1; // can take bus as long as board >= end

        vector<vector<Edge>> graph(n+1, vector<Edge>{});
        for (int i = 0; i < m; i++) {
            int u, v, bus, walk; cin >> u >> v >> bus >> walk;
            graph[v].push_back(Edge{.to=u, .bus=bus, .walk=walk});
            graph[u].push_back(Edge{.to=v, .bus=bus, .walk=walk});
        }

        multimap<int, int> queue;
        vector<int> processed(n+1, -1);
        processed[n] = 0;
        queue.insert({0, n});
        while (queue.size() > 0) {
            auto [time, node] = *queue.begin();
            queue.erase(queue.begin());
            if (processed[node] > 0) {
                continue;
            }
            processed[node] = time;
            for (auto e : graph[node]) {
                if (time + e.bus <= start || time >= end) {
                    if (time + e.bus <= t0) {
                        // take bus
                        queue.insert({time + e.bus, e.to});
                    }
                } else {
                    // wait for bus or walk
                    int bus = end + e.bus;
                    int walk = time + e.walk;
                    if (min(bus, walk) <= t0) {
                        queue.insert({min(bus, walk), e.to});
                    }
                }
            }
            graph[node].clear();
        }
        if (processed[1] < 0) {
            cout << -1 << "\n";
        } else {
            cout << t0 - processed[1] << "\n";
        }
    }
}