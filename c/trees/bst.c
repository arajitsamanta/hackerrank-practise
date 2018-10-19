#include <stdio.h>
#include <string.h>
#include <math.h>
#include <stdlib.h>

//RepresentTion of a tree node
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

//Calculate maximum depth of a BST considering depth of root as one.
int getMaxDepth(struct node *root)
{
    if (root == NULL)
    {
        return 0;
    }
    return 1 + max(getMaxDepth(root->left), getMaxDepth(root->right));
}

//Calculate height of a BST considering height of leaf is at zero.
int getHeight(struct node *root)
{
    if (root == NULL)
    {
        return -1;
    }
    return 1 + max(getHeight(root->left), getHeight(root->right));
}

//Insert a node into a BST
struct node *insert(struct node *root, int data)
{

    if (root == NULL)
    {
        //When root is NULL. insert point is reached, create the node and return it
        struct node *node = (struct node *)malloc(sizeof(struct node));
        node->data = data;
        node->left = NULL;
        node->right = NULL;
        return node;
    }
    else
    {
        struct node *cur;

        //Recursivley search for insert point on left subtree if data <= node->data
        if (data <= root->data)
        {
            cur = insert(root->left, data);
            root->left = cur;
        }
        else
        {
            //Recursivley search for insert point on right subtree if data > node->data
            cur = insert(root->right, data);
            root->right = cur;
        }
        return root;
    }
}

//Function to find out lowest common ancestor of two nodes.
struct node *lca(struct node *root, int v1, int v2)
{

    if (root == NULL)
        return NULL;

    //Decide if you have to call recursively Samller than both
    if (root->data > v1 && root->data > v2)
        return lca(root->left, v1, v2);

    //Bigger than both
    if (root->data < v1 && root->data < v2)
        return lca(root->right, v1, v2);

    //Else solution already found
    return root;
}

//Recursive inorder traversal
void inorder(struct node *root)
{

    if (root == NULL)
        return;

    inorder(root->left);
    printf("%d ", root->data);
    inorder(root->right);
}

//Recursive inorder traversal which return the output as an array
void inorderArray(struct node *root, int arr[], int *i)
{
    if (root == NULL){
        arr[(*i)++]='$';
         return;
    }

    inorderArray(root->left, arr, i);
    arr[(*i)++] = root->data;
    inorderArray(root->right, arr, i);
}

//Recurive preorder traversal
void preorder(struct node *root)
{

    if (root == NULL)
        return;

    printf("%d ", root->data);
    inorder(root->left);
    inorder(root->right);
}

//Preorder traversal that returns an array.
void preorderArray(struct node *root, int arr[], int *i)
{

    if (root == NULL)
        return;

    arr[(*i)++] = root->data;
    preorderArray(root->left,arr,i);
    preorderArray(root->right,arr,i);
}

//Recurive postorder traversal
void postorder(struct node *root)
{

    if (root == NULL)
        return;

    postorder(root->left);
    postorder(root->right);
    printf("%d ", root->data);
}

//Postorder traversal that returns an array.
void postorderArray(struct node *root,int arr[],int *i)
{

    if (root == NULL)
        return;

    postorderArray(root->left,arr,i);
    postorderArray(root->right,arr,i);
    arr[(*i)++]=root->data;
}

//Populate abinary seacrh tree from an array
struct node *populateTree(int *arr, int n)
{
    //Determine array size.

    //Initialize root
    struct node *root = NULL;

    //Loop over array and invoke insert()
    for (int i = 0; i < n; i++)
    {
        root = insert(root, arr[i]);
    }
    return root;
}

void printArray(int arr[], int n)
{
    for (int j = 0; j < n; j++)
    {
        printf("%d ", arr[j]);
    }
}

int main()
{

    struct node *root = NULL;

    /*
    int t, n;
    int data;

    scanf("%d", &t);

    while (t-- > 0)
    {
        scanf("%d", &n);
        for (int i = 0; i < n; i++)
        {
            scanf("%d", &data);
            root = insert(root, data);
        }
    }*/

    int TESTCASE_1[] = {10, 4, 6, 3, 5, 7};
    int n = sizeof(TESTCASE_1) / sizeof(TESTCASE_1[0]);
    root = populateTree(TESTCASE_1, n);

    //Max depth
    printf("Max Depth: %d", getMaxDepth(root));

    //Height
    printf("\nHeight : %d", getHeight(root));

    //LCA
    struct node *lcaNode = lca(root, 2, 6);
    printf("\nLowest common ancester of 2 and 6 is: %d", lcaNode->data);

    /*
    Tree traversal
    */
    printf("\nInorder Traversal: ");
    inorder(root);

    //Inorder travel tha  //In order travel that returns an array insted of printing.
   /* int inorderArr[n];
    int i = 0;
    inorderArray(root, inorderArr, &i);
    n = sizeof(inorderArr) / sizeof(inorderArr[0]);
    printf("\nInorder Traversal[]: ");
    printArray(inorderArr,n); insted of printing.
    int inorderArr[n*2];  //In order travel that returns an array insted of printing.
   /* int inorderArr[n];
    int i = 0;
    inorderArray(root, inorderArr, &i);
    n = sizeof(inorderArr) / sizeof(inorderArr[0]);
    printf("\nInorder Traversal[]: ");
    printArray(inorderArr,n);
    int i = 0;
    inorderArray(root, inorderArr, &i);
    n = sizeof(inorderArr) / sizeof(inorderArr[0]);
    printf("\nInorder Traversal[]: ");
    printArray(inorderArr,n);

    printf("\nPreorder Traversal: ");
    preorder(root);

    printf("\nPostorder Traversal: ");
    postorder(root);

    return 0;
}