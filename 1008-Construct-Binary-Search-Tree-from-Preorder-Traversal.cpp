/**
 * Definition for a binary tree node.
 * struct TreeNode {
 *     int val;
 *     TreeNode *left;
 *     TreeNode *right;
 *     TreeNode(int x) : val(x), left(NULL), right(NULL) {}
 * };
 */

// Input: [8,5,1,7,10,12]
// Output: [8,5,10,1,7,null,12]

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
    int findBiggerOne(vector<int>& preorder, int left, int right, int val) {
        for (int i = left; i <= right; i++) {
            if (preorder[i] > val) {
                return i;
            }
        }
        return -1;
    }

    TreeNode* solve(vector<int>& preorder, int left, int right) {
        if (left > right) {
            return NULL;
        }

        int rootVal = preorder[left];
        TreeNode* tn = new TreeNode(rootVal);
        int idx = findBiggerOne(preorder, left, right, rootVal);
        if (idx == -1) {
            tn->left = solve(preorder, left + 1, right);
            tn->right = NULL;
        } else {
            tn->left = solve(preorder, left + 1, idx - 1);
            tn->right = solve(preorder, idx, right);
        }
        return tn;
    }

    TreeNode* bstFromPreorder(vector<int>& preorder) {
        return solve(preorder, 0, preorder.size() - 1);
    }
};