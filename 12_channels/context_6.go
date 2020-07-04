package main
import(
	"fmt"
	"context"
)
func main(){
	ctx := context.WithValue(context.Background(), "GO", "Language")
	str(ctx, "GO")
	ctx = context.WithValue(ctx, "dog", "Gaston")
	str(ctx, "dog")

	str(ctx, "color")

}
func str(ctx context.Context, s string){
	if v := ctx.Value(s); v != nil{
		fmt.Println("value found", v)
		return
	}
	fmt.Println("value not found")
}