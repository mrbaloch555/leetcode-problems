package main

import "fmt"

type ThroneInheritance struct {
	dead   map[string]bool
	family map[string][]string
	king   string
}

func Constructor(kingName string) ThroneInheritance {
	return ThroneInheritance{
		dead:   make(map[string]bool),
		family: make(map[string][]string),
		king:   kingName,
	}
}

func (this *ThroneInheritance) Birth(parentName string, childName string) {
	this.family[parentName] = append(this.family[parentName], childName)
}

func (this *ThroneInheritance) Death(name string) {
	this.dead[name] = true
}

func (this *ThroneInheritance) GetInheritanceOrder() []string {
	var ans []string
	this.dfs(this.king, &ans)
	return ans
}

func (this *ThroneInheritance) dfs(name string, ans *[]string) {
	if !this.dead[name] {
		*ans = append(*ans, name)
	}
	for _, child := range this.family[name] {
		this.dfs(child, ans)
	}
}

//     king
//   /     |    \
//  andy   bob   catherine
//       /       |
//    metthew  /  \
//             alex  asha

func main() {
	obj := Constructor("king")
	obj.Birth("king", "andy")
	obj.Birth("king", "bob")
	obj.Birth("king", "catherine")
	obj.Birth("andy", "matthew")
	obj.Birth("bob", "alex")
	obj.Birth("bob", "asha")
	fmt.Println(obj.GetInheritanceOrder())
	obj.Death("bob")
	fmt.Println(obj.GetInheritanceOrder())
}
