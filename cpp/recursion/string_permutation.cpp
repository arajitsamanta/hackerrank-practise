#include <iostream>
using namespace std;

void indent(int n){
  for(int i=0;i<n;i++){
    cout<<"    ";
  }
}

void permuteHelper(string soFar, string rest)
{
  indent(soFar.length());
  cout << "permute(\"" << soFar << "\",\""<< rest << "\")" << endl;
  if (rest.empty())
  {
    cout << soFar << endl;
  }
  else
  {
    for (int i = 0; i < rest.length(); i++)
    {

      //cout<<rest.substr(0, i)<<" "<<rest.substr(i + 1)<<endl;
      string remaining = rest.substr(0, i) + rest.substr(i + 1);
      //cout<<soFar+rest[i]<<" "<<remaining<<endl;
      permuteHelper(soFar + rest[i], remaining);
    }
  }
}

void permute(string s)
{
  permuteHelper("", s);
}

int main()
{
  permute("abcd");
}