#include<algorithm>
#include<iostream>
#include<map>
#include<set>
#include<vector>

using namespace std;
using ll = long long;

struct TreeNode {
    TreeNode* parent;
    TreeNode* left;
    TreeNode* right;
    int val;
    int count;
    int height;
};

int count(TreeNode* root) {
    if (root == nullptr) {
        return 0;
    }
    return root->count;
}

int height(TreeNode* root) {
    if (root == nullptr) {
        return 0;
    }
    return root->height;
}

void print(TreeNode* root, string prefix = "") {
    if (root == nullptr) {
        return;
    }
    cout << prefix << root->val << '\n';
    print(root->left, prefix + "  ");
    print(root->right, prefix + "  ");
}

void print(const vector<int>& v) {
    bool first = true;
    for (auto val : v) {
        if (!first) {
            cout << ' ';
        }
        first = false;
        cout << val;
    }
    cout << '\n';
}

TreeNode* rotateLeftUp(TreeNode* root) {
    if (root == nullptr || root->left == nullptr) {
        return root;
    }
    TreeNode* oldLeft = root->left;
    TreeNode* oldLeftRight = oldLeft->right;

    root->left = oldLeftRight;
    root->count = count(root->left) + 1 + count(root->right);
    root->height = max(height(root->left), height(root->right)) + 1;

    oldLeft->right = root;
    oldLeft->count = count(oldLeft->left) + 1 + count(oldLeft->right);
    oldLeft->height = max(height(oldLeft->left), height(oldLeft->right)) + 1;

    return oldLeft;
}

TreeNode* rotateRightUp(TreeNode* root) {
    if (root == nullptr || root->right == nullptr) {
        return root;
    }
    TreeNode* oldRight = root->right;
    TreeNode* oldRightLeft = oldRight->left;

    root->right = oldRightLeft;
    root->count = count(root->left) + 1 + count(root->right);
    root->height = max(height(root->left), height(root->right)) + 1;

    oldRight->left = root;
    oldRight->count = count(oldRight->left) + 1 + count(oldRight->right);
    oldRight->height = max(height(oldRight->left), height(oldRight->right)) + 1;

    return oldRight;
}

TreeNode* insert(TreeNode* root, int val) {
    if (root == nullptr) {
        return new TreeNode {.parent=nullptr, .left=nullptr, .right=nullptr, .val=val, .count=1, .height=1}; 
    }
    if (val <= root->val) {
        root->left = insert(root->left, val);
        root->left->parent = root;
        root->height = max(root->height, root->left->height + 1);
    } else {
        root->right = insert(root->right, val);
        root->right->parent = root;
        root->height = max(root->height, root->right->height + 1);
    }
    root->count++;

    int leftHeight = height(root->left);
    int rightHeight = height(root->right);
    if (leftHeight - rightHeight == 2) {
        if (height(root->left->left) < height(root->left->right)) {
            root->left = rotateRightUp(root->left);
        }
        return rotateLeftUp(root);
    } else if (rightHeight - leftHeight == 2) {
        if (height(root->right->right) < height(root->right->left)) {
            root->right = rotateLeftUp(root->right);
        }
        return rotateRightUp(root);
    }
    return root;
}

void freeTree(TreeNode* root) {
    if (root == nullptr) {
        return;
    }
    if (root->left != nullptr) {
        freeTree(root->left);
    }
    if (root->right != nullptr) {
        freeTree(root->right);
    }
    delete root;
    return;
}

int countLE(TreeNode* root, int q) {
    if (root == nullptr) {
        return 0;
    }
    if (root->val <= q) {
        int prev = 1;
        if (root->left != nullptr) {
            prev += root->left->count;
        }
        return prev + countLE(root->right, q);
    } else {
        return countLE(root->left, q);
    }
}

int main() {
    int n, q; cin >> n >> q;
    vector<int> levels(n);
    for (int i = 0; i < n; i++) {
        cin >> levels[i];
    }

    vector<int> smallestKForYes(n);
    smallestKForYes[0] = 1;

    TreeNode* root = nullptr;
    root = insert(root, 1);

    for (int i = 1; i < n; i++) {
        int low = 0;
        int high = 200000;
        while (high - low > 1) {
            int mid = (high + low) / 2;
            if (countLE(root, mid) / mid + 1 > levels[i]) {
                low = mid;
            } else {
                high = mid;
            }
        }
        smallestKForYes[i] = high;
        root = insert(root, high);
    }
    freeTree(root);

    for (int i = 0; i < q; i++) {
        int idx, k; cin >> idx >> k;
        idx--;
        if (k >= smallestKForYes[idx]) {
            cout << "YES\n";
        } else {
            cout << "NO\n";
        }
    }
    return 0;
}