#include <iostream>

using namespace std;

/**
Suppose you are standing in a line to buy a movie ticket. You want to know the
price of the ticket, but only the first person in line knows the price.

You can approach the first person and ask him the price directly.
However, if you leave the line, someone else will take your place.
Therefore, you will have to use the following approach.

You will ask for the ticket price from the person in front of you.

That person does not want to leave the line either. Therefore, they will ask
the same question from the person in front of them. This process will continue
until the price is asked from the first person in the line.

The first person in line will tell the ticket price. After that, each person in
line will know the ticket price through the person standing in front of them.
*/

int ticketPriceForNumPeopleInLine(int person)
{
  int price;

  // base case.
  if (person == 1)
  {
    price = 100;
    return price;
  }
  else
  {
    // recursive case.
    cout << "Person" << person << " is asking price" << endl;
    person--;
    price = ticketPriceForNumPeopleInLine(person);
    cout << "Person" << person << " is telling price" << endl;
    return price;
  }
}

int main()
{
  int price;
  price = ticketPriceForNumPeopleInLine(5);
  cout << "Ticket price = " << price << endl;
}