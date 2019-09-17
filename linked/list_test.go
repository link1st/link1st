/**
* Created by GoLand.
* User: link1st
* Date: 2019/9/17
* Time: 16:46
 */

package linked

import (
	"fmt"
	"testing"
)

func printNode(nodes []*ListNode) (values []string) {
	values = make([]string, 0, len(nodes))
	for _, node := range nodes {
		values = append(values, node.GetValue())
	}

	return
}

func TestList_Index(t *testing.T) {
	list := NewList()

	list.RPush("1001")
	list.RPush("1002")
	list.RPush("1003")
	list.RPush("1004")

	fmt.Println("index", list.Index(0).GetValue(), list.Index(-1).GetValue())
	fmt.Println("index", list.Index(3).GetValue(), list.Index(-5).GetValue())

	fmt.Println("Range", printNode(list.Range(0, -1)))

	fmt.Println("pop", list.LPop().GetValue())
	fmt.Println("Range", printNode(list.Range(0, -1)))

	fmt.Println("pop", list.LPop().GetValue())
	fmt.Println("Range", printNode(list.Range(0, -1)))

	fmt.Println("pop", list.LPop().GetValue())
	fmt.Println("Range", printNode(list.Range(0, -1)))

	fmt.Println("pop", list.LPop().GetValue())
	fmt.Println("Range", printNode(list.Range(0, -1)))

	fmt.Println("pop", list.LPop().GetValue())
	fmt.Println("Range", printNode(list.Range(0, -1)))

}

func TestList_Range(t *testing.T) {
	list := NewList()
	fmt.Println("Range", printNode(list.Range(0, -1)))
	fmt.Println("Range", printNode(list.Range(0, 2)))
	fmt.Println("Range", printNode(list.Range(0, 100)))
	fmt.Println("Range", printNode(list.Range(-1, -2)))
	fmt.Println("Range", printNode(list.Range(-2, -1)))
	fmt.Println("Range", printNode(list.Range(-2, 3)))

	list.RPush("1001")
	list.RPush("1002")
	list.RPush("1003")
	list.RPush("1004")

	fmt.Println("Range", printNode(list.Range(0, -1)))
	fmt.Println("Range", printNode(list.Range(0, 2)))
	fmt.Println("Range", printNode(list.Range(0, 100)))
	fmt.Println("Range", printNode(list.Range(-1, -2)))
	fmt.Println("Range", printNode(list.Range(-2, -1)))
	fmt.Println("Range", printNode(list.Range(-2, 3)))

}
