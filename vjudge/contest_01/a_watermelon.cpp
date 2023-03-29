#include <iostream>
using namespace std;

/*
Difficult: A
Score: -
Link: https://codeforces.com/problemset/problem/155/A
Contest: DIM0410 - Contest #1
Status: Solved
*/

int main() {
    int k;
    cin >> k;

    // Time: O(1)
    // Space: O(1)

    if(k%2==0 && k!=2) {
        cout << "YES" << endl;
    } else {
        cout << "NO" << endl;
    }
}