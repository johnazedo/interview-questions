#include <iostream>
#include <map>
using namespace std;

int main() {
    int m, n, temp;
    cin >> m;
    cin >> n;

    map<int, bool> values;

    for(int i = 0; i<m+n; i++) {
        cin >> temp;
        if(values.find(temp) != values.end()) {
            values[temp] = true;
        } else {
            values[temp] = false;
        }
    }

    int count = 0;
    for(auto it = values.begin(); it != values.end(); it++) {
        if(it->second == true) {
            count++;
        }
    }

    cout << count << endl;
}