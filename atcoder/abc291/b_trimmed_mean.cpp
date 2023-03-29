#include <iostream>
#include <vector>
#include <bits/stdc++.h>
using namespace std;

/*
Difficult: B
Score: 200
Link: https://atcoder.jp/contests/abc291/tasks/abc291_b
Contest: abc291
*/

int main() {
    // Time: O(N*LogN)
    // Space: O(N)

    int n, temp;
    cin >> n;
    vector<int> grade;
    float sum = 0;
    
    for(int i=0; i<n*5; i++) {
        cin >> temp;
        grade.push_back(temp);
    }

    sort(grade.begin(), grade.end());

    for(int i=n; i<grade.size()-n; i++){
        sum += grade[i];    
    }

    cout << sum/(n*3) << endl;
}