#include <vector>
#include <stack>

using std::vector;
using std::stack;

struct TreeNode {
    int val;
    TreeNode *left;
    TreeNode *right;
    TreeNode(int x) : val(x), left(NULL), right(NULL) {}
 };
/**
 * Definition for a binary tree node.
 * struct TreeNode {
 *     int val;
 *     TreeNode *left;
 *     TreeNode *right;
 *     TreeNode(int x) : val(x), left(NULL), right(NULL) {}
 * };
 */

// class Solution {
// public:
//     void visit(TreeNode* root, vector<int>& result) {
//         if (root == NULL) { return; }
//         result.push_back(root->val);
//         visit(root->left, result);
//         visit(root->right, result);

//     }
//     vector<int> preorderTraversal(TreeNode* root) {
//         vector<int> result;
//         visit(root, result);
//         return result;
//     }
// };

class Solution {
public:
    vector<int> preorderTraversal(TreeNode* root) {
        vector<int> result;
        if (root == NULL) {
            return result;
        }

        stack<TreeNode*> stk;
        stk.push(root);
        while (!stk.empty())
        {
            TreeNode* tn = stk.top();
            stk.pop();
            result.push_back(tn->val);
            if (tn->right != NULL) {
                stk.push(tn->right);
            }

            if (tn->left != NULL) {
                stk.push(tn->left);
            }
        }
        
        return result;
    }
};