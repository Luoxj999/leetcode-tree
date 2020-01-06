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
    vector<vector<int>> levelOrder(TreeNode* root) {
        vector<vector<int>> result;
		if (root == NULL) {
			return result;
		}

		vector<vector<TreeNode*>> roll;
		vector<TreeNode*> a;
		roll.push_back(a);
		roll.push_back(a);

		int curIdx = 0;
		roll[curIdx].push_back(root);
		while (!roll[curIdx].empty()) {
			vector<int> temp;
			int nextIdx = (curIdx + 1) % 2;
			roll[nextIdx].clear();
			for (auto it = roll[curIdx].begin(); it != roll[curIdx].end(); it++) {
				temp.push_back((*it)->val);
				if ((*it)->left) {
					roll[nextIdx].push_back((*it)->left);
				}
				if ((*it)->right) {
					roll[nextIdx].push_back((*it)->right);
				}

			}
			curIdx = nextIdx;
			result.push_back(temp);
		}

		return result;
    }
};