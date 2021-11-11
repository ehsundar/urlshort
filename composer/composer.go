package composer

type Composer interface {
	Compose(long, nonce string) (short string)
}
