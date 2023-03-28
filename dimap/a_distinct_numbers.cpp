#include <iostream>
#include <set>
using namespace std;

/*
Difficult: A
Score: 100
Link: https://vjudge.net/contest/550194#problem/A
Contest: Treino 02
*/

int main() {
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