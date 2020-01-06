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


// class Solution {
// public:
//     int maxDepth(TreeNode* root) {
//         if (root == NULL) {
//             return 0;
//         }

//         int left = maxDepth(root->left) + 1;
//         int right = maxDepth(root->right) + 1;

//         return left > right ? left : right;
//     }
// };

class Solution {
public:
    int maxDepth(TreeNode* root) {
        if (root == NULL) {
            return 0;
        }

        vector<vector<TreeNode*>> temp{{}, {}};

        int result = 0;
        int currIdx = 0;
        temp[currIdx].push_back(root);
        while (!temp[currIdx].empty()) {
            int nextIdx = (currIdx + 1) % 2;
            temp[nextIdx].clear();
            for (auto it = temp[currIdx].begin(); it != temp[currIdx].end(); it++) {
                if ((*it)->left != NULL) {
                    temp[nextIdx].push_back((*it)->left);
                }
                if ((*it)->right != NULL) {
                    temp[nextIdx].push_back((*it)->right);
                }
            }
            currIdx = nextIdx;
            result++;
        }

        return result;
    }
};