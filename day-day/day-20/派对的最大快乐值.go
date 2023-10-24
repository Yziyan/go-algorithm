// @Author: Ciusyan 10/24/23

package day_20

// 公司的每个员工都符合 Employee 类的描述。整个公司的人员结构可以看作是一棵标准的、 没有环的多叉树
// 树的头节点是公司唯一的老板，除老板之外的每个员工都有唯一的直接上级
// 叶节点是没有任何下属的基层员工(subordinates列表为空)，除基层员工外每个员工都有一个或多个直接下级
// 这个公司现在要办party，你可以决定哪些员工来，哪些员工不来，规则：
// 1. 如果某个员工来了，那么这个员工的所有直接下级都不能来
// 2. 派对的整体快乐值是所有到场员工快乐值的累加
// 3. 你的目标是让派对的整体快乐值尽量大
// 给定一棵多叉树的头节点boss，请返回派对的最大快乐值。

type Employee struct {
	Happy int         // 这名员工可以带来的快乐值
	nexts []*Employee // 这名员工有哪些直接下级
}

func NewEmployee(happy int) *Employee {
	return &Employee{
		Happy: happy,
		nexts: make([]*Employee, 0),
	}
}

// MaxHappy 最大快乐值
func MaxHappy(boss *Employee) int {
	if boss == nil {
		return 0
	}

	// 定义 info
	type info struct {
		noMaxHappy  int // 不来的最大快乐值
		yesMaxHappy int // 来的最大快乐值
	}

	// 构建 info
	var getInfo func(boss *Employee) info
	getInfo = func(boss *Employee) info {
		if boss == nil {
			// 空节点好设置的情况
			return info{}
		}

		var (
			noMaxHappy int // 不来的最大快乐值
			// 来的最大快乐值，默认是自己的
			yesMaxHappy = boss.Happy
		)

		// 先收集所有子树的信息
		for _, next := range boss.nexts {
			nextInfo := getInfo(next)
			// 1. boss 不来：0 + 子节点max{来, 不来} 的最大快乐值
			nextMaxHappy := nextInfo.yesMaxHappy
			if nextInfo.noMaxHappy > nextMaxHappy {
				nextMaxHappy = nextInfo.noMaxHappy
			}
			noMaxHappy += nextMaxHappy

			// 2. boss 来：boss.happy + 所有子节点不来的最大快乐值
			yesMaxHappy += nextInfo.noMaxHappy
		}

		return info{
			noMaxHappy:  noMaxHappy,
			yesMaxHappy: yesMaxHappy,
		}
	}

	happyInfo := getInfo(boss)

	// 返回来和不来中，最大的
	if happyInfo.noMaxHappy > happyInfo.yesMaxHappy {
		return happyInfo.noMaxHappy
	}

	return happyInfo.yesMaxHappy
}
