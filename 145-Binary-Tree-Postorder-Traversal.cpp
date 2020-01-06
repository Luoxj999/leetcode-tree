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
#include <stack>
#include <set>

using std::vector;
using std::stack;
using std::set;

struct TreeNode {
    int val;
    TreeNode *left;
    TreeNode *right;
    TreeNode(int x) : val(x), left(NULL), right(NULL) {}
};

class Solution {
public:
    vector<int> postorderTraversal(TreeNode* root) {
        if (root == NULL) {
            return {};
        }

        
        vector<int> result;
        stack<TreeNode*> stk;

        set<TreeNode*> visitedLeft;
        set<TreeNode*> visitedRight;

        stk.push(root);
        while (!stk.empty()) {
            TreeNode* tn = stk.top();
            stk.pop();

            if (visitedLeft.find(tn) == visitedLeft.end()) {
                visitedLeft.insert(tn);
                stk.push(tn);
                if (tn->left != NULL) {
                    stk.push(tn->left);
                }
                continue;
            }
            
            if (visitedRight.find(tn) == visitedRight.end()) {
                visitedRight.insert(tn);
                stk.push(tn);
                if (tn->right != NULL) {
                    stk.push(tn->right);
                }
                continue;
            }
            
            result.push_back(tn->val);
        }

        return result;
    }
};