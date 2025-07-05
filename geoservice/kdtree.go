package geoservice

// KDTree insert and search

type KDNode struct {
	Location
	Axis  int
	Left  *KDNode
	Right *KDNode
}

func Insert(root *KDNode, loc Location, depth int) *KDNode {
	if root == nil {
		return &KDNode{Location: loc, Axis: depth % 2}
	}

	var cmp float64

	if root.Axis == 0 {
		cmp = loc.Latitude - root.Latitude
	} else {
		cmp = loc.Longitude - root.Longitude
	}

	if cmp < 0 {
		root.Left = Insert(root.Left, loc, depth+1)
	} else {
		root.Right = Insert(root.Right, loc, depth+1)
	}
	return root
}

func RadiusSearch(root *KDNode, lat, lon, radius float64, depth int, result *[]Location) {
	if root == nil {
		return
	}

	distance := Haversine(lat, lon, root.Latitude, root.Longitude)
	if distance <= radius {
		*result = append(*result, root.Location)
	}

	var cmp float64
	if root.Axis == 0 {
		cmp = lat - root.Latitude
	} else {
		cmp = lon - root.Longitude
	}

	if cmp-radius <= 0 {
		RadiusSearch(root.Left, lat, lon, radius, depth+1, result)
	}

	if cmp-radius >= 0 {
		RadiusSearch(root.Right, lat, lon, radius, depth+1, result)
	}
}
