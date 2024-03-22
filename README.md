# golang-concurrency
* TODO implement the key ideas from below.

# Channel patterns : https://go.dev/talks/2012/concurrency.slide#1
* Multiplexing
* Fan in
* Restoring sequence
* Daisy chain

# Channel use cases: https://go101.org/article/channel-use-cases.html
* Futures/Promises
  * Return receive-only channels as results
  * Pass send-only channels as arguments
  * The first response wins
* Channels for notifications
  * 1-to-1 notification by sending a value to a channel
  * 1-to-1 notification by receiving a value from a channel
  * N-to-1 and 1-to-N notifications
  * Broadcast (1-to-N) notifications by closing a channel
  * Timer: scheduled notification
* Mutex locks
* As Counting Semaphores
* Dialogue (Ping-Pong)
* Channel Encapsulated in Channel
* Peak/burst limiting
* timeout
* Ticker
* Rate limiting
* Data Flow Manipulations
  * Data generation/collecting/loading
  * Data aggregation
  * Data division
  * Data composition
  * Data decomposition
  * Data duplication/proliferation
  * Data calculation/analysis
  * Data validation/filtering
  * Data serving/saving
  * Data flow system assembling


