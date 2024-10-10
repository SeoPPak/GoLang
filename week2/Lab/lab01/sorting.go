package main

func sorting(list *[]User) {
	for i := 0; i < len(*list); i++ {
		for j := i + 1; j < len(*list); j++ {
			if (*list)[i].Name > (*list)[j].Name {
				Change(&(*list)[i], &(*list)[j])
			}
		}
	}
}
