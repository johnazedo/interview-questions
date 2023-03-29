#include <iostream>
#include <map>
using namespace std;

/*
Difficult: -
Score: -
Link: https://onlinejudge.org/index.php?option=com_onlinejudge&Itemid=8&page=show_problem&problem=2949
Contest: DIM0410 - Treino #2 STL
Status: Solved
*/

int main() {
    // Time: O(N)
    // Space: O(N)
    
    int m, n, temp;
    while (cin >> m >> n && m != 0 && n != 0)
    {   
        map<int, bool> values;
        for(int i = 0; i<m+n; i++) {
            cin >> temp;
            if(values.find(temp) != values.end()) {
                values[temp] = true;
            } else {
                values[temp] = false;
            }
        }

        int count = 0;
        for(auto it = values.begin(); it != values.end(); it++) {
            if(it->second == true) {
                count++;
            }
        }

        cout << count << endl;
    } 
}