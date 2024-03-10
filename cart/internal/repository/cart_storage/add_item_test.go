package cartstorage

import (
	"context"
	"testing"
)

func BenchmarkAddItem1(b *testing.B) {
	ctx := context.Background()
	cartStorage := NewCartStorage()

	b.Run("Добавляем один продукт одному пользователю", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			err := cartStorage.AddItem(ctx, 1, 1, 1)
			if err != nil {
				b.Fail()
			}
		}
	})
}

func BenchmarkAddItem2(b *testing.B) {
	ctx := context.Background()
	cartStorage := NewCartStorage()

	b.Run("Добавляем один продукт разным пользователям", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			err := cartStorage.AddItem(ctx, int64(i), 1, 1)
			if err != nil {
				b.Fail()
			}
		}
	})

}

func BenchmarkAddItem3(b *testing.B) {
	ctx := context.Background()
	cartStorage := NewCartStorage()

	cartStorage.Reset()
	b.Run("Добавляем разные продукты одному пользователям", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			err := cartStorage.AddItem(ctx, 1, int64(i), 1)
			if err != nil {
				b.Fail()
			}
		}
	})
}

func BenchmarkAddItem4(b *testing.B) {
	ctx := context.Background()
	cartStorage := NewCartStorage()

	b.Run("Добавляем разные продукты разным пользоваелям", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			err := cartStorage.AddItem(ctx, int64(i), int64(i), 1)
			if err != nil {
				b.Fail()
			}
		}
	})
}
