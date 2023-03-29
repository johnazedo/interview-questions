#include <iostream>
using namespace std;

/*
Difficult: I
Score: -
Link: https://codeforces.com/gym/103960/problem/I
Contest: DIM0410 - Contest #1 
Status: Solved
*/

int main() {
    // Time: O(N)
    // Space: O(1)

    int v;
    for(int i=0; i<8; i++) {
        cin >> v;
        if(v==9) {
            cout << "F" << endl;
            return 0;
        }
    }

    cout << "S" << endl;
}