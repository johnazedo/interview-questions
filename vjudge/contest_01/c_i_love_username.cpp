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
    // Time: O(N)
    // Space: O(1) 

    int m, temp;
    int count = 0;
    int high = -1;
    int lower = 10001;
    cin >> m;

    for(int i=0; i<m; i++) {
        cin >> temp;
        
        if(temp > high) {
            high = temp;
            count++;
        }

        if (temp < lower) {
            lower = temp;
            count++;
        }
    }

    cout << count - 2 << endl;
}