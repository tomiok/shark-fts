package main

type index map[string][]int

func (idx index) add(docs []Doc) {
	for _, doc := range docs {
		for _, token := range analyze(doc.Text) {
			ids := idx[token]
			if ids != nil && ids[len(ids)-1] == doc.ID {
				continue
			}
			idx[token] = append(ids, doc.ID)
		}
	}
}

func (idx index) search(s string) [][]int {
	var r [][]int
	for _, token := range analyze(s) {
		if ids, ok := idx[token]; ok {
			r = append(r, ids)
		}
	}
	return r
}
