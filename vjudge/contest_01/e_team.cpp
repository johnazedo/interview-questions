#include <iostream>
#include <vector>
using namespace std;

/*
Difficult: A
Score: -
Link: https://codeforces.com/problemset/problem/231/A
Contest: DIM0410 - Contest #1 
Status: Solved
*/

int main() {
    // Time: O(N)
    // Space: O(1)

    int m;
    cin >> m;

    int count = 0;
    int peyta, vasya, tonya;

    for(int i = 0; i<m; i++) {
        cin >> peyta >> vasya >> tonya;

        if((peyta == 1 && vasya == 1) || 
           (vasya == 1 && tonya == 1) || 
           (tonya == 1 && peyta == 1)) {
            count++;
        }
    }

    cout << count << endl;
}