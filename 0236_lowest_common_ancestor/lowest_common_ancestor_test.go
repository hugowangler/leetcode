package leetcode

import (
	"reflect"
	"testing"
)

var val0 = 0
var val1 = 1
var val2 = 2
var val3 = 3
var val4 = 4
var val5 = 5
var val6 = 6
var val7 = 7
var val8 = 8

var exampleTree = []*int{&val3, &val5, &val1, &val6, &val2, &val0, &val8, nil, nil, &val7, &val4}
var exampleTree2 = []*int{&val1, &val2}
var exampleTree3 = []*int{&val2, &val1}
var exampleTree4 = []*int{&val3, &val1, &val4, nil, &val2}

func Test_lowestCommonAncestor(t *testing.T) {
	type args struct {
		tree []*int
		p    int
		q    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "example 1",
			args: args{
				tree: exampleTree,
				p:    5,
				q:    1,
			},
			want: 3,
		},
		{
			name: "example 2",
			args: args{
				tree: exampleTree,
				p:    5,
				q:    4,
			},
			want: 5,
		},
		{
			name: "example 3",
			args: args{
				tree: exampleTree2,
				p:    1,
				q:    2,
			},
			want: 1,
		},
		{
			name: "example 4",
			args: args{
				tree: exampleTree3,
				p:    1,
				q:    2,
			},
			want: 2,
		},
		{
			name: "example 5",
			args: args{
				tree: exampleTree4,
				p:    2,
				q:    4,
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				root := buildTree(tt.args.tree, 0)
				if got := lowestCommonAncestor(
					root,
					findNode(tt.args.p, root),
					findNode(tt.args.q, root),
				); !(got.Val == tt.want) {
					t.Errorf("lowestCommonAncestor() = %v, want %v", got.Val, tt.want)
				}
			},
		)
	}
}

func Test_findNode(t *testing.T) {
	type args struct {
		node int
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{
			name: "should find the node",
			args: args{
				node: 5,
				root: buildTree(exampleTree, 0),
			},
			want: &TreeNode{
				Val: 5,
				Left: &TreeNode{
					Val:   6,
					Left:  nil,
					Right: nil,
				},
				Right: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val:   7,
						Left:  nil,
						Right: nil,
					},
					Right: &TreeNode{
						Val:   4,
						Left:  nil,
						Right: nil,
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				if got := findNode(tt.args.node, tt.args.root); !reflect.DeepEqual(got, tt.want) {
					t.Errorf("findNode() = %v, want %v", got, tt.want)
				}
			},
		)
	}
}
