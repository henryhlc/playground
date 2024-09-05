#include <iostream>
#include <vector>
#include <set>
#include <map>
#include <algorithm>
#include <utility>

using namespace std;

int main() {
    int tt; cin >> tt;
    for (int t = 0; t < tt; t++) {
        int n; cin >> n;
        string s; cin >> s;
        map<char, int> freq;
        for (int i = 0; i < s.length(); i++) {
            if (freq.find(s[i]) == freq.end()) {
                freq[s[i]] = 0;
            }
            freq[s[i]]++;
        }


        auto comp = [](pair<char,int> a, pair<char,int> b) -> bool {
            if (a.second > b.second) {
                return true;
            } else if (a.second < b.second) {
                return false;
            } else if (a.first < b.first) {
                return true;
            } else if (a.first > b.first) {
                return false;
            } else {
                return false;
            }
        };

        set<pair<char,int>,decltype(comp)> chars(comp);
        for (auto k : freq) {
            chars.insert(k);
        }

        auto iter = chars.begin();
        while (iter != chars.end()) {
            auto [c, f] = *iter;
            cout << c;
            chars.erase(iter);
            auto secondIter = chars.begin();
            if (secondIter != chars.end()) {
                auto [c, f] = *secondIter;
                cout << c;
                chars.erase(secondIter);
                if (f > 1) {
                    chars.insert({c, f-1});
                }
            }
            if (f > 1) {
                chars.insert({c, f-1});
            }
            iter = chars.begin();
        }
        cout << '\n';
    }
    return 0;
}