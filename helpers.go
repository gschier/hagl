package hagl

func FragmentRange(n int, child func(i int) Node) Node {
	fragment := Fragment()
	for i := 0; i < n; i++ {
		fragment.Children(child(i))
	}
	return fragment
}
