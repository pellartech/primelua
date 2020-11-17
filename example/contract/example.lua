function concat(a)
	a = a .. "-concat"
	set_value("a", a)
	return get_value("a")
end
function set(a)
	set_value("a", a)
end
function get()
   return get_value("a")
end