/**
 * Definition for a binary tree node.
 * struct TreeNode {
 *     int val;
 *     TreeNode *left;
 *     TreeNode *right;
 *     TreeNode(int x) : val(x), left(NULL), right(NULL) {}
 * };
 */

// inorder = [9,3,15,20,7]
// postorder = [9,15,7,20,3]

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

    TreeNode* solve(vector<int>& postorder, int pLeft, int pRight,
        vector<int>& inorder, int iLeft, int iRight) {

        if (pLeft > pRight || iLeft > iRight) {
            return NULL;
        }

        int rootVal = postorder[pRight];
        TreeNode* tn = new TreeNode(rootVal);
        int inorderRootIdx = findRootInInorder(rootVal, inorder, iLeft, iRight);
        int inorderLeftTreeNodeNum = inorderRootIdx - iLeft;
        tn->left = solve(postorder, pLeft, pLeft + inorderLeftTreeNodeNum - 1,
            inorder, iLeft, inorderRootIdx - 1);
        
        tn->right = solve(postorder, pLeft + inorderLeftTreeNodeNum, pRight - 1,
            inorder, inorderRootIdx + 1, iRight);

        return tn;
    }

    TreeNode* buildTree(vector<int>& inorder, vector<int>& postorder) {
        if (postorder.size() == 0 || inorder.size() == 0) {
            return NULL;
        }

        return solve(postorder, 0, postorder.size() - 1, inorder, 0, inorder.size() - 1);
    }
};
