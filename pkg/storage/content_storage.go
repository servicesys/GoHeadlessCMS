package storage



type ContentStorage interface {
	SearchContent(strSearch string ) []string
	GetContent(uuid string) (, error)
	//CreateContent(ctx context.Context, u User) (string, error)
	///UpdateContent(ctx context.Context, u User) error
	//DeleteContent(ctx context.Context, id string) error
}

