package interprog

import (
	"fmt"
	"github.com/gookit/color"
	"time"
)

// InterProg represents the progress tracker.
type InterProg struct {
	name          string
	logInterval   time.Duration
	startTime     time.Time
	count         int
	begin         time.Time
	printedUpdate bool
}

// NewInterProg creates a new InterProg instance.
func NewInterProg(name string, logInterval ...time.Duration) *InterProg {
	var interval = 5 * time.Second
	if len(logInterval) > 0 {
		interval = logInterval[0]
	}
	color.Tag("info").Println(name)
	return (&InterProg{
		name:        name,
		logInterval: interval,
		begin:       time.Now(),
	}).reset()
}

// reset restarts the timer.
func (ip *InterProg) reset() *InterProg {
	ip.startTime = time.Now()
	ip.count = 0
	return ip
}

// Update increments the counter and prints updates if needed.
func (ip *InterProg) Update(step ...int) {
	var updateStep = 1
	if len(step) > 0 {
		updateStep = step[0]
	}
	ip.count += updateStep
	var elapsedTime = time.Since(ip.startTime)
	var magenta = color.FgMagenta.Render
	if elapsedTime >= ip.logInterval {
		fmt.Printf("%s @ %s it/s \r", magenta(ip.name), magenta(ip.count/int(elapsedTime.Seconds())))
		ip.printedUpdate = true
		ip.reset()
	}
}

// Done prints the final elapsed time.
func (ip *InterProg) Done() {
	var elapsedTime = time.Since(ip.begin)
	if ip.printedUpdate {
		fmt.Println()
	}
	fmt.Printf("%s elapsed time: %v\n", ip.name, formatDuration(elapsedTime))
}
