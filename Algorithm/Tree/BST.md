<!-- START doctoc generated TOC please keep comment here to allow auto update -->
<!-- DON'T EDIT THIS SECTION, INSTEAD RE-RUN doctoc TO UPDATE -->
**Table of Contents**  *generated with [DocToc](https://github.com/thlorenz/doctoc)*

- [二叉搜索树](#%E4%BA%8C%E5%8F%89%E6%90%9C%E7%B4%A2%E6%A0%91)
  - [前菜](#%E5%89%8D%E8%8F%9C)
    - [树节点的定义](#%E6%A0%91%E8%8A%82%E7%82%B9%E7%9A%84%E5%AE%9A%E4%B9%89)
    - [前序遍历](#%E5%89%8D%E5%BA%8F%E9%81%8D%E5%8E%86)
    - [中序遍历](#%E4%B8%AD%E5%BA%8F%E9%81%8D%E5%8E%86)
    - [后序遍历](#%E5%90%8E%E5%BA%8F%E9%81%8D%E5%8E%86)
  - [98. 验证二叉搜索树](#98-%E9%AA%8C%E8%AF%81%E4%BA%8C%E5%8F%89%E6%90%9C%E7%B4%A2%E6%A0%91)
    - [方法一：](#%E6%96%B9%E6%B3%95%E4%B8%80)
  - [99.恢复二叉搜索树](#99%E6%81%A2%E5%A4%8D%E4%BA%8C%E5%8F%89%E6%90%9C%E7%B4%A2%E6%A0%91)

<!-- END doctoc generated TOC please keep comment here to allow auto update -->

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

```shell
输入:
    2
   / \
  1   3
输出: true
```

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

## 99.恢复二叉搜索树

给你二叉搜索树的根节点 root ，该树中的两个节点被错误地交换。请在不改变其结构的情况下，恢复这棵树。

进阶：使用 O(n) 空间复杂度的解法很容易实现。你能想出一个只使用常数空间的解决方案吗？

示例 1：

```
输入：root = [1,3,null,null,2]
输出：[3,1,null,null,2]
解释：3 不能是 1 左孩子，因为 3 > 1 。交换 1 和 3 使二叉搜索树有效。
```

示例 2：

```
输入：root = [3,1,4,null,null,2]
输出：[2,1,4,null,null,3]
解释：2 不能在 3 的右子树中，因为 2 < 3 。交换 2 和 3 使二叉搜索树有效。
```


提示：

树上节点的数目在范围 [2, 1000] 内

