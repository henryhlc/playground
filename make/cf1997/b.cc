#include <iostream>
#include <vector>

using namespace std;

int main() {
    int tt; cin >> tt;
    for (int t = 0; t < tt; t++) {
        int n; cin >> n;
        vector<string> rs(2);
        cin >> rs[0] >> rs[1];
        
        int count = 0;
        for (int i = 1; i < n-1; i++) {
            for (int j = 0; j < 2; j++) {
                int o = (j + 1) % 2;
                if (rs[j][i-1] == '.'
                    && rs[j][i] == '.'
                    && rs[j][i+1] == '.'
                    && rs[o][i-1] == 'x'
                    && rs[o][i] == '.'
                    && rs[o][i+1] == 'x') {
                    count++;
                }
            }
        }
        cout << count << '\n';
    }
    return 0;
}