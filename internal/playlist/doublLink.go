package playlist

import "context"


type doubLink struct{
	 // https://pkg.go.dev/container/list
	storage
}


func newDoubLink() {
	return &doubLink{}
}



func (d *doubLink) PushBack(ctx context.Context) error{
	return nil
}


