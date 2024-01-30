#include <bits/stdc++.h>

std::vector<int> select_sort_array(std::vector<int>& array) {
    std::vector<int> sorted_array;
    int size = array.size();
    int max, index;

    // All lines inside this for has O(N) as base time complexity 
    for(;;) {
        if(sorted_array.size() == size) break;
        max = INT_MIN;

        // Search on array takes O(N) so this for has 
        // O(N*N) time complexity
        for(int i = 0; i<array.size(); i++) {
            if(array[i] > max) {
                max = array[i];
                index = i;
            }
        }

        sorted_array.push_back(max);

        // Delete on array takes O(N) so this piece of code 
        // has O(N*N) time complexity
        array.erase(array.begin()+index);
    }

    return sorted_array;
}
