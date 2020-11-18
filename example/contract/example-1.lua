function concat(a)
	a = a .. "-concat"
	set_value("a", a)
	return get_value("a")
end