#include <iostream>
#include <map>
using namespace std;

/*
Difficult: C
Score: 300
Link: https://vjudge.net/contest/550194#problem/C
Contest: Treino 02
*/

int main() {
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