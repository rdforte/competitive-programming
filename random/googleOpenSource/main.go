package main

import (
	"context"
	"fmt"
	"time"

	"github.com/rdforte/competitive-programming/random/googleOpenSource/errGroup"
)

func main() {
	circleRadiuses := []int{1, 2, 3}
	results := make([]int, len(circleRadiuses))

	parentCtx := context.Background()
	g, ctx := errGroup.WithContext(parentCtx)

	for i, r := range circleRadiuses {
		index := i
		radius := r

		g.Go(func() error {
			d, err := getCircleDiameter(radius)
			if err != nil {
				return err
			}
			results[index] = d
			return nil
		})
	}

	if err := g.Wait(); err != nil {
		fmt.Println(ctx.Err())
		panic(err)
	}
}

func getCircleDiameter(radius int) (int, error) {
	if radius == 3 {
		return 0, fmt.Errorf("test error")
	}

	time.Sleep(time.Second * 3) // simulate latency
	return radius * 2, nil
}
