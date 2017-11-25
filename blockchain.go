package main

type Blockchain struct {
	chain []string
	currentTransactions []string
}

func newBlockchain() Blockchain {
	return Blockchain{}
}

func (b *Blockchain) newBlock() {
	// create new block and add it to chain
	return
}

func (b *Blockchain) newTransaction() {
	// add new transaction to list
	return
}

func (b *Blockchain) Hash(block string) {
	// hash block
	return
}

func (b *Blockchain) LastBlock() {
	// return last block of chain
	return
}
