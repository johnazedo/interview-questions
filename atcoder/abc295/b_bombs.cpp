#include <iostream>
#include <vector>
#include <bits/stdc++.h>
#define DESTROYED_SQUARE '.'
using namespace std;

/*
Difficult: B
Score: 200
Link: https://atcoder.jp/contests/abc295/tasks/abc295_b
Contest: abc295
*/

class Bomb {
    public:
        int i;
        int j;
        int power;
};

// TODO: Try other solution
int main() {
    /*
    ASCII
    # = 35
    . = 46
    1..9 = 49..57
    */

   // Time: O(M*N^2) - Where M is the number of bombs;
   // Space: O(N^2)

    int m, n;
    cin >> m >> n;

    char temp;
    char board[m][n];

    vector<Bomb> bombs;

    for(int i=0; i<m; i++) {
        for(int j=0; j<n; j++) {
            cin >> temp;
            board[i][j] = temp;
            if(int(temp) > 48) {
                Bomb bomb;
                bomb.i = i;
                bomb.j = j;
                bomb.power = int(temp)-48;
                bombs.push_back(Bomb(bomb));
            }
        }
    }

    for(auto b : bombs) {
        for(int i=0; i<m; i++) {
            for(int j=0; j<n; j++) {
                if(abs(i - b.i) + abs(j - b.j) <= b.power) {
                    board[i][j] = DESTROYED_SQUARE;
                } 
            }
        }
    }

    for(int i=0; i<m; i++) {
        for(int j=0; j<n; j++) {
            cout << board[i][j];
        }
        cout << endl;
    }
}