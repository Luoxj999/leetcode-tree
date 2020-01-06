#include <vector>
#include <set>
#include <stack>

using std::vector;
using std::set;
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
class Solution {
public:
    vector<int> inorderTraversal(TreeNode* root) {
        vector<int> result;
        if (root == NULL) {
            return result;
        }
        
        set<TreeNode*> st;
        stack<TreeNode*> stk;
        stk.push(root);
        while (stk.size() > 0) {
            TreeNode* tn = stk.top();
            stk.pop();

            // 第一次visit
            if (st.find(tn) == st.end()) {
                st.insert(tn);
                if (tn->left != NULL) {
                    stk.push(tn);
                    stk.push(tn->left);
                } else {
                    result.push_back(tn->val);
                    if (tn->right != NULL) {
                        stk.push(tn->right);
                    }
                }
            // 第二次visit
            } else {
                result.push_back(tn->val);
                if (tn->right != NULL) {
                    stk.push(tn->right);
                } 
            }
        }

        return result;
    }
};