#include <iostream>
#include <vector>
using namespace std;

/*
Difficult: -
Score: -
Link: https://cses.fi/problemset/task/2216/
Contest: DIM0410 - Contest #1 
Status: Unsolved
*/

int main() {
    int m, temp;
    cin >> m;

    vector<int> v;

    for(int i=0; i<m; i++) {
        cin >> temp;
        v.push_back(temp);
    }

    int count = 0;
    int next = 1;


    for(int j=0; j<m; j++) {
        count++;

        for(int i=0; i<m; i++) {
            if(v[i] == next) {
                next++;
            }
        }

        if((next-1) == v.size()) {
            break;
        }
    }

    cout << count << endl;

}