# 二叉搜索树
题目来源：[leetcode](https://leetcode-cn.com/problems)



## 前菜

二叉树的三种遍历

- 前序遍历
- 中序遍历
- 后续遍历

### 树节点的定义

```python
# Definition for a binary tree node.
class TreeNode:
    def __init__(self, val=0, left=None, right=None):
		    self.val = val
        self.left = left
        self.right = right
```

### 前序遍历

```python
def PreOrder(root):
  if root:
    print(root.val) # 访问根节点
    PreOrder(root.left) # 遍历访问左子树
    PreOrder(root.right) # 遍历访问右子树
```

### 中序遍历

```python
def PreOrder(root):
  if root:
    PreOrder(root.left) # 遍历访问左子树
    print(root.val) # 访问根节点
    PreOrder(root.right) # 遍历访问右子树
```

### 后序遍历

```python
def PreOrder(root):
  if root:
    PreOrder(root.left) # 遍历访问左子树
    PreOrder(root.right) # 遍历访问右子树
    print(root.val) # 访问根节点
```



## 98. 验证二叉搜索树

给定一个二叉树，判断其是否是一个有效的二叉搜索树。

假设一个二叉搜索树具有如下特征：

- 节点的左子树只包含小于当前节点的数。
- 节点的右子树只包含大于当前节点的数。
- 所有左子树和右子树自身必须也是二叉搜索树。

``````shell
输入:
    2
   / \
  1   3
输出: true
``````

### 方法一：

根据二叉搜索树的特点，采用**中序遍历**可将树遍历为递增的数组。因此，解题时可以先将树做一个**中序遍历**。经过遍历之后的数组若为严格递增，则可任务该树为二叉搜索树。

代码如下

```python
# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution:
    def isValidBST(self, root: TreeNode) -> bool:
        lis = []
        def helper(root: TreeNode):
            if root.left:
                helper(root.left)
            if root:
                lis.append(root.val)
            if root.right:
                helper(root.right)
        helper(root)
        if len(lis) < 1:
            return True
        # 若遍历结果为严格递增，则验证通过
        for i in range(1, len(lis)):
            if lis[i-1] >= lis[i]:
                return False
        return True
```

