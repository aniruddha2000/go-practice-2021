package main

type data struct {
	data string
}

func main() {
	k_new := new(data)
	k_new.data = "data by new"

	k_make := make([]data, 2)
	k_make = append(k_make, data{data: "data by make"})
}
