package main

import (
	"context"
	"fmt"

	"route256.ozon.ru/project/cart/internal/pkg/errgroup"
)

func main() {

	eg, _ := errgroup.NewErrGroup(context.Background(), 10)

	for i := range 100 {
		i := i
		eg.Go(func() error {
			if i == 5 || i == 7 || i == 9 {
				return fmt.Errorf("error %d", i)
				// return nil
			} else {
				fmt.Println(i)
				return nil
			}
		})
	}

	if err := eg.Wait(); err != nil {
		fmt.Println(err)
	}

}
