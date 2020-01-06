/**
 * Definition for a binary tree node.
 * struct TreeNode {
 *     int val;
 *     TreeNode *left;
 *     TreeNode *right;
 *     TreeNode(int x) : val(x), left(NULL), right(NULL) {}
 * };
 */

// Input: pre = [1,2,4,5,3,6,7], post = [4,5,2,6,7,3,1]
// Output: [1,2,3,4,5,6,7]

// 观察规律，递归解决

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
    int findLeftValInPost(int leftVal, vector<int>& postorder, int postLeft, int postRight) {
        for (int i = postLeft; i <= postRight; i++) {
            if (leftVal == postorder[i]) {
                return i;
            }
        }
        return -1;
    }

    TreeNode* solve(vector<int>& pre, int preLeft, int preRight,
        vector<int>& post, int postLeft, int postRight) {

        if (preLeft > preRight || postLeft > postRight) {
            return NULL;
        }

        if (preLeft == preRight && postLeft == postRight) {
            return new TreeNode(pre[preLeft]);
        }

        TreeNode* tn = new TreeNode(pre[preLeft]);

        int leftVal = pre[preLeft + 1];
        int preLeftInPost = findLeftValInPost(leftVal, post, postLeft, postRight);
        int leftTreeNum = preLeftInPost - postLeft + 1;
        tn->left = solve(pre, preLeft + 1, preLeft + leftTreeNum,
            post, postLeft, postLeft + leftTreeNum - 1);
        
        tn->right = solve(pre, preLeft + leftTreeNum + 1, preRight,
            post, postLeft + leftTreeNum, postRight - 1);

        return tn;
    }


    TreeNode* constructFromPrePost(vector<int>& pre, vector<int>& post) {
        return solve(pre, 0, pre.size() - 1, post, 0, post.size() - 1);
    }
};