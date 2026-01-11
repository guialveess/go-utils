package main

import (
	"context"
	"fmt"
)

func main() {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "token", "senha") // a ideia é que esse token seja passado por todo o fluxo da aplicação
	bookHotel(ctx, "Hotel Do Gui")
}

func bookHotel(ctx context.Context, name ...string) {
	token := ctx.Value("token").(string)
	fmt.Println(token)
}
