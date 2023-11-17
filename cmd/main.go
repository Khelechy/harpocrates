package main

import (
	"fmt"

	"github.com/khelechy/harpocrates"
)

func main() {

	//harpocrates.Initialize(5)

	//var keys = []string{"ea995b38fc06e25844e815c22d9945fbfcd04afb0205568799cbecc0498501ceb69cb90edd", "a69b3a64e25ff1215fd0fd9eb156d3cf7c027f7161927100dbf72e21465c7b5737379d9f34", "4bed84274d8fc0167280a2507a8778a8a25c2505c6a0f73e38574b5792427336806fca86b0", "7f51338bc7e19de47e5516a0e3a84f1c8620cf6e5b7d68e1beca8bf1deb46e682921488287"}

	//harpocrates.Unseal(keys)

	ss := harpocrates.GetItem("name")

	fmt.Println(ss)
}
