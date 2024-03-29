// https://leetcode.com/problems/design-a-food-rating-system/
#include <vector>
#include <string>
#include <unordered_map>
#include <set>
#include <initializer_list>

using namespace std;

class FoodRatings
{
public:
  // sets automatically order pairs by their first property. If the first property is the same
  // for both pairs then it will order it by its second.
  // sets order in ascending order. Thereby when you set the rating to negative. The largest rating will
  // be first.
  unordered_map<string, set<pair<int, string>>> cuisine_ratings;
  unordered_map<string, string> food_cuisine;
  unordered_map<string, int> food_rating;
  FoodRatings(vector<string> &foods, vector<string> &cuisines, vector<int> &ratings)
  {
    for (int i = 0; i < foods.size(); ++i)
    {
      cuisine_ratings[cuisines[i]].insert({-ratings[i], foods[i]});
      food_cuisine[foods[i]] = cuisines[i];
      food_rating[foods[i]] = ratings[i];
    }
  }
  void changeRating(string food, int newRating)
  {
    auto &cuisine = food_cuisine.find(food)->second;
    int rating = food_rating[food];
    cuisine_ratings[cuisine].erase({-rating, food});
    cuisine_ratings[cuisine].insert({-newRating, food});
    food_rating[food] = newRating;
  }
  string highestRated(string cuisine)
  {
    return cuisine_ratings[cuisine].begin()->second;
  }
};