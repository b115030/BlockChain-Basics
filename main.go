package main
import(
        "fmt"
        "bytes"
        "crypto/sha256"
)
type BlockChain struct {
blocks []*Block
}
type Block struct{
        Hash []byte
        Data []byte
        PrevHash []byte
}
func (b *Block) DeriveHash(){
info := bytes.Join([][]byte{b.Data, b.PrevHash}, []byte{})
hash := sha256.Sum256(info)
b.Hash = hash[:]
}
func CreateBlock(data string, prevHash []byte) *Block{
block := &Block{[]byte{}, []byte(data), prevHash}
block.DeriveHash()
return block
}
func (chain *BlockChain) AddBlock(data string){
prevBlock := chain.blocks[len(chain.blocks)-1]
new :=CreateBlock(data, prevBlock.Hash)
chain.blocks = append(chain.blocks, new)
}
func Genesis() *Block {
return CreateBlock("Genesis", []byte{})
}
func InitBlockChain() *BlockChain {
return &BlockChain{[]*Block{Genesis()}}
}
func main(){
        chain := InitBlockChain()
        chain.AddBlock("First after Gen")
        chain.AddBlock("Seconf after Gen")
		chain.AddBlock("hird after Gen")
		
		for _,block := range chain.blocks{
			fmt.Println("-----Block Starts-----")
			fmt.Printf("Previous Hash: %x\n", block.PrevHash)
			fmt.Printf("Data in Block: %s\n", block.Data)
			fmt.Printf("Hash of itself: %x\n", block.Hash)
			fmt.Println("-----Block ends-----")
		}
}
