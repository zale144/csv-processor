package retry

import "time"

type Operation func(attempt uint) error

func Retry(action Operation, limit uint, delay time.Duration) (err error) {

	for attempt := uint(0); attempt == 0 || (attempt < limit && err != nil); attempt++ {

		if attempt > 0 {
			time.Sleep(delay)
		}

		err = action(attempt)
	}

	return
}
