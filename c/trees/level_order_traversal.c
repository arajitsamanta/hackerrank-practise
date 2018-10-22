#include <stdio.h>
#include <stdlib.h>

struct node
{

  int data;
  struct node *left;
  struct node *right;
};

int max(int a, int b)
{
  return a > b ? a : b;
}

int height(struct node *root)
{
  if (root == NULL)
  {
    return 0;
  }
  return 1 + max(height(root->left), height(root->right));
}

void printLevel(struct node *root, int level)
{
  //Best case - when rot is NULL, return
  if (root == NULL)
  {
    return;
  }

  //If level is 1, print it. else recursively print left and right subtree.
  if (level == 1)
  {
    printf("%d ", root->data);
  }
  else
  {
    printLevel(root->left, level - 1);
    printLevel(root->right, level - 1);
  }
}

//https://www.hackerrank.com/challenges/tree-level-order-traversal/problem?utm_campaign=challenge-recommendation&utm_medium=email&utm_source=24-hour-campaign
void levelOrder(struct node *root)
{
  int h = height(root);
  for (int i = 1; i <= h; i++)
  {
    printLevel(root, i);
  }
}

struct node *new (int data)
{

  struct node *node = (struct node *)malloc(sizeof(struct node));
  node->data = data;
  node->left = NULL;
  node->right = NULL;
  return node;
}

int main()
{

  struct node *root = new (1);

  root->left = new (2);
  root->right = new (3);

  root->left->left = new (4);
  root->left->right = new (5);
  levelOrder(root);
  return 0;
}