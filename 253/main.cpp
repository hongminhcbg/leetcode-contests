
// Note that by default C++ creates a max-heap
// for priority queue
#include <iostream>
#include <queue>
  
using namespace std;
  
void showpq(priority_queue<int> gq)
{
    priority_queue<int> g = gq;
    while (!g.empty()) {
        cout << '\t' << g.top();
        g.pop();
    }
    cout << '\n';
}

int minStoneSum(vector<int>& piles, int k){
    priority_queue<int> gquiz;
    for(int i = 0; i < piles.size(); i++){
        gquiz.push(piles[i]);
    }
   for(int i = 0; i < k; i++){
        int temp = gquiz.top();
        gquiz.pop();
        temp = temp - temp/2;
        gquiz.push(temp)
    }
   int sum = 0;
   while(!gquiz.empty()){
        sum += gquiz.top();
        gquiz.pop();
   }
    return sum;
}

func minSwaps(s string) int {
    numClose := 0
    numOpen := 0
    for int i := 0; i < len(s); i++ {
        if s[i] == '[' {
            numOpen++
            continue
        }
        if s[i] == ']' {
            if numOpen > 0 {
                numOpen--
                continue
            }

            numClose++
        }
    }
    if numClose % 2 == 0 {
        return numClose / 2
    }

    return numClose / 2 + 1
}
int main()
{
    fmt.Println(minSwaps("[[[[]]]")) 
    return 0;
}
