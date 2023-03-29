#include <iostream>
#include <set>
using namespace std;

/*
Difficult: A
Score: -
Link: https://cses.fi/problemset/task/1621
Contest: DIM0410 - Treino #2 STL
Status: Solved
*/

int main() {
    // Time: O(N*M)
    // Space: O(N+M)
    
    int m, n, temp;
    set<int> vm, vn;
    set<int>::iterator itr, jtr;

    cin >> m;
    for(int i = 0; i<m; i++) {
        cin >> temp;
        vm.insert(temp);
    }

    cin >> n;
    for(int i = 0; i<n; i++) {
        cin >> temp;
        vn.insert(temp);
    }

    for(itr = vm.begin(); itr != vm.end(); itr++) {
        for(jtr = vn.begin(); jtr != vn.end(); jtr++) {
            int value = *itr+*jtr;
            if(vm.find(value) == vm.end() && vn.find(value) == vn.end()) {
                cout << *itr << " " << *jtr << endl;
                return 0;
            }
        }
    }
}