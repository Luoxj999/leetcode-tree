/**
 * Definition for a binary tree node.
 * struct TreeNode {
 *     int val;
 *     TreeNode *left;
 *     TreeNode *right;
 *     TreeNode(int x) : val(x), left(NULL), right(NULL) {}
 * };
 */

#include <vector>

using std::vector;

struct TreeNode {
    int val;
    TreeNode *left;
    TreeNode *right;
    TreeNode(int x) : val(x), left(NULL), right(NULL) {}
};

class Solution {
public:
    int findRootInInorder(int rootVal, vector<int>& inorder, int iLeft, int iRight) {
        for (int i = iLeft; i <= iRight; i++) {
            if (rootVal == inorder[i]) {
                return i;
            }
        }
        return -1;
    }

    TreeNode* solve(vector<int>& preorder, int pLeft, int pRight,
        vector<int>& inorder, int iLeft, int iRight) {

        if (pLeft > pRight || iLeft > iRight) {
            return NULL;
        }

        int rootVal = preorder[pLeft];
        TreeNode* tn = new TreeNode(rootVal);
        int inorderRootIdx = findRootInInorder(rootVal, inorder, iLeft, iRight);
        int inorderLeftTreeNodeNum = inorderRootIdx - iLeft;
        tn->left = solve(preorder, pLeft + 1, pLeft + inorderLeftTreeNodeNum,
            inorder, iLeft, inorderRootIdx - 1);
        
        tn->right = solve(preorder, pLeft + inorderLeftTreeNodeNum + 1, pRight,
            inorder, inorderRootIdx + 1, iRight);

        return tn;
    }

    TreeNode* buildTree(vector<int>& preorder, vector<int>& inorder) {
        if (preorder.size() == 0 || inorder.size() == 0) {
            return NULL;
        }

        return solve(preorder, 0, preorder.size() - 1, inorder, 0, inorder.size() - 1);
    }
};