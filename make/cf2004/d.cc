#include <iostream>
#include <vector>
#include <algorithm>
#include <map>

using namespace std;

int dist(int a, int b) {
    if (a >= b) {
        return a - b;
    }
    return b - a;
}

bool overlap(string a, string b) {
    return a[0] == b[0] || a[0] == b[1] || a[1] == b[0] || a[1] == b[1];
}

vector<string> hop_candidates(string a, string b) {
    vector<string> res{};
    vector<string> candidates {"BG", "BR", "BY", "GR", "GY", "RY"};
    for (int i = 0; i < candidates.size(); i++) {
        if (candidates[i] != a && candidates[i] != b) {
            res.push_back(candidates[i]);
        }
    }
    return res;
}

int main() {
    int tt; cin >> tt;
    for (int t = 0; t < tt; ++t) {
        int n, q; cin >> n >> q;
        vector<string> cities(n);
        map<string, vector<int>> type_locations;
        for (int i = 0; i < n; i++) {
            cin >> cities[i];
            if (type_locations.find(cities[i]) == type_locations.end()) {
                type_locations[cities[i]] = vector<int>{ i };
            } else {
                type_locations[cities[i]].push_back(i);
            }
        }
        vector<int> from(q);
        vector<int> to(q);
        for (int i = 0; i < q; i++) {
            cin >> from[i] >> to[i];
        }

        for (int i = 0; i < q; i++) {
            // cout << "q: " << i << "\n";
            string s = cities[from[i]-1];
            string t = cities[to[i]-1];
            int max_i = max(from[i]-1, to[i]-1);
            int min_i = min(from[i]-1, to[i]-1);
            // cout << s << t << "\n";
            if (overlap(s, t)) {
                cout << dist(from[i], to[i]) << "\n";
            } else {
                int best = -1;
                // for each possible intermediate step, find the closest
                for (auto& c : hop_candidates(s, t)) {
                    if (type_locations.find(c) == type_locations.end()) {
                        continue;
                    }
                    auto& locations = type_locations[c];
                    if (locations.size() == 0) {
                        continue;
                    }
                    // cout << "locations: ";
                    // for (int i = 0; i < locations.size(); ++i) {
                        // cout << locations[i] << ' ';
                    // }
                    // cout << "\n";
                    int low = -1;
                    int high = locations.size();
                    while (high - low > 1) {
                        auto mid = (high + low) / 2;
                        if (locations[mid] > max_i) {
                            high = mid;
                        } else {
                            low = mid;
                        }
                    }
                    if (low >= 0) {
                        if (best == -1) {
                            best = dist(locations[low], from[i]-1) + dist(locations[low], to[i]-1);
                        } else {
                            best = min(best, dist(locations[low], from[i]-1) + dist(locations[low], to[i]-1));
                        } 
                    }
                    if (high < locations.size()) {
                        if (best == -1) {
                            best = dist(locations[high], from[i]-1) + dist(locations[high], to[i]-1);
                        } else {
                            best = min(best, dist(locations[high], from[i]-1) + dist(locations[high], to[i]-1));
                        }
                    }
                }
                cout << best << "\n";
            }
        }
    }
    return 0;
}
