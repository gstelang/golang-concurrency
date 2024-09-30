# golang-concurrency
* Implement the key ideas from below.

# Most common ones 
* Worker Pool Pattern
* Pipeline Pattern
* Fan-In Pattern
* Semaphore Pattern
* Generator Pattern

# Channel patterns: 
Rob pike's https://go.dev/talks/2012/concurrency.slide#1
* Multiplexing
* Fan in
* Restoring sequence
* Daisy chain

# Channel use cases: 
https://go101.org/article/channel-use-cases.html
* Futures/Promises
  * [Return receive-only channels as results](futures/example1/main.go)
  * [Pass send-only channels as arguments](futures/example2/main.go)
  * [The first response wins](futures/example3/main.go)
* Channels for notifications
  * Fasters are notified by slowers
    * [1-to-1 notification by sending a value to a channel](notifications/example1/main.go)
    * [1-to-1 notification by receiving a value from a channel](notifications/example1/main.go)
  * N notifications [N-to-1 and 1-to-N notifications](notificatinos/example3/main.go)
  * [Timer: scheduled notification](notifications/example4/main.go)
* Mutex locks
  * [Send only](mutex-locks/example1/main.go)
  * [Receive only](mutex-locks/example2/main.go)
* Data Flow Manipulations
  * [Data generation/collecting/loading](data-flow/example1/main.go)
  * Data aggregation
  * Data division
  * Data composition
  * Data decomposition
  * Data duplication/proliferation
  * Data calculation/analysis
  * Data validation/filtering
  * Data serving/saving
  * Data flow system assembling
* Other
  * As Counting Semaphores
  * Dialogue (Ping-Pong)
  * Channel Encapsulated in Channel
  * Peak/burst limiting
  * timeout
  * Ticker
  * Rate limiting

