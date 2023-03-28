#include <iostream>
#include <set>
using namespace std;

/*
Difficult: B
Score: 200
Link: https://vjudge.net/contest/550194#problem/B
Contest: Treino 02
*/

int main() {
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
            if(vm.find(value) == vm.end() || vn.find(value) == vn.end()) {
                cout << *itr << " " << *jtr << endl;
                return 0;
            }
        }
    }
}