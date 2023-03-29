#include <iostream>
#include <set>
using namespace std;

/*
Difficult: -
Score: -
Link: https://cses.fi/problemset/task/1621
Contest: DIM0410 - Treino #2 STL
Status: Solved and Reviewed
*/

int main() {
    // Time: O(N)
    // Space: O(N)

    int n;
    cin >> n;

    int temp;
    set<int> values;

    for(int i = 0; i < n; i++) {
        cin >> temp;
        values.insert(temp);    
    }

    cout << values.size() << endl;
}