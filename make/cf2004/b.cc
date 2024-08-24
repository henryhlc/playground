
#include <iostream>
#include <vector>
#include <algorithm>

using namespace std;

int main() {
    int tt; cin >> tt;
    for (int t = 0; t < tt; ++t) {
        int l, r, L, R; cin >> l >> r >> L >> R;
        
        // If the intervals do not overlap, lock between, answer is 1
        // [    ]
        //        [     ]

        // If the intervals overlap, need to block all the doors in the overlap
        // [    ]
        //    [     ]

        // [    o]
        //      [    ]
        // overlap size 1, close 2 doors

        int intL = max(l, L);
        int intR = min(r, R);

        if (intL > intR) {
            // [   ]
            //       [   ]
            cout << 1 << "\n";
            continue;
        }


        int doors = 0;
        if (l < intL || L < intL) {
            // block rooms before intersection to the intersection
            //  [ A| ]
            //     [B, ,]
            doors++;
        }

        if (r > intR || R > intR) {
            doors++;
        }

        doors += intR - intL;
        cout << doors << "\n";
    }
    return 0;
}